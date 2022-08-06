package gindemo

import (
	"encoding/base64"
	"strings"

	"github.com/dyjwl/gin-web-plugin-demo/pkg/resp"
	"github.com/gin-gonic/gin"
)

// BasicStrategy defines Basic authentication strategy.
type BasicStrategy struct {
	compare func(username string, password string) bool
}

var _ AuthStrategy = &BasicStrategy{}

// NewBasicStrategy create basic strategy with compare function.
func NewBasicStrategy(compare func(username string, password string) bool) BasicStrategy {
	return BasicStrategy{
		compare: compare,
	}
}

// AuthFunc defines basic strategy as the gin authentication middleware.
func (b BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			resp.JSON(c, resp.ErrSignatureInvalid, nil)
			c.Abort()

			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			resp.JSON(c, resp.ErrSignatureInvalid, nil)
			c.Abort()

			return
		}

		c.Set(UsernameKey, pair[0])

		c.Next()
	}
}
