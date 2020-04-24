package Message

import (
	"elearn100/Database"
	"elearn100/pkg/setting"
	"github.com/jinzhu/gorm"
	"time"
)

// 留言表
type Message struct {
	Database.Model

	Mname string `json:"mname" gorm:"type:varchar(100);not null; default ''; comment:'留言姓名' "`
	Area string `json:"area" gorm:"type:varchar(100);not null; default ''; comment:'区域' "`
	Tel string `json:"tel" gorm:"type:varchar(20);not null; default ''; comment:'留言电话' "`
	Content string `json:"content" gorm:"type:text;not null; default ''; comment:'留言内容' "`
	Com string `json:"com" gorm:"type:varchar(190);not null; default ''; comment:'留言来源页' "`
	Client string `json:"client" gorm:"type:varchar(190);not null; default ''; comment:'客户端' "`
	Ip string `json:"ip" gorm:"type:varchar(50);not null; default ''; comment:'ip地址' "`
	Channel string `json:"channel" gorm:"type:varchar(50);not null; default ''; comment:'留言板块' "`
}

// @Summer 留言总数
func GetMessageTotal() (count int) {
	Database.Db.Model(&Message{}).Count(&count)
	return
}

// @Summer 留言列表
// @Param int page 当前页
func GetMessages(page int) (messages []Message)  {
	offset :=0
	if page >= 1 {
		offset =  (page-1)*setting.PageSize
	}
	Database.Db.Offset(offset).Limit(setting.PageSize).Find(&messages)
	return
}

// @Summer 插入
func (message *Message) BeforeCreate(score *gorm.Scope) error  {
	score.SetColumn("CreatedAt",time.Now().Format("2006:01:02 15:04:05"))
	return nil
}

// @Summer 更新
func (message *Message) BeforeUpdate(score *gorm.Scope) error {
	score.SetColumn("UpdatedAt",time.Now().Format("2006:01:02 15:04:05"))
	return nil
}