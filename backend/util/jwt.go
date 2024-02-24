package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your_secret_key") // 使用自定义密钥

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT 生成JWT令牌
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // 令牌有效期1小时
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}
