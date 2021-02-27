package SmsServices

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var magicDb *gorm.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = "mysql"
	dbName = "mofashuxue"
	user = "root"
	password = "123456"
	host = "127.0.0.1:3306"

	magicDb, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		fmt.Println("connected failed:", err)
	}

	magicDb.SingularTable(true)
	magicDb.LogMode(true)
	magicDb.DB().SetMaxIdleConns(20)
	magicDb.DB().SetMaxOpenConns(100)
}

type SysSmsConfig struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `json:"account" gorm:"not null; default:''; comment:'APIID" binding:"required"`
	Password string `json:"password" gorm:"not null;default:''; comment:'APIKEY" binding:"required"`
	Url      string `json:"url" gorm:"not null;default:''; comment:'url'" binding:"required"`
}

// @Summer 获取短信配置
func GetSmsConfig() (smsConfig SysSmsConfig) {
	magicDb.First(&smsConfig)
	return
}

// @Summer 调用第三方
func SendSms(mobile, msg string) {
	v := url.Values{}

	smsConfig := GetSmsConfig()
	_account := smsConfig.Account
	_password := smsConfig.Password
	_url := smsConfig.Url

	v.Set("account", _account)
	v.Set("password", _password)
	v.Set("mobile", mobile)
	v.Set("content", msg)
	fmt.Println(msg, mobile, _account, _password)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, err := http.NewRequest("POST", _url, body)
	if err != nil {
		fmt.Println("post失败:", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("看下发送的结构 %+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	if err == nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	} else {
		fmt.Println("发送失败:", err)
	}
}
