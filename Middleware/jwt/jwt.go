package jwt

import (
	"elearn100/Pkg/e"
	"elearn100/Pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
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

		isOk := util.CheckLoginParam(c)
		if !isOk {
			e.Error(c, "缺少参数", "")
			c.Abort()
			return
		}
		if util.GetSignContent(c) == c.PostForm("sign") {
			e.Error(c, "非法签名", "")
			c.Abort()
			return
		}
		c.Next()
	}
}
