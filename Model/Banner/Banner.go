package Banner

import (
	db "elearn100/Database"
	"elearn100/Model/Nav"
	"elearn100/Pkg/setting"
	"fmt"
	"log"
	"time"
)

type Banner struct {
	db.Model

	Navs Nav.Nav `json:"nav" gorm:"FOREIGNKEY:Bposition;ASSOCIATION_FOREIGNKEY:ID"`

	Province   string    `json:"province" gorm:"type:varchar(190);not null;default '0';comment:'省'"`
	City       string    `json:"city" gorm:"type:varchar(190);not null;default '0';comment:'市'"`
	Area       string    `json:"area" gorm:"type:varchar(190);not null;default '0';comment:'区'"`
	Bname      string    `json:"bname" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	Bposition  int       `json:"bposition" gorm:"index;comment:'位置'"`
	Imgurl     string    `json:"imgurl" gorm:"type:varchar(190);not null;default '';comment:'图片地址'"`
	TargetLink string    `json:"target_link" gorm:"type:varchar(190);not null;default '';comment:'跳转链接'"`
	BeginTime  time.Time `json:"begin_time" time_format:"2006-01-02 15:04:05" gorm:"default '';comment:'显示开始时间'"`
	EndTime    time.Time `json:"end_time" time_format:"2006-01-02 15:04:05" gorm:"default '';comment:'显示结束时间'"`
	IsShow     int       `json:"is_show" gorm:"default '1';comment:'状态 1显示 2隐藏'"`
	ImageSize  string    `json:"image_size" gorm:"type:varchar(190);not null;default '';comment:'图片大小 长*高*宽'"`
	Info       string    `json:"info" gorm:"type:varchar(190);not null;default '';comment:'备注'"`
}

// @Summer 添加banner
func AddBanner(data map[string]interface{}) bool {
	startTime := time.Now().Add(100 * time.Hour)
	err := db.Db.Create(&Banner{
		Province:   "10000",
		City:       "0",
		Area:       "0",
		Bname:      data["bname"].(string),
		Bposition:  data["bposition"].(int),
		Imgurl:     data["imgurl"].(string),
		Info:       data["info"].(string),
		TargetLink: data["target_link"].(string),
		IsShow:     data["is_show"].(int),
		BeginTime:  startTime,
		EndTime:    startTime,
	})

	if err.Error != nil {
		log.Printf("添加banner失败,%v", err)
		return false
	}
	return true
}

// @Summer 编辑banner
func EditBanner(id interface{}, data interface{}) bool {
	edit := db.Db.Model(&Banner{}).Where("id = ?", id).Update(data)
	if edit.Error != nil {
		fmt.Print("编辑banner错误:", edit)
		return false
	}
	return true
}

// @Summer获取所有banner
func GetBanners(page int) (banner []Banner) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Preload("Navs").Offset(offset).Limit(setting.PageSize).Find(&banner)
	return
}

// @Summer 统计图片总数
func GetBannerTotal() (count int) {
	db.Db.Model(&Banner{}).Count(&count)
	return
}

// @Summer 获取图片列表
func GetBanner(id int) (banner Banner) {
	db.Db.Preload("Navs").Where("id = ?", id).First(&banner)
	return
}

// @Summer获取所有banner
func GetBannerData(bposition int) (banner []Banner) {
	db.Db.Where("bposition = ?", bposition).Find(&banner)
	return
}
