package util

import (
	"elearn100/Model/Admin"
	"elearn100/Pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
	"time"
)

var (
	pubParam  = []string{"timestamp", "version", "client", "sign"}
	jwtSecret = []byte(setting.JwrSecret)
)

const (
	secre   = "brocaedu"
	version = "1.0"
)

type Claims struct {
	LoginName string `json:"login_name"`
	Pword     string `json:"pword"`
	jwt.StandardClaims
}

// @Summer 生成token
func GenerateToken(loginName, pwd string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(8 * time.Hour)
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

// @Summer 排序
func SortParam(c *gin.Context) string {
	sort.Strings(pubParam)
	newParam := ""
	for _, v := range pubParam {
		_defParam := strings.TrimSpace(c.PostForm(v))
		if _defParam == "" {
			continue
		}
		newParam += v + strings.TrimSpace(c.PostForm(v))
	}
	return newParam
}

// @Summer 加密
func GetSignContent(c *gin.Context) string {
	newParam := SortParam(c) + secre
	return Admin.Md5Pwd(newParam)
}

// @Summer 校验公共参数
func CheckLoginParam(c *gin.Context) (isOk bool) {
	for _, v := range pubParam {
		_, isOk := c.GetPostForm(v)
		if !isOk {
			return false
		}
	}
	return true
}
