package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
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

func (j *JWT)ParseToken(tokenString string)(*CustomClaims, error)  {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT)RefreshToken(tokenString string)(string, error)  {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0,0)
	}
	token , err := jwt.ParseWithClaims(tokenString,&CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey,nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2*24 *time.Hour).Unix()
		return j.GenToken(*claims)
	}
	return "",TokenInvalid
}
