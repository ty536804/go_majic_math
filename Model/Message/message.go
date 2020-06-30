package Message

import (
	db "elearn100/Database"
	"elearn100/Pkg/setting"
	"fmt"
)

// 留言表
type Message struct {
	db.Model

	Mname   string `json:"mname" gorm:"type:varchar(100);not null; default ''; comment:'留言姓名' "`
	Area    string `json:"area" gorm:"type:varchar(100);not null; default ''; comment:'区域' "`
	Tel     string `json:"tel" gorm:"type:varchar(20);not null; default ''; comment:'留言电话' "`
	Content string `json:"content" gorm:"type:text;not null; default ''; comment:'留言内容' "`
	Com     string `json:"com" gorm:"type:varchar(190);not null; default ''; comment:'留言来源页' "`
	Client  string `json:"client" gorm:"type:varchar(190);not null; default ''; comment:'客户端' "`
	Ip      string `json:"ip" gorm:"type:varchar(50);not null; default ''; comment:'ip地址' "`
	Channel int    `json:"channel" gorm:"type:varchar(50);not null; default ''; comment:'留言板块' "`
}

// @Summer 留言总数
func GetMessageTotal() (count int) {
	db.Db.Model(&Message{}).Count(&count)
	return
}

// @Summer 留言列表
// @Param int page 当前页
func GetMessages(page int) (messages []Message) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Offset(offset).Limit(setting.PageSize).Find(&messages)
	return
}

// @Summer添加留言
func AddMessage(data map[string]interface{}) bool {
	result := db.Db.Create(&Message{
		Mname:   data["mname"].(string),
		Area:    data["area"].(string),
		Tel:     data["tel"].(string),
		Content: data["content"].(string),
		Com:     data["com"].(string),
		Client:  data["client"].(string),
		Ip:      data["ip"].(string),
		Channel: data["channel"].(int),
	})

	if result.Error != nil {
		fmt.Print("添加留言失败", result)
		return false
	}
	return true
}
