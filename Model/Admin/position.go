package Admin

import (
	"elearn100/Database"
)

//职位表
type SysAdminPosition struct {
	Database.Model

	PositionName string `json:"position_name" gorm:"type:varchar(30); not null; comment:'职位名称'" binding:"required"`
	DepartmentId int64 `json:"department_id" gorm:"not null; default:'0'; comment:'归属部门' " binding:"required"`
	Desc string `json:"desc" gorm:"not null; comment:'职位描述' "`
	Powerid int64 `json:"powerid" gorm:"default:'0'; comment:'职位权限' " binding:"required"`
	Status int64 `json:"status" gorm:"default:'1'; comment:'职位状态 1 正常'" binding:"required"`
	CityId int64 `json:"city_id" gorm:"not null; default '0'; comment:'城市ID' "`
}
