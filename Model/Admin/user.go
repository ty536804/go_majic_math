package Admin

import (
	"crypto/md5"
	db "elearn100/Database"
	"elearn100/Pkg/setting"
	"encoding/hex"
	"errors"
	"log"
)

//管理员表
type SysAdminUser struct {
	db.Model

	LoginName    string `json:"login_name" gorm:"type:varchar(100);not null;unique;comment:'账号'" binding:"required"`
	NickName     string `json:"nick_name" gorm:"type:varchar(100);not null;default'';comment:'昵称'"`
	Email        string `json:"email" gorm:"type:varchar(100);not null;default '';comment:'邮箱'"`
	Tel          string `json:"tel" gorm:"type:char(20);not null;default '';comment:'电话'"`
	Pwd          string `json:"pwd" gorm:"type:char(32);not null;default '';unique;comment:'密码'" binding:"required"`
	Avatar       string `json:"avatar" gorm:"type:varchar(100);comment:'头像'"`
	DepartmentId int64  `json:"department_id" gorm:"not null;default '';comment:'部门'"`
	PositionId   string `json:"position_id" gorm:"comments:'职位 角色'" binding:"required"`
	CityId       string `json:"city_id" gorm:"type:text;comments:'城市ID'"`
	Statues      int64  `json:"statues" gorm:"not null;default:'0';comment:'状态 1显示'" binding:"required"`
	Token        string `json:"token" gorm:"varchar(100);not null;default:'';comment:'token'"`
	projectId    int64  `json:"project_id" gorm:"not null;default:'0';comment:'归属项目 0系统'"`
}

// @Summary 通过用户ID判断用户是否存在
// @param id int 管理员ID
func GetAdmin(id int) (admins SysAdminUser) {
	db.Db.First(&admins, id)
	return
}

// @Summary 获取用户信息
// @param loginName string 用户名
// @param pwd string 密码
func GetUserInfo(loginName, pwd string) (err error, uuid int) {
	admin := SysAdminUser{}
	userInfo := db.Db.Where("login_name = ?", loginName).First(&admin)
	if err := userInfo.Error; err != nil {
		return errors.New("用户不存在"), 0
	}

	if admin.Pwd != Md5Pwd(pwd) {
		return errors.New("密码不正确"), 0
	}
	return nil, admin.ID
}

// @Summary md5加密 16进制32位
// @param data string 密码
func Md5Pwd(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}

// @Summer 用户列表
func Users(page int) (users []SysAdminUser) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Offset(offset).Limit(setting.PageSize).Find(&users)
	return
}

// @Summer 用户总数
func GetUserTotal() (count int) {
	db.Db.Model(&SysAdminUser{}).Count(&count)
	return
}

// @Summer 判断当前账号是否已经注册
func ExistsByLoginName(loginName string) bool {
	var user SysAdminUser
	db.Db.Select("id").Where("login_name = ?", loginName).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

// @Summer 判断当前手机号码是否已注册
func ExistsByTel(tel string) bool {
	var user SysAdminUser
	db.Db.Select("id").Where("tel = ?", tel).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

// @Summer 添加用户
func AddUser(user SysAdminUser) bool {
	if err := db.Db.Create(&user); err.Error != nil {
		log.Printf("添加用户失败,%v", err)
		return false
	}
	return true
}

// @Summer 编辑用户
func EditUser(id int, user SysAdminUser) bool {
	if err := db.Db.Model(&SysAdminUser{}).Where("id = ?", id).Updates(user); err.Error != nil {
		log.Printf("编辑用户失败,%v", err)
		return false
	}
	return true
}

func GetLastUserId() (admin SysAdminUser) {
	db.Db.Select("id").Order("id desc").Limit(1).Find(&admin)
	return
}
