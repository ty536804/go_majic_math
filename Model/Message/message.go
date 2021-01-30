package Message

import (
	db "elearn100/Database"
	"elearn100/Pkg/setting"
	"encoding/json"
	"fmt"
)

// 留言表
type Message struct {
	db.Model

	Mname     string `json:"mname" gorm:"type:varchar(100);not null; default ''; comment:'留言姓名' "`
	Area      string `json:"area" gorm:"type:varchar(100);not null; default ''; comment:'区域' "`
	Tel       string `json:"tel" gorm:"type:varchar(20);not null; default ''; comment:'留言电话' "`
	Client    string `json:"client" gorm:"type:varchar(190);not null; default ''; comment:'客户端' "`
	Ip        string `json:"ip" gorm:"type:varchar(50);not null; default ''; comment:'ip地址' "`
	VisitUuid string `json:"visit_uuid" gorm:"type:varchar(32);not null; default ''; comment:'用户ID' "`
	MsgType   int    `json:"msg_type" gorm:"type:not null; default '0'; comment:'1 魔法数学 2布卢卡斯' "`
}

// @Summer 留言总数
func GetMessageTotal() (count int) {
	db.Db.Model(&Message{}).Count(&count)
	return
}

type ListMessage struct {
	Id           int    `json:"id"`
	Mname        string `json:"mname"`
	Area         string `json:"area"`
	Tel          string `json:"tel"`
	Client       string `json:"client"`
	Ip           string `json:"ip"`
	VisitUuid    string `json:"visit_uuid"`
	MsgType      int    `json:"msg_type"`
	FirstUrl     string `json:"first_url"`
	FromUrl      string `json:"from_url"`
	VisitHistory string `json:"visit_history"`
	CreatedAt    string `json:"created_at"`
}

// @Summer 留言列表
// @Param page int 当前页
func GetMessages(page int) (messages []ListMessage) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Raw("SELECT m.id,m.mname,m.area,m.tel,m.client,m.ip,m.visit_uuid,UNIX_TIMESTAMP(m.created_at) as created_at,"+
		"m.msg_type,v.first_url,v.from_url,h.visit_history"+
		" FROM message as m LEFT JOIN visit as v on v.uuid=m.visit_uuid"+
		" LEFT JOIN history as h on m.visit_uuid=h.uuid ORDER BY m.id DESC LIMIT ?,?", offset, setting.PageSize).
		Find(&messages)
	return
}

type Info struct {
	MName   string
	Area    string
	Tel     string
	Client  string
	Ip      string
	Uid     string
	MsgType int
}

// @Summer添加留言
func AddMessage(s string) {
	var msg Info
	err := json.Unmarshal([]byte(s), &msg)
	if err == nil {
		result := db.Db.Create(&Message{
			Mname:     msg.MName,
			Area:      msg.Area,
			Tel:       msg.Tel,
			Client:    msg.Client,
			Ip:        msg.Ip,
			VisitUuid: msg.Uid,
			MsgType:   msg.MsgType,
		})
		if result.Error != nil {
			fmt.Print("添加留言失败", result)
		}
	}
}

func GetTotalMessage(uid string, ftime, ltime string) (count int) {
	db.Db.Model(&Message{}).Where("ip = ? AND created_at >= ? AND created_at <= ?", uid, ftime, ltime).Count(&count)
	return
}
