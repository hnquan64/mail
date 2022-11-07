package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/emersion/go-sasl"
	"github.com/gin-gonic/gin"
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/entities"
	"gitlab.com/meta-node/mail/helper"
	"gitlab.com/meta-node/mail/server"
)

const (
	APPROVED = "Approved"
)

var dbConn = database.InstanceDB()

type EmailProcessRequest struct {
	EmailID uint   `json:"email_id"`
	Status  string `json:"status"`
}

func EmailProcessByManager() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, _ := ctx.Get("User")
		user := value.(entities.User)
		var emailProcessResp EmailProcessRequest
		if err := ctx.ShouldBind(&emailProcessResp); err != nil {
			ctx.AbortWithStatusJSON(
				401,
				gin.H{"code": 1,
					"message": err.Error(),
					"data":    nil})
			return
		}

		var email entities.Email
		if err := dbConn.DB.First(&email, "id = ?", emailProcessResp.EmailID).Error; err != nil {
			ctx.AbortWithStatusJSON(
				401,
				gin.H{"code": 1,
					"message": err.Error(),
					"data":    nil})
			return
		}

		if emailProcessResp.Status == APPROVED {
			go sendMail(email)
		}

		go func() {
			dbConn.DB.Model(&email).Update("status", emailProcessResp.Status)

			var emailManager entities.EmailManager
			emailManager.HandlerID = user.ID
			emailManager.EmailID = email.ID
			emailManager.Status = emailProcessResp.Status
			dbConn.DB.Create(&emailManager)
		}()

		ctx.AbortWithStatusJSON(200, gin.H{"code": 0, "message": fmt.Sprintf("%v successful email id: %v", emailProcessResp.Status, emailProcessResp.EmailID), "data": nil})
	}
}

func sendMail(email entities.Email) {
	password, err := helper.GetPassword(email.From)
	if err != nil {
		log.Fatal("Can't get password from email")
		return
	}
	mailHandler := server.MailHandler{
		Flag:  make(chan error),
		Email: &email,
		Auth:  sasl.NewPlainClient("", email.From, password),
		// Auth:  sasl.NewPlainClient("", "kuquandx1999@be.earning.com", "123456"),
	}
	mailHandler.SendmailHandler()
}

func CreateEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var email entities.Email
		if err := c.ShouldBind(&email); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"code": 1, "message": err.Error(), "data": nil})
			return
		}
		// fmt.Println(email)
		// email.Status = PENDING
		if err := dbConn.DB.Create(&email).Error; err != nil {
			c.AbortWithStatusJSON(400, gin.H{"code": 1, "message": err.Error(), "data": nil})
			return
		}
		c.AbortWithStatusJSON(200, gin.H{"code": 0, "message": "success", "data": email})
	}
}

func GetEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID, _ := strconv.Atoi(c.Param("id"))
		var email entities.Email
		if err := dbConn.DB.Find(&email, "id = ?", emailID).Error; err != nil {
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
				"message": "Get Email successful",
				"data":    email})
	}
}

func GetEmails() gin.HandlerFunc {
	return func(c *gin.Context) {
		var emails []entities.Email
		if err := dbConn.DB.Find(&emails).Error; err != nil {
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
				"message": "Get Emails successful",
				"data":    emails})
	}
}
func UpdateEmail() gin.HandlerFunc {
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
		var email entities.Email

		if err := dbConn.DB.Where("id = ?", id).First(&email).Error; err != nil {
			c.AbortWithStatusJSON(
				400, gin.H{
					"code":    1,
					"message": err.Error(),
					"data":    nil})
			return
		}
		c.ShouldBind(&email)
		dbConn.DB.Save(&email)

		c.AbortWithStatusJSON(
			200, gin.H{
				"code":    0,
				"message": "Update successful",
				"data":    email})
	}
}

func DeleteEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID, _ := strconv.Atoi(c.Param("id"))

		if err := dbConn.DB.Delete(&entities.Email{}, "id = ?", emailID).Error; err != nil {
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
