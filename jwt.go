package u

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

type jwtObj struct{}

var _jwt *jwtObj
var _jwtOnce sync.Once

// JWT 获取jwtObj对象
func JWT() *jwtObj {
	_jwtOnce.Do(func() {
		_jwt = &jwtObj{}
	})
	return _jwt
}

// GenerateToken 生成token
func (receiver *jwtObj) GenerateToken(secret interface{}, claims jwt.MapClaims) (string, error) {
	// 默认24小时
	if claims == nil {
		claims = jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

// ParseToken 解析token
func (receiver *jwtObj) ParseToken(tokenString string, secret []byte) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if token.Valid {
		return token.Claims, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New(fmt.Sprintf("That's not even a token"))
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, errors.New(fmt.Sprintf("Timing is everything"))
		} else {
			return nil, errors.New(fmt.Sprintf("Couldn't handle this token: %v", err))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Couldn't handle this token: %v", err))
	}
}
