package Campus

import (
	db "elearn100/Database"
	"elearn100/Pkg/setting"
	"fmt"
	"log"
)

type Campus struct {
	db.Model

	SchoolName   string `json:"school_name" gorm:"type:varchar(190);not null;default '';comment:'学习名称'"`
	SchoolTel    string `json:"school_tel" gorm:"type:varchar(20);not null default 0;comment:'学校电话'"`
	WorkerTime   string `json:"worker_time" gorm:"type:varchar(50);not null; default '';comment:'工作时间'"`
	Address      string `json:"address" gorm:"type:varchar(190);not null;default '';comment:'学校地址'"`
	SchoolImg    string `json:"school_img" gorm:"type:varchar(190);not null;default '';comment:'学校封面照片'"`
	Province     int    `json:"province" gorm:"comment:'省'"`
	ProvinceName string `json:"province_name" gorm:"type:varchar(190);not null;default '';comment:'省名称'"`
	City         string `json:"city" gorm:"type:varchar(190);not null;default '0';comment:'市'"`
	Area         string `json:"area" gorm:"type:varchar(190);not null;default '0';comment:'区'"`
	IsShow       int    `json:"is_show" gorm:"default '0';comment:'是否显示 0否 1是'"`
	IsDel        int    `json:"is_del" gorm:"default '0';comment:'是否删除 0否 1是'"`
}

// @Summer 新增校区
func AddCampus(data map[string]interface{}) bool {
	err := db.Db.Create(&Campus{
		SchoolName:   data["school_name"].(string),
		SchoolTel:    data["school_tel"].(string),
		WorkerTime:   data["worker_time"].(string),
		Address:      data["address"].(string),
		SchoolImg:    data["school_img"].(string),
		ProvinceName: data["province_name"].(string),
		Province:     data["province"].(int),
		IsShow:       data["is_show"].(int),
		IsDel:        '0',
	})
	if err.Error != nil {
		log.Printf("添加校区失败,%v", err)
		return false
	}
	return true
}

// @Summer 编辑校区
func EditCampus(id int, data interface{}) bool {
	edit := db.Db.Model(&Campus{}).Where("id = ?", id).Update(data)
	if edit.Error != nil {

		fmt.Print("编辑校区错误:", edit)
		return false
	}
	return true
}

// @Summer 获取校区列表
func GetCampus(page int, where interface{}) (campuses []Campus) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Where(where).Offset(offset).Limit(setting.PageSize).Find(&campuses)
	return
}

// @Summer 统计校区数量
func CountCampus(where interface{}) (count int) {
	db.Db.Where(where).Model(&Campus{}).Count(&count)
	return
}

// @Summer 获取校区详情
func DetailCampus(id int) (campus Campus) {
	db.Db.Where("id = ?", id).First(&campus)
	return
}

// @Summer 分省统计
type SubUser struct {
	CProvince int    `json:"c_province"`
	Name      string `json:"name"`
}

// @Summer 分组
func GroupCampus() (subUser []SubUser) {
	db.Db.Raw("SELECT COUNT(province_name) AS c_province, province_name AS name FROM campus").Group("name").Scan(&subUser)
	return subUser
}

// @Summer 删除校区
func DelCampus(id int) bool {
	err := db.Db.Delete(&Campus{}, "id = ?", id)
	if err.Error != nil {
		fmt.Print("删除校区错误:", err)
		return false
	}
	return true
}
