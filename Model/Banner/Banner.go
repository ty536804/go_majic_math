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
	Info       string    `json:"info" gorm:"type:varchar(255);not null;default '';comment:'备注'"`
	Tag        string    `json:"tag" gorm:"type:varchar(190);not null;default '';comment:'标签'"`
	Type       int       `json:"type" gorm:"not null;default '1';comment:'1 PC 2 WAP'"`
	Sort       int       `json:"sort" gorm:"not null;default '0';comment:'排序'"`
}

// @Summer 添加banner
func AddBanner(banner Banner) bool {
	err := db.Db.Create(&banner)
	if err.Error != nil {
		log.Printf("添加banner失败,%v", err)
		return false
	}
	return true
}

// @Summer 编辑banner
func EditBanner(id int, banner Banner) bool {
	edit := db.Db.Model(&Banner{}).Where("id = ?", id).Updates(banner)
	if edit.Error != nil {
		fmt.Print("编辑banner错误:", edit)
		return false
	}
	return true
}

type BannerData struct {
	ID         int    `json:"id"`
	Province   string `json:"province" gorm:"type:varchar(190);not null;default '0';comment:'省'"`
	City       string `json:"city" gorm:"type:varchar(190);not null;default '0';comment:'市'"`
	Area       string `json:"area" gorm:"type:varchar(190);not null;default '0';comment:'区'"`
	Bname      string `json:"bname" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	Bposition  int    `json:"bposition" gorm:"index;comment:'位置'"`
	Imgurl     string `json:"imgurl" gorm:"type:varchar(190);not null;default '';comment:'图片地址'"`
	TargetLink string `json:"target_link" gorm:"type:varchar(190);not null;default '';comment:'跳转链接'"`
	IsShow     int    `json:"is_show" gorm:"default '1';comment:'状态 1显示 2隐藏'"`
	ImageSize  string `json:"image_size" gorm:"type:varchar(190);not null;default '';comment:'图片大小 长*高*宽'"`
	Info       string `json:"info" gorm:"type:varchar(255);not null;default '';comment:'备注'"`
	Tag        string `json:"tag" gorm:"type:varchar(190);not null;default '';comment:'标签'"`
	Type       int    `json:"type" gorm:"not null;default '1';comment:'1 PC 2 WAP'"`
	Sort       int    `json:"sort" gorm:"not null;default '0';comment:'排序'"`
}

// @Desc 获取一张banner图片
// @Param bPosition int 导航ID
// @Param clientType int 客户端类型
// @Param tag string 标签
func GetOneBanner(bPosition, clientType int, tag string) (banner BannerData) {
	db.Db.Table("banner").Where("bposition = ? and type = ? and tag =? ", bPosition, clientType, tag).Scan(&banner)
	return
}

// @Summer获取所有banner
func GetBanners(page int) (banner []Banner) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Preload("Navs").Offset(offset).Limit(setting.PageSize).Order("id desc").Find(&banner)
	return
}

// @Desc 统计图片总数
func GetBannerTotal() (count int) {
	db.Db.Model(&Banner{}).Count(&count)
	return
}

// @Desc 获取图片列表
// @Param id int id 主键ID
func GetBanner(id int) (banner Banner) {
	db.Db.Preload("Navs").Where("id = ?", id).First(&banner)
	return
}

// @Desc 获取所有banner
// @Param posi int 栏目ID
// @Param clientType int 客户端类型
// @Param tag string 标签
func GetBannerData(posi, clientType int, tag string) (banner []Banner) {
	db.Db.Where("bposition = ? and type = ? and tag = ? and is_show=1", posi, clientType, tag).Order("sort desc").Find(&banner)
	return
}

// @Desc 删除banner
// @Param id int 主键ID
func DelBanner(id int) bool {
	if id < 1 {
		return false
	}
	err := db.Db.Delete(&Banner{}, "id =? ", id)
	if err.Error != nil {
		log.Printf("删除banner失败,%v", err)
		return false
	}
	return true
}

// @Summer通过描述获取图片
func GetBannerList(info string) (banner []Banner) {
	db.Db.Where("info = ?", info).Find(&banner)
	return
}
