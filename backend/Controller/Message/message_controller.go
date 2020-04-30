package Message

import (
	"elearn100/Model/Message"
	"elearn100/Services"
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func List(c *gin.Context) {
	c.HTML(e.SUCCESS, "message/message.html", gin.H{
		"title": "留言列表",
	})
}

// @Summer 留言列表
func ListData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Services.GetMessages(page)
	data["count"] = e.GetPageNum(Message.GetMessageTotal())
	e.Success(c, "留言列表", data)
}

// @Summer 添加留言
func AddMessage(c *gin.Context) {
	Services.AddMessage(c)
	data := make(map[string]interface{})
	e.Success(c, "留言成功", data)
}
