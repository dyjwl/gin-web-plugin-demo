package gindemo

import (
	"strings"

	"github.com/dyjwl/gin-web-plugin-demo/pkg/resp"
	"github.com/gin-gonic/gin"
)

const authHeaderCount = 2

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
// according `Authorization` header.
type AutoStrategy struct {
	basic BasicStrategy
	jwt   JWTStrategy
}

var _ AuthStrategy = &AutoStrategy{}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(basic BasicStrategy, jwt JWTStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

// AuthFunc defines auto strategy as the gin authentication middleware.
func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeaderCount {
			resp.JSON(
				c,
				resp.ErrInvalidAuthHeader,
				nil,
			)
			c.Abort()

			return
		}

		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "Bearer":
			operator.SetStrategy(a.jwt)
			// a.JWT.MiddlewareFunc()(c)
		default:
			resp.JSON(c, resp.ErrSignatureInvalid, nil)
			c.Abort()

			return
		}

		operator.AuthFunc()(c)

		c.Next()
	}
}
