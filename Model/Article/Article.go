package Article

import (
	db "elearn100/Database"
	"elearn100/pkg/setting"
	"fmt"
)

type Article struct {
	db.Model

	Title    string `json:"title" gorm:"type:varchar(190);not null;default '';comment:'标题'"`
	Summary  string `json:"summary" gorm:"type:varchar(255);not null;default '';comment:'摘要'"`
	ThumbImg string `json:"thumb_img" gorm:"type:varchar(190);not null;default '';comment:'缩率图'"`
	Admin    string `json:"admin" gorm:"type:varchar(190);not null;default '';comment:'发布者'"`
	Com      string `json:"com" gorm:"type:varchar(190);not null;default '';comment:'来源'"`
	IsShow   int    `json:"is_show" gorm:"not null;default '1';comment:'是否展示 1展示 2不展示'"`
	Content  string `json:"content" gorm:"type:text;not null;default '';comment:'内容'"`
	Hot      int    `json:"hot" gorm:"not null;default '2';comment:'是否热点 1是 2否'"`
	Sort     int    `json:"sort" gorm:"not null;default '0';comment:'优先级 数字越大，排名越前'"`
}

// @Summer 添加文章
func AddArticle(data map[string]interface{}) bool {
	article := db.Db.Create(&Article{
		Title:    data["title"].(string),
		Summary:  data["summary"].(string),
		ThumbImg: data["thumb_img"].(string),
		Admin:    data["admin"].(string),
		Com:      data["com"].(string),
		IsShow:   data["is_show"].(int),
		Content:  data["content"].(string),
		Hot:      data["hot"].(int),
		Sort:     data["sort"].(int),
	})

	if article.Error != nil {
		fmt.Print("添加文章失败", article)
		return false
	}
	return true
}

// @Summer 编辑文章
func EditArticle(id int, data interface{}) bool {
	edit := db.Db.Model(&Article{}).Where("id = ?", id).Update(data)
	if edit.Error != nil {
		fmt.Print("编辑文章失败", edit)
		return false
	}
	return true
}

// @Summer 获取所有文章
func GetArticles(page int, data interface{}) (articleS []Article) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Where(data).Offset(offset).Limit(setting.PageSize).Find(&articleS)
	return
}

// @Summer 获取单票文章
func GetArticle(id int) (article Article) {
	db.Db.Where("id = ?", id).First(&article)
	return
}

// @Summer 统计
func GetArticleTotal() (count int) {
	db.Db.Model(&Article{}).Count(&count)
	return
}
