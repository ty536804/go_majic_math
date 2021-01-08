package WeChat

import (
	db "elearn100/Database"
	"elearn100/Pkg/setting"
	"fmt"
)

// 微信公众号 关注用户表
type ChatUser struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Mobile    string `json:"mobile" gorm:"type:varchar(20);not null;default '';comment:'手机号' "`
	OpenId    string `json:"open_id" gorm:"type:varchar(100);not null;default '';comment:'openid' "`
	NickName  string `json:"nick_name" gorm:"type:varchar(100);not null;default '';comment:'昵称' "`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(100);not null;default '';comment:'头像' "`
	Gender    int    `json:"gender" gorm:"not null;default 0;comment:'性别 1时是男性，值为2时是女性，值为0时是未知' "`
	Province  string `json:"province" gorm:"type:varchar(100);not null;default '';comment:'省' "`
	City      string `json:"city" gorm:"type:varchar(100);not null;default '';comment:'市' "`
	Country   string `json:"country" gorm:"type:varchar(100);not null;default 0;comment:'国家' "`
	IsBack    int    `json:"is_back" gorm:"not null;default 0;comment:'是否拉黑 0否 1是'"`
	LookCode  string `json:"look_code" gorm:"type:varchar(100);not null;default '';comment:'观看视频密码"`
	Subscribe int    `json:"subscribe" gorm:"not null;default 0;comment:'用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息' "`
	CreatedAt string `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func AddUser(data map[string]interface{}) bool {
	res := db.Db.Create(&ChatUser{
		Mobile:    data["mobile"].(string),
		OpenId:    data["open_id"].(string),
		NickName:  data["nick_name"].(string),
		AvatarUrl: data["avatar_url"].(string),
		Gender:    data["gender"].(int),
		Province:  data["province"].(string),
		City:      data["city"].(string),
		Country:   data["country"].(string),
		Subscribe: data["subscribe"].(int),
		IsBack:    data["is_back"].(int),
		LookCode:  data["look_code"].(string),
	})
	if res.Error != nil {
		fmt.Println("添加用户失败")
		return false
	}
	return true
}

// @Summer 获取单个微信用户信息
func GetUser(id int, openid string) (u ChatUser) {
	db.Db.Where("id =? ", id).Or("open_id =? ", openid).Find(&u)
	return
}

// @Summer 编辑微信用户信息
func EditUser(id int, openid string, data map[string]interface{}) bool {
	res := db.Db.Model(&ChatUser{}).Where("id =? ", id).Or("open_id =? ", openid).Update(data)
	if res.Error != nil {
		fmt.Println("更新微信公众号信息失败", res)
		return false
	}
	return true
}

// @Summer 获取微信用户列表
func GetUsers(page int) (u []ChatUser) {
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * setting.PageSize
	db.Db.Limit(setting.PageSize).Offset(offset).Order("created_at desc").Find(&u)
	return
}

// @Summer 统计微信用户
func GetTotalUser() (count int) {
	db.Db.Model(&ChatUser{}).Count(&count)
	return
}
