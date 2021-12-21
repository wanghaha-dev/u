package u

import (
	"errors"
	"fmt"
	"github.com/develop1024/jwt-go"
	"sync"
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
func (receiver *jwtObj) GenerateToken(secret string, claims Map) (token string, err error) {
	_mapClaims := jwt.MapClaims{}
	for k, v := range claims {
		_mapClaims[k] = v
	}
	// 创建一个新的令牌对象，指定签名方法和声明
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, _mapClaims)
	// 使用密码签名并获得完整的编码令牌作为字符串
	token, err = tokenObj.SignedString([]byte(secret))
	return
}

// ParseToken 解析token
func (receiver *jwtObj) ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	// Bad token
	if token == nil {
		return nil, errors.New("bad token")
	}

	// Invalid token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
