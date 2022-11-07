package helper

import (
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/entities"
	"golang.org/x/crypto/bcrypt"
)

var dbConn = database.InstanceDB()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetPassword(email string) (string, error) {
	var password string
	if err := dbConn.DB.Model(&entities.User{}).Select("password").Where("email = ?", email).Scan(&password).Error; err != nil {
		return "", err
	}
	return password, nil
}
