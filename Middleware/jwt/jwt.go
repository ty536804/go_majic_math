package jwt

import (
	"elearn100/Pkg/e"
	"elearn100/Pkg/util"
	"elearn100/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		var code int

		code = e.SUCCESS
		isOk, token := e.GetVal("token")
		uuid, uOk := c.Request.Cookie("uuid")

		if uOk != nil && !isOk || len(uuid.Value) == 0 {
			Services.LogOut(c)
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != e.SUCCESS {
			fmt.Println("uuid:,token:", uuid, token)
			c.Redirect(http.StatusMovedPermanently, "/admin")
			c.Abort()
			return
		}
		c.Next()
	}
}
