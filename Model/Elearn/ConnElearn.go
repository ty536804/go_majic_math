package Elearn

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var elearnDb *gorm.DB

func init() {
	var err error
	dbName := "elearn100"
	user := "root"
	password := "elearnedu"
	host := "39.98.213.147:3306"
	elearnDb, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		fmt.Printf("elearn100 connected faild:%s", err)
	}
	fmt.Println("elelarn100 connect success")

	elearnDb.SingularTable(true)
	elearnDb.LogMode(true)
	elearnDb.DB().SetMaxIdleConns(10)
	elearnDb.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer elearnDb.Close()
}
