package gindemo

import (
	"fmt"

	"github.com/dyjwl/gin-web-plugin-demo/configs"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	initRouter(router)
	router.Run(fmt.Sprintf(":%d", configs.Config.Server.Port))
}
