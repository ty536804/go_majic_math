package Commands

import (
	"elearn100/Model/Admin"
)

func init() {
	AddUser()
}
func AddUser() {
	if !Admin.ExistsByLoginName("admin") {
		user := Admin.SysAdminUser{
			LoginName:    "admin",
			NickName:     "admin",
			Email:        "admin@126.com",
			Tel:          "",
			Pwd:          Admin.Md5Pwd("admin"),
			Avatar:       "",
			DepartmentId: 0,
			PositionId:   "10000",
			CityId:       "10000",
			Statues:      int64(1),
		}
		Admin.AddUser(user)
	}
}
