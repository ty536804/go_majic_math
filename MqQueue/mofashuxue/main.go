package main

import (
	"elearn100/MqQueue/Mq"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var magicDb *gorm.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = "mysql"
	dbName = "mofashuxue"
	user = "root"
	password = "123456"
	host = "127.0.0.1:3306"

	magicDb, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		fmt.Println("connected failed:", err)
	}

	magicDb.SingularTable(true)
	magicDb.LogMode(true)
	magicDb.DB().SetMaxIdleConns(20)
	magicDb.DB().SetMaxOpenConns(100)
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

type Message struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Mname     string    `json:"mname" gorm:"type:varchar(100);not null; default ''; comment:'留言姓名' "`
	Area      string    `json:"area" gorm:"type:varchar(100);not null; default ''; comment:'区域' "`
	Tel       string    `json:"tel" gorm:"type:varchar(20);not null; default ''; comment:'留言电话' "`
	Client    string    `json:"client" gorm:"type:varchar(190);not null; default ''; comment:'客户端' "`
	Ip        string    `json:"ip" gorm:"type:varchar(50);not null; default ''; comment:'ip地址' "`
	VisitUuid string    `json:"visit_uuid" gorm:"type:varchar(32);not null; default ''; comment:'用户ID' "`
	MsgType   int       `json:"msg_type" gorm:"type:not null; default '0'; comment:'1 魔法数学 2布卢卡斯' "`
}

func main() {
	Mq.ConsumeEx("mofashuxue", "fanout", "", AddMessage)
}

// @Summer添加留言
func AddMessage(s string) {
	var msg Info
	err := json.Unmarshal([]byte(s), &msg)
	if err == nil {
		result := magicDb.Create(&Message{
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

func CloseDB() {
	defer magicDb.Close()
}
