package Site

import (
	db "elearn100/Database"
	"fmt"
)

type Site struct {
	db.Model

	SiteTitle     string `json:"site_title" gorm:"type:varchar(190);not null;default '';comment:'网站标题'"`
	SiteDesc      string `json:"site_desc" gorm:"type:text;not null;default '';comment:'网站描述'"`
	SiteKeyboard  string `json:"site_keyboard" gorm:"type:text;not null;default '';comment:'网站关键字'"`
	SiteCopyright string `json:"site_copyright" gorm:"type:varchar(190);not null;default '';comment:'版权'"`
	SiteTel       string `json:"site_tel" gorm:"type:varchar(20);not null;default '';comment:'电话'"`
	SiteEmail     string `json:"site_email" gorm:"type:varchar(50);not null;default '';comment:'邮箱'"`
	SiteAddress   string `json:"site_address" gorm:"type:varchar(100);not null;default '';comment:'地址'"`
	LandLine      string `json:"land_line" gorm:"type:varchar(50);not null;default '';comment:'座机'"`
	ClientTel     string `json:"client_tel" gorm:"type:varchar(50);not null;default '';comment:'400电话'"`
	RecordNumber  string `json:"record_number" gorm:"type:varchar(100);not null;default '';comment:'备案号'"`
	AdminTel      string `json:"admin_tel" gorm:"type:varchar(255);not null;default '';comment:'接收短信的管理员手机号码'"`
}

func GetSite() (site Site) {
	db.Db.First(&site)
	return
}

type WebSite struct {
	ID            int    `gorm:"primary_key" json:"id"`
	SiteTitle     string `json:"site_title" gorm:"type:varchar(190);not null;default '';comment:'网站标题'"`
	SiteDesc      string `json:"site_desc" gorm:"type:text;not null;default '';comment:'网站描述'"`
	SiteKeyboard  string `json:"site_keyboard" gorm:"type:text;not null;default '';comment:'网站关键字'"`
	SiteCopyright string `json:"site_copyright" gorm:"type:varchar(190);not null;default '';comment:'版权'"`
	SiteTel       string `json:"site_tel" gorm:"type:varchar(20);not null;default '';comment:'电话'"`
	SiteEmail     string `json:"site_email" gorm:"type:varchar(50);not null;default '';comment:'邮箱'"`
	SiteAddress   string `json:"site_address" gorm:"type:varchar(100);not null;default '';comment:'地址'"`
	LandLine      string `json:"land_line" gorm:"type:varchar(50);not null;default '';comment:'座机'"`
	ClientTel     string `json:"client_tel" gorm:"type:varchar(50);not null;default '';comment:'400电话'"`
	RecordNumber  string `json:"record_number" gorm:"type:varchar(100);not null;default '';comment:'备案号'"`
	AdminTel      string `json:"admin_tel" gorm:"type:varchar(255);not null;default '';comment:'接收短信的管理员手机号码'"`
}

func GetWebSite() (site WebSite) {
	db.Db.Table("site").Where("id = 1 ").Scan(&site)
	return
}

// @Summer网站信息添加
func AddSite(site Site) bool {
	if err := db.Db.Create(&site); err.Error != nil {
		fmt.Print("基础信息添加失败", err)
		return false
	}
	return true
}

// @Summer 编辑网站信息
func EditSite(id int, site Site) bool {
	if err := db.Db.Model(&Site{}).Where("id = ?", id).Updates(site); err.Error != nil {
		fmt.Print("基础信息编辑失败", err)
		return false
	}
	return true
}
