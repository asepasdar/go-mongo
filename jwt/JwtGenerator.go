package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//MyClaims : custom claims info on jwt
type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	DeviceID string `json:"DeviceID"`
}

//M : map interface
type M map[string]interface{}

//ApplicationName : app name for application name
var ApplicationName = "My Simple Golang"

//LoginExpirationDuration : login expiration duration
var LoginExpirationDuration = time.Duration(1) * time.Hour

//JwtSigningMethod : jwt method to hash
var JwtSigningMethod = jwt.SigningMethodHS256

//JwtSignatureKey : some secret key
var JwtSignatureKey = []byte("this is my secret key")

type key string

const (
	//KeyInfo : get info from jwt token
	KeyInfo key = "userInfo"
)

//CreateJWT : function to create JWT token
func CreateJWT(username string, deviceID string) string {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: time.Now().Add(LoginExpirationDuration).Unix(),
		},
		Username: username,
		DeviceID: deviceID,
	}

	token := jwt.NewWithClaims(JwtSigningMethod, claims)
	tokenResult, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return tokenResult
}

//DecodeJWT : decode jwt
func DecodeJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JwtSignatureKey, nil
	})

	return token, err
}
