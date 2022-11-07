package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/entities"
	"gitlab.com/meta-node/mail/helper"
)

var dbConn = database.InstanceDB()

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	c.ShouldBind(&loginRequest)
	if loginRequest.Email == "" || loginRequest.Password == "" {
		c.AbortWithStatusJSON(
			400, gin.H{
				"code":    1,
				"message": "Email or Password can't empty",
				"data":    nil})
		return
	}

	user := getUser(loginRequest.Email)

	check := helper.CheckPasswordHash(loginRequest.Password, user.Password)
	if !check {
		c.AbortWithStatusJSON(
			400, gin.H{
				"code":    1,
				"message": "Password is incorrect",
				"data":    nil})
		return
	}

	token := helper.GenerateToken(user)

	c.AbortWithStatusJSON(
		200, gin.H{
			"code":    0,
			"message": "Login successful",
			"data":    token})
}

func getUser(email string) entities.User {
	var dbConn = database.InstanceDB()
	var user entities.User
	if err := dbConn.DB.Find(&user, "email = ? ", email).Error; err != nil {
		log.Println("Email not found: ", email)
	}
	return user
}
