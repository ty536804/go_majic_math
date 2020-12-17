package Message

import (
	"elearn100/Model/Message"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 留言列表
func ListData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Message.GetMessages(page)
	data["count"] = Message.GetMessageTotal()
	data["size"] = setting.PageSize
	e.Success(c, "留言列表", data)
}

// @Summer 添加留言
func AddMessage(c *gin.Context) {
	_, msg := Services.AddMessage(c)
	e.Success(c, msg, "")
}
