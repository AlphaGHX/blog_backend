package initialize

import (
	"blog/global"
	"crypto/rsa"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
)

func InitKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKeyByte, err := ioutil.ReadFile(global.CONFIG.Jwt.PrivateKeyPath)
	if err != nil {
		global.STD_LOG.Errorln("Loading PrivateKey: ", err)
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		global.STD_LOG.Errorln("Loading PrivateKey: ", err)
		panic(err)
	}

	publickKeyByte, err := ioutil.ReadFile(global.CONFIG.Jwt.PublicKeyPath)
	if err != nil {
		global.STD_LOG.Errorln("Loading PublicKey: ", err)
		panic(err)
	}
	publickKey, err := jwt.ParseRSAPublicKeyFromPEM(publickKeyByte)
	if err != nil {
		global.STD_LOG.Errorln("Parse PublicKey: ", err)
		panic(err)
	}
	return privateKey, publickKey
}
