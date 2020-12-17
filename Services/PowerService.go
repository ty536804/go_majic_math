package Services

import "elearn100/Model/Admin"

type Menus struct {
	Name  string      `json:"name"`
	Id    int         `json:"id"`
	Icon  string      `json:"icon"`
	Path  string      `json:"path"`
	Child interface{} `json:"child"`
}

// @Summer 获取导航
func GetPower() []interface{} {
	var menuData []interface{}
	parenPowerList := Admin.GetParentPower(0)
	for _, item := range parenPowerList {
		menuData = append(menuData, &Menus{
			Name:  item.PowerName,
			Id:    item.Id,
			Icon:  item.Icon,
			Path:  item.Path,
			Child: Admin.GetChildPower(item.Id),
		})
	}
	return menuData
}
