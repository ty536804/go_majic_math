package e

import (
	"elearn100/Pkg/setting"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Conn() (client *redis.Client) {
	Rclient := redis.NewClient(&redis.Options{
		Addr:     setting.RedisHost,
		Password: setting.RedisPwd, // no password set
		DB:       0,                // use default DB
	})
	return Rclient
}

func SetVal(key, val string) bool {
	redisCon := Conn()
	if err := redisCon.Set(key, val, time.Duration(3*time.Hour)); err.Err() != nil {
		fmt.Println("reids:", err)
		return false
	}
	return true
}

func SetMenuVal(key, val string) bool {
	redisCon := Conn()
	if err := redisCon.Set(key, val, time.Duration(720*time.Hour)); err.Err() != nil {
		fmt.Println("reids:", err)
		return false
	}
	return true
}

func GetVal(key string) (isOk bool, val string) {
	redisCon := Conn()
	if err := redisCon.Get(key); err.Err() != nil {
		return false, ""
	}
	return true, redisCon.Get(key).Val()
}

func DelVal(key string) bool {
	redisCon := Conn()
	if err := redisCon.Del(key); err.Err() != nil {
		return false
	}
	return true
}

// @Summer 微信token缓存
func SetAccessToken(val string) bool {
	redisCon := Conn()
	if err := redisCon.Set("access_token", val, time.Duration(2*time.Hour)); err.Err() != nil {
		fmt.Println("reids:", err)
		return false
	}
	return true
}

// @Summer 微信token缓存
func SetQiNiuToken(val string) bool {
	redisCon := Conn()
	if err := redisCon.Set("qiniu_token", val, time.Duration(8*time.Hour)); err.Err() != nil {
		fmt.Println("reids:", err)
		return false
	}
	return true
}
