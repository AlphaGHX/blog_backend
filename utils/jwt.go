package utils

import (
	"blog/global"
	"blog/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetToken(userName string) (string, error) {
	jwtSecret := global.KEY_FILE
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
	jwtSecret := global.KEY_FILE
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	//保证算法一致
	if token.Header["alg"] == "HS256" && token.Valid && err == nil {
		return claims, nil
	} else {
		return nil, err
	}
}
