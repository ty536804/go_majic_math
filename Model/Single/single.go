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

	Name     string `json:"name" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	Content  string `json:"content" gorm:"type:text;default '';comment:'内容'"`
	NavId    int    `json:"nav_id" gorm:"default '';comment:'栏目ID'"`
	ThumbImg string `json:"thumb_img" gorm:"not null;default '';comment:'缩率图'"`
	Summary  string `json:"summary" gorm:"type:varchar(255);not null;default '';comment:'摘要'"`
	Tag      string `json:"tag" gorm:"type:varchar(100);not null;default '';comment:'标签'"`
}

// @Summer 新增内容
func AddSingle(data map[string]interface{}) bool {
	single := db.Db.Create(&Single{
		Name:     data["name"].(string),
		Content:  data["content"].(string),
		NavId:    data["nav_id"].(int),
		ThumbImg: data["thumb_img"].(string),
		Summary:  data["summary"].(string),
	})
	if single.Error != nil {
		fmt.Print("添加文章失败", single)
		return false
	}
	return true
}

func EditSingle(id int, data interface{}) bool {
	edit := db.Db.Model(&Single{}).Where("id = ?", id).Update(data)
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
	db.Db.Preload("Navs").Where(data).Offset(offset).Limit(setting.PageSize).Find(&singles)
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

// @Summer 获取所有文章
func GetAllSingle(id int) (singles []Single) {
	db.Db.Where("nav_id =?", id).Find(&singles)
	return
}

// @Summer 获取tag
func GetTag(id int) (singles []Single) {
	db.Db.Select("name").Where("nav_id = ? ", id).Group("name").Find(&singles)
	return
}

// @Summer 通过tag获取内容
func GetCon(tit string) (singles []Single) {
	db.Db.Where("name = ? ", tit).Find(&singles)
	return
}
