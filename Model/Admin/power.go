package Admin

import (
	"elearn100/Database"
)

//权限
type SysAdminPower struct {
	Database.Model

	Pname    string `json:"pname" gorm:"type:varchar(50); not null; default ''; comment:'权限名称' "`
	Ptype    int64  `json:"ptype" gorm:"not null;default:'1'; unique; comment:'1 左侧菜单 2顶部菜单' "`
	Icon     string `json:"icon" gorm:"varchar(50); not null;default ''; comment:'权限ICON样式名称' "`
	Desc     string `json:"desc" gorm:"varchar(50); not null; default ''; comment:'权限描述' "`
	Purl     string `json:"purl" gorm:"varchar(100;); not null; default:''; comment:'权限地址' "`
	ParentId int64  `json:"parent_id" gorm:"not null; default:'0'; comment:'上级地址' "`
	Pindex   int64  `json:"pindex" gorm:"not null;unique; default:'0'; comment:'显示排序' "`
	Status   int64  `json:"status" gorm:"not null; default:1; comment:'状态 1 显示 0不显示'" binding:"required"`
}
