package jwt

import (
	"elearn100/pkg/e"
	"elearn100/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		isOk, token:= e.GetVal("token")

		if !isOk {
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
			//e.Error(c,e.GetMsg(code),data)
			c.Redirect(http.StatusMovedPermanently, "/admin")
			c.Abort()
			return
		}
		c.Next()
	}
}