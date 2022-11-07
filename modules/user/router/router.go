package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/modules/user/controller"
)

func InitRouterUser(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user")
	userRouter.POST("/", controller.CreateUser())
	userRouter.GET("/", controller.GetUsers())
	userRouter.GET("/:id", controller.GetUser())
	userRouter.PUT("/:id", controller.UpdateUser())
	userRouter.DELETE("/:id", controller.DeleteUser())
}
