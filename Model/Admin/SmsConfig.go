package Admin

import (
	db "elearn100/Database"
	"fmt"
)

type SysSmsConfig struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `json:"account" gorm:"not null; default:''; comment:'APIID" binding:"required"`
	Password string `json:"password" gorm:"not null;default:''; comment:'APIKEY" binding:"required"`
	Url      string `json:"url" gorm:"not null;default:''; comment:'url'" binding:"required"`
}

// 添加短信配置
func AddSms(data map[string]interface{}) bool {
	result := db.Db.Create(&SysSmsConfig{
		Account:  data["account"].(string),
		Password: data["password"].(string),
		Url:      data["url"].(string),
	})
	if result.Error != nil {
		fmt.Print("添加短信失败", result)
		return false
	}
	return true
}

// @Summer 编辑短信
func EditSms(id int, data interface{}) bool {
	edit := db.Db.Model(&SysSmsConfig{}).Where("id = ?", id).Update(data)
	if edit.Error != nil {
		fmt.Print("编辑短信失败", edit)
		return false
	}
	return true
}

// @Summer 获取短信配置
func GetSmsConfig() (smsConfig SysSmsConfig) {
	db.Db.First(&smsConfig)
	return
}
