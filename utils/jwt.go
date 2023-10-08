package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
	"zhongxite/t-template/models"
)

var key = []byte("zhongxite/gin_chat")

type MycustomClaims struct {
	Id   uint
	Name string
	jwt.RegisteredClaims
}

// 生成
func GenerateTokenUsingHs256(user *models.User) (string, error) {
	claims := &MycustomClaims{
		Id:   user.ID,
		Name: user.AccountName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "zhongxite",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解密
func ParseTokenHs256(token string) (*jwt.Token, string, error) {
	claims := &MycustomClaims{}
	tokenStr, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	id := strconv.Itoa(int(claims.Id))
	return tokenStr, id, err
}
