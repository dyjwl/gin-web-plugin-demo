package user

import (
	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"github.com/gin-gonic/gin"
)

type ListOptions struct {
}

func (u *UserController) List(c *gin.Context) {
	log.Info("list user function called.")
}
