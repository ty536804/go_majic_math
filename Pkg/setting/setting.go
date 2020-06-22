package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string //调试模式

	HTTPPort     int           //端口号
	ReadTimeout  time.Duration //读取时间
	WriteTimeout time.Duration //写入时间

	PageSize  int    //每页显示数量
	JwrSecret string //json web token

	RedisHost string
	RedisPwd  string

	Domain string
)

func init() {
	var err error
	Cfg, err = ini.Load("Conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'Conf/app.ini':%v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadRedis()
}

// 开启调试模式
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	Domain = Cfg.Section("").Key("HOST_NAME").MustString("127.0.0.0.1:8000")
}

// 加载服务器初始配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 每页显示数量
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app':%v", err)
	}
	JwrSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'Redis':%v", err)
	}
	RedisHost = sec.Key("Host").MustString("127.0.0.1:6379")
	RedisPwd = sec.Key("Password").MustString("")
}
