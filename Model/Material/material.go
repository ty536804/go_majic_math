package Material

import (
	db "elearn100/Database"
	"elearn100/Pkg/setting"
	"fmt"
)

// 视频
type Material struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"type:varchar(100);not null;default '';comment:'标题' "`
	VideoSrc  string `json:"video_src" gorm:"type:varchar(100);not null;default '';comment:'七牛视频地址' "`
	LocalSrc  string `json:"local_src" gorm:"type:varchar(200);not null;default '';comment:'本地视频地址' "`
	IsShow    int    `json:"is_show" gorm:"not null;default 0;comment:'是否展示 0展示 1禁止' "`
	IsHot     int    `json:"is_hot" gorm:"not null;default 0;comment:'排序 数字越高排序越靠前' "`
	Code      string `json:"code" gorm:"type:varchar(30);not null;default '';comment:'观看码' "`
	CreatedAt string `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

// @Summer 添加视频
func AddMaterial(data map[string]interface{}) bool {
	res := db.Db.Create(Material{
		Title:    data["title"].(string),
		VideoSrc: data["video_src"].(string),
		LocalSrc: data["local_src"].(string),
		IsShow:   data["is_show"].(int),
		Code:     data["code"].(string),
		IsHot:    data["is_hot"].(int),
	})
	if res.Error != nil {
		fmt.Println("视频添加失败:", res)
		return false
	}
	return true
}

// @Summer获取 编辑单个视频
func EditMaterial(id int, data map[string]interface{}) bool {
	res := db.Db.Model(&Material{}).Where("id = ?", id).Update(data)
	if res.Error != nil {
		fmt.Println("视频编辑错误", res)
		return false
	}
	return true
}

// @Summer 获取单个视频列表
func GetMaterial(id int) (m Material) {
	if id >= 1 {
		db.Db.Where("id = ?", id).Find(&m)
		return
	}
	return m
}

// @Summer 获取视频列表
func GetMaterials(page int, where map[string]interface{}) (m []Material) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * setting.PageSize
	db.Db.Limit(setting.PageSize).Offset(offset).Order("is_hot desc,updated_at desc").Where(where).Find(&m)
	return
}

// @Summer 统计视频数量
func GetTotalMaterial() (cont int) {
	db.Db.Model(&Material{}).Count(&cont)
	return
}
