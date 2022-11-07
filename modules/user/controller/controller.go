package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/entities"
	"gitlab.com/meta-node/mail/helper"
)

var dbConn = database.InstanceDB()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"code": 1, "message": err.Error(), "data": nil})
			return
		}
		user.Password = helper.HashPassword(user.Password)
		if err := dbConn.DB.Create(&user).Error; err != nil {
			c.AbortWithStatusJSON(400, gin.H{"code": 1, "message": err.Error(), "data": nil})
			return
		}
		c.AbortWithStatusJSON(200, gin.H{"code": 0, "message": "success", "data": user})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Param("id"))
		var user entities.User
		if err := dbConn.DB.Preload("MailStorage").Find(&user, "id = ?", userID).Error; err != nil {
			c.AbortWithStatusJSON(
				400,
				gin.H{
					"code":    1,
					"message": err.Error(),
					"data":    nil})
			return
		}
		c.AbortWithStatusJSON(
			200,
			gin.H{
				"code":    0,
				"message": "Get user successful",
				"data":    user})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []entities.User
		if err := dbConn.DB.Preload("MailStorage").Find(&users).Error; err != nil {
			c.AbortWithStatusJSON(
				400, gin.H{
					"code":    1,
					"message": err.Error(),
					"data":    nil})
			return
		}
		c.AbortWithStatusJSON(
			200, gin.H{
				"code":    0,
				"message": "Get users successful",
				"data":    users})
	}
}
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.AbortWithStatusJSON(
				400, gin.H{
					"code":    1,
					"message": err.Error(),
					"data":    nil})
			return
		}
		var user entities.User

		if err := dbConn.DB.Where("id = ?", id).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(
				400, gin.H{
					"code":    1,
					"message": err.Error(),
					"data":    nil})
			return
		}
		c.ShouldBind(&user)
		dbConn.DB.Save(&user)

		c.AbortWithStatusJSON(
			200, gin.H{
				"code":    0,
				"message": "Update successful",
				"data":    user})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Param("id"))

		if err := dbConn.DB.Delete(&entities.User{}, "id = ?", userID).Error; err != nil {
			c.AbortWithStatusJSON(
				400, gin.H{
					"code":    1,
					"message": err.Error(),
					"data":    nil})
			return
		}
		c.AbortWithStatusJSON(
			200, gin.H{
				"code":    0,
				"message": "Delete successful",
				"data":    nil})
	}
}
