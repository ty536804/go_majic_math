package jwt

import (
	"elearn100/Services"
	"elearn100/pkg/e"
	"elearn100/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
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
