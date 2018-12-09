package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

//JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (

	TokenExpired       = errors.New("Token is expired")
	TokenNotValidYet   = errors.New("Token not active yet")
	TokenMalformed     = errors.New("That's not even a token")
	TokenInvalid       = errors.New("Couldn't handle this token:")
	SignKey           = "newtrekWang"
)


//自定义载荷
type CustomClaims struct {
	jwt.StandardClaims
	ID   		string  	`json:"user_id"`
	Account 	string		`json:"account"`
}

func GetSignKey() string  {
	return SignKey
}

func SetSignKey(key string)   {
	SignKey = key
}


func NewJWT() *JWT  {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func (j *JWT)GenToken(claims CustomClaims)(string, error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}
