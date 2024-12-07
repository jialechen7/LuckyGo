package utility

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jialechen7/go-lottery/common/constants"
	"time"
)

// GenerateJWT 生成一个 JWT
func GenerateJWT(secretKey []byte, expirationSeconds int64, userId int64) (string, error) {
	// 创建令牌，设置声明
	claims := jwt.MapClaims{
		"exp":                       time.Now().Add(time.Second * time.Duration(expirationSeconds)).Unix(),
		constants.JwtClaimUserIdKey: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名令牌
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWT 解析和验证一个 JWT
func ParseJWT(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	// 解析令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌是否有效并返回声明
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
