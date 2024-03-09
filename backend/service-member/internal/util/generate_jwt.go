package util

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var LoginExpirationDuration = time.Duration(24) * time.Hour
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("secretThisIsSecretKey")

type ServiceClaim struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type JwtGeneratorImpl struct {
}

func NewJWTGenerator() JwtGenerator {
	return &JwtGeneratorImpl{}
}

type JwtGenerator interface {
	NewAccessToken(email string, name string) (string, error)
	VerifyToken(token string) error
}

func (j JwtGeneratorImpl) NewAccessToken(email string, name string) (string, error) {
	claims := ServiceClaim{
		Username: name,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Service Auth",
			ExpiresAt: time.Now().Add(LoginExpirationDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(JwtSigningMethod, claims)
	return token.SignedString([]byte(JwtSignatureKey))
}

func (j JwtGeneratorImpl) VerifyToken(token string) error {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return JwtSignatureKey, nil
	})
	if err != nil {
		log.Printf("Error to parse jwt: {%s}", err.Error())
		return err
	}

	if !parse.Valid {
		log.Printf("Token not vaild")
		return err
	}

	return nil
}
