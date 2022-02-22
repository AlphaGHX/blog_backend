package utils

import (
	"blog/global"
	"blog/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(global.CONFIG.Jwt.Secret)

func GetToken(userName string) (string, error) {
	expireTime := time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Second)
	claims := models.Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Alpha",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, err
}

func VerifyToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
