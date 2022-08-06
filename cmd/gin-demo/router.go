package gindemo

import (
	"github.com/dyjwl/gin-web-plugin-demo/internal/controller/user"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/mysql"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/resp"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	jwtStrategy, _ := newJWTAuth().(JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	// Refresh time can be longer than token timeout
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		resp.JSON(c, resp.ErrPageNotFound, nil)
	})

	storeIns, _ := mysql.GetMysqlFactoryOr(nil)
	v1 := g.Group("/api/v1")
	{
		userv1 := v1.Group("/user")
		{
			userController := user.NewUserController(storeIns)

			userv1.POST("", userController.Create)
			userv1.Use(auto.AuthFunc(), Validation())
		}
	}

	return g
}
