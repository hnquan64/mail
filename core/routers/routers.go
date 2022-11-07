package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/middlewares"
	emailRouter "gitlab.com/meta-node/mail/modules/email/router"
	rbacRouter "gitlab.com/meta-node/mail/modules/rbac/router"
	userRouter "gitlab.com/meta-node/mail/modules/user/router"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	apiGroup := r.Group("api")
	rbacRouter.InitRouterRBAC(apiGroup)
	userRouter.InitRouterUser(apiGroup)
	apiGroup.Use(middlewares.Authorization)
	emailRouter.InitRouterEmail(apiGroup)

	return r
}
