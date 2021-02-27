package Single

import (
	db "elearn100/Database"
	"elearn100/Model/Nav"
	"elearn100/Pkg/setting"
	"fmt"
)

type Single struct {
	db.Model

	Navs Nav.Nav `json:"nav" gorm:"FOREIGNKEY:NavId;ASSOCIATION_FOREIGNKEY:ID"`

	Name       string `json:"name" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	Content    string `json:"content" gorm:"type:text;default '';comment:'内容'"`
	NavId      int    `json:"nav_id" gorm:"default '';comment:'栏目ID'"`
	ThumbImg   string `json:"thumb_img" gorm:"not null;default '';comment:'缩率图'"`
	Summary    string `json:"summary" gorm:"type:varchar(255);not null;default '';comment:'摘要'"`
	Tag        string `json:"tag" gorm:"type:varchar(100);not null;default '';comment:'标签'"`
	ClientType int    `json:"client_type" gorm:"type:int(2);not null;default '1';comment:'1PC 2移动'"`
}

// @Summer 新增内容
func AddSingle(single Single) bool {
	if single := db.Db.Create(&single); single.Error != nil {
		fmt.Print("添加文章失败", single)
		return false
	}
	return true
}

func EditSingle(id int, single Single) bool {
	edit := db.Db.Model(&Single{}).Where("id = ?", id).Updates(single)
	if edit.Error != nil {
		fmt.Print("编辑文章失败", edit)
		return false
	}
	return true
}

// @Summer 获取所有文章
func GetSingles(page int, data interface{}) (singles []Single) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	fmt.Println("当前页数", page)
	db.Db.Preload("Navs").Where(data).Offset(offset).Limit(setting.PageSize).Order("id desc").Find(&singles)
	return
}

// @Summer 获取单篇文章
func GetSingle(id int) (single Single) {
	db.Db.Preload("Navs").Where("id = ?", id).First(&single)
	return
}

// @Summer 统计
func GetSingleTotal() (count int) {
	db.Db.Model(&Single{}).Count(&count)
	return
}

type SingleData struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	Content    string `json:"content" gorm:"type:text;default '';comment:'内容'"`
	NavId      int    `json:"nav_id" gorm:"default '';comment:'栏目ID'"`
	ThumbImg   string `json:"thumb_img" gorm:"not null;default '';comment:'缩率图'"`
	Summary    string `json:"summary" gorm:"type:varchar(255);not null;default '';comment:'摘要'"`
	Tag        string `json:"tag" gorm:"type:varchar(100);not null;default '';comment:'标签'"`
	ClientType int    `json:"client_type" gorm:"type:int(2);not null;default '1';comment:'1PC 2移动'"`
}

// @Summer 获取tag
func GetSingleByOne(navId, clientType int, tag string) (singles SingleData) {
	db.Db.Table("single").Where("nav_id = ? and client_type = ? and tag = ? ", navId, clientType, tag).Scan(&singles)
	return
}

// @Desc 获取tag
// @Param navId int 导航ID
// @Param clientType int 客户端类型
// @Param tag string tag标签
func GetAllSingle(navId, clientType int, tag string) (singles []Single) {
	db.Db.Where("nav_id = ? and client_type = ? and tag = ? ", navId, clientType, tag).Find(&singles)
	return
}
