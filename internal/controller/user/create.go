package user

import (
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/model"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/auth"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/resp"
	"github.com/gin-gonic/gin"
)

// Create add new user to the storage.
func (u *UserController) Create(c *gin.Context) {
	log.Info("user create function called.")
	var r model.User

	if err := c.ShouldBindJSON(&r); err != nil {
		resp.JSON(c, resp.ErrBind, nil)
		return
	}

	r.Password, _ = auth.Encrypt(r.Password)
	// Insert the user to the storage.
	if err := u.srv.Users().Create(c, &r); err != nil {
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, nil)
}
