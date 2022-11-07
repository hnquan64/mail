package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/modules/rbac/controller"
)

func InitRouterRBAC(rg *gin.RouterGroup) {
	rbacRouter := rg.Group("/auth")
	rbacRouter.POST("/login", controller.Login)
}
