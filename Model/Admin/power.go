package Admin

import (
	db "elearn100/Database"
)

// @Summer 权限表
type SysAdminPower struct {
	Id        int    `json:"id" gorm:"primary_key:true;unique;"`
	PowerName string `json:"power_name" gorm:"comment:'权限名称' " `
	Level     int    `json:"level" gorm:"not null;default 0;comment:'级别'"`
	Pid       int    `json:"pid" gorm:"not null;default 0;comment:'父ID' " `
	Status    int    `json:"status" gorm:"not null;default 1;comment:'状态 1有效 0无效'" `
	Icon      string `json:"icon" gorm:"varchar(50);not null;default '';comment:'icon 图标'"`
	Path      string `json:"path" gorm:"varchar(100);not null; default ''; comment:'访问地址'"`
	Desc      string `json:"desc" gorm:"type:varchar(200); not null; default '';comment:'描述'" `

	db.Model
}

// @Summer 获取权限列表
func GetParentPower(level int) (power []SysAdminPower) {
	db.Db.Select("power_name, level, pid, id, status, icon, path").Where("pid = ? and status=1", level).Find(&power)
	return
}

// @Summer 获取权限列表
func GetChildPower(pid int) (power []SysAdminPower) {
	db.Db.Select("power_name, level, pid, id, status, icon, path").Where("pid = ? and status=1", pid).Find(&power)
	return
}
