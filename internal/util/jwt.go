package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserId uint
	jwt.RegisteredClaims
}

const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

var signingMethods = map[string]jwt.SigningMethod{
	HS256: jwt.SigningMethodHS256,
	HS384: jwt.SigningMethodHS384,
	HS512: jwt.SigningMethodHS512,
}

func getSignedMethod(signed string) jwt.SigningMethod {
	val, ok := signingMethods[signed]
	if !ok {
		return signingMethods[HS256]
	}
	return val
}

func GenerateToken(userId uint, secretKey string, signed string, expireTime time.Duration) (string, error) {
	myClaims := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "xHHx",
		},
	}
	token := jwt.NewWithClaims(getSignedMethod(signed), myClaims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseToken(tokenString string, secretKey string) (*MyClaims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证Claims
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
