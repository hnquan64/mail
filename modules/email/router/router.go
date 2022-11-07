package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/modules/email/controller"
)

func InitRouterEmail(rg *gin.RouterGroup) {
	emailRouter := rg.Group("/email")
	emailRouter.POST("/", controller.CreateEmail())
	emailRouter.GET("/", controller.GetEmails())
	emailRouter.GET("/:id", controller.GetEmail())
	emailRouter.PUT("/:id", controller.UpdateEmail())
	emailRouter.DELETE("/:id", controller.DeleteEmail())
	emailRouter.POST("/manager", controller.EmailProcessByManager())
}
