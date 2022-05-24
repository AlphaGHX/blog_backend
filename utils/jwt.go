package utils

import (
	"blog/global"
	"blog/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetToken(userName string) (string, error) {
	expireTime := time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Second)
	claims := models.Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(global.PRIVATE_KEY)
	return token, err
}

func VerifyToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("验证Token的加密类型错误")
		}
		return global.PUBLICK_KEY, nil
	})

	if token.Valid && err == nil {
		return claims, nil
	} else {
		return nil, err
	}
}
