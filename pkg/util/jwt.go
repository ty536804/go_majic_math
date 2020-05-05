package util

import (
	"elearn100/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwrSecret)

type Claims struct {
	LoginName string `json:"login_name"`
	Pword     string `json:"pword"`
	jwt.StandardClaims
}

// @Summer 生成token
func GenerateToken(loginName, pwd string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		loginName,
		pwd,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "elearn100",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// @Summer 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
