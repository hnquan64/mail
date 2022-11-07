package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/entities"
	"gitlab.com/meta-node/mail/helper"
)

var dbConn = database.InstanceDB()

func Authorization(c *gin.Context) {
	var user entities.User
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		c.AbortWithStatusJSON(401, gin.H{"code": 98, "message": "Missing Authorization", "data": nil})
		return
	}
	token := auth[0]
	claims, err := helper.VerifyToken(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"code": 98, "message": "Not Authorize " + err.Error(), "data": nil})
		return
	}

	if err := dbConn.DB.First(&user, "id = ?", claims.AccountID).Error; err != nil {
		c.AbortWithStatusJSON(401, gin.H{"code": 98, "message": "Not Authorize " + err.Error(), "data": nil})
		return
	}
	c.Set("User", user)
	c.Next()

}
