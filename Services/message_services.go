package Services

import (
	"elearn100/Database"
	"elearn100/Model/Message"
	"elearn100/pkg/setting"
)

var m Message.Message

func List(page int,maps ...interface{}) (message []Message.Message)  {
	offset := (page-1)*setting.PageSize
	Database.Db.Where(maps).Offset(offset).Limit(setting.PageSize).Find(&m)
	return
}
