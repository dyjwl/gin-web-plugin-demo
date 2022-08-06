package gindemo

import (
	"errors"
	"net/http"

	"github.com/dyjwl/gin-web-plugin-demo/internal/store"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/resp"
	"github.com/gin-gonic/gin"
)

// Validation make sure users have the right resource permission and operation.
func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := isAdmin(c); err != nil {
			switch c.FullPath() {
			case "/v1/users":
				if c.Request.Method != http.MethodPost {
					resp.JSON(c, resp.ErrPermissionDenied, nil)
					c.Abort()

					return
				}
			case "/v1/users/:name", "/v1/users/:name/change_password":
				username := c.GetString("username")
				if c.Request.Method == http.MethodDelete ||
					(c.Request.Method != http.MethodDelete && username != c.Param("name")) {
					resp.JSON(c, resp.ErrPermissionDenied, nil)
					c.Abort()

					return
				}
			default:
			}
		}

		c.Next()
	}
}

func isAdmin(c *gin.Context) error {
	username := c.GetString(UsernameKey)
	user, err := store.Client().Users().Get(c, username)
	if err != nil {
		return err
	}

	if user.IsAdmin != 1 {
		return errors.New("current user is no permission")
	}

	return nil
}
