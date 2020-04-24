package Controller

import (
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context)  {
	c.HTML(e.SUCCESS,"index/index.html",gin.H{
		"title":"首页",
	})
}