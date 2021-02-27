package e

import (
	"elearn100/Pkg/setting"
	"github.com/garyburd/redigo/redis"
	"time"
)

//连接池连接
func PoolConnect() redis.Conn {
	pool := &redis.Pool{
		MaxIdle:     1,                 //最大空闲连接数
		MaxActive:   10,                //最大连接数
		IdleTimeout: 180 * time.Second, //超时时间
		Wait:        true,              //超过最大连接数之后的操作  等待还是报错   等待
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisHost)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	return pool.Get()
}
