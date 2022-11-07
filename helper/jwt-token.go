package helper

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/meta-node/mail/core/entities"
)

var secret_key = []byte("mailservice")

type Claims struct {
	AccountID uint   `json:"account_id"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(account entities.User) string {
	expirationTime := time.Now().Add(1200000 * time.Hour)
	claims := &Claims{
		AccountID: account.ID,
		Email:     account.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(secret_key)
	if err != nil {
		log.Fatal("Can't generate token on this user: ", err)
	}
	return token
}

func VerifyToken(token string) (*Claims, error) {
	err := errors.New("invalid token")
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, err
		}
		return secret_key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
