package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string
	jwt.StandardClaims
}
