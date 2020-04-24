package Message

import (
	"elearn100/Model/Message"
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func List(c *gin.Context)  {
	c.HTML(e.SUCCESS,"message/message.html",gin.H{
		"title":"留言列表",
	})
}

func ListData(c *gin.Context)  {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Message.GetMessages(page)
	data["count"] =  e.GetPageNum(Message.GetMessageTotal())
	c.JSON(e.SUCCESS,gin.H{
		"code" : e.SUCCESS,
		"msg" : "留言列表",
		"data" : data,
	})
}