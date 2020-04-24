package Admin

import "elearn100/Database"

//部门表
type SysAdminDepartment struct {
	Database.Model

	DpName string `json:"dp_name" gorm:"type:varchar(30); not null;default:''; comment:'部门名称'" binding:"required"`
	ParentId int64 `json:"parent_id" gorm:"not null; default:'0'; comment:'父部门' " binding:"required"`
	RootId int64 `json:"root_id" gorm:"not null; default:'0'; comment:'根部门'" binding:"required"`
	Level int64		`json:"level" gorm:"not null;default 0; comment:'部门登记'" binding:"required"`
	Path string `json:"path" gorm:"type:varchar(190);not null; default:'|';comment:'部门归属'" binding:"required"`
	Powerid string `json:"powerid" gorm:"type:text;not null; comment:'部门权限'"`
	Status int64 `json:"status" gorm:"not null;default:'1';comment:'部门状态 1 正常'" binding:"required"`
}
