package initialize

import (
	"blog/global"
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
)

func InitKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKeyByte, err := ioutil.ReadFile(global.CONFIG.Jwt.PrivateKeyPath)
	fmt.Println("INIT\t PrivateKey & PublicKey")
	if err != nil {
		panic(fmt.Errorf("ERROR\t Loading PrivateKey: %v", err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		panic(fmt.Errorf("ERROR\t Parse PrivateKey: %v", err))
	}

	publickKeyByte, err := ioutil.ReadFile(global.CONFIG.Jwt.PublicKeyPath)
	if err != nil {
		panic(fmt.Errorf("ERROR\t Loading PublicKey: %v", err))
	}
	publickKey, err := jwt.ParseRSAPublicKeyFromPEM(publickKeyByte)
	if err != nil {
		panic(fmt.Errorf("ERROR\t Parse PublicKey: %v", err))
	}
	fmt.Println("DONE\t PrivateKey & PublicKey")
	return privateKey, publickKey
}
