package Services

import (
	"elearn100/Model/Single"
	"elearn100/Pkg/e"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加文章
func AddSingle(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return e.ERROR, "操作失败"
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	name := com.StrTo(c.PostForm("name")).String()
	navId := com.StrTo(c.PostForm("nav_id")).MustInt()
	content := com.StrTo(c.PostForm("content")).String()
	thumbImg := com.StrTo(c.PostForm("thumb_img")).String()
	summary := com.StrTo(c.PostForm("summary")).String()
	tag := com.StrTo(c.PostForm("tag")).String()
	ClientType := com.StrTo(c.PostForm("client_type")).MustInt()

	if code, err := validSingle(name, content, tag, navId); code == e.ERROR {
		return code, err
	}

	single := Single.Single{
		Name:       name,
		Content:    content,
		NavId:      navId,
		ThumbImg:   thumbImg,
		Summary:    summary,
		Tag:        tag,
		ClientType: ClientType,
	}

	isOk := false
	if id < 1 {
		isOk = Single.AddSingle(single)
	} else {
		isOk = Single.EditSingle(id, single)
	}
	if isOk {
		return e.ReSuccess()
	}
	return e.ReError()
}

// #Desc数据验证
func validSingle(name, content, tag string, navId int) (int, string) {
	valid := validation.Validation{}
	valid.Required(name, "name").Message("标题不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(navId, "nav_id").Message("栏目不能为空")
	valid.Required(tag, "tag").Message("标签不能为空")

	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Desc 获取单页信息列表
// @Param navId int 导航ID
// @Param clientType int 客户端类型
// @Param tag string tag标签
// @Param redisKey string 缓存key
func RedisGetAllSingle(navId, clientType int, tag, redisKey string) []Single.Single {
	conn := e.PoolConnect()
	defer conn.Close()

	redisKey = e.REDISKey + redisKey
	var singles []Single.Single

	if exists, _ := redis.Bool(conn.Do("exists", redisKey)); exists {
		values, _ := redis.Values(conn.Do("lrange", redisKey, 0, -1))
		var single Single.Single
		for _, v := range values {
			if err := json.Unmarshal(v.([]byte), &single); err == nil {
				singles = append(singles, single)
			}
		}
	} else {
		singles = Single.GetAllSingle(navId, clientType, tag)
		for _, v := range singles {
			if jsonStr, err := json.Marshal(v); err == nil {
				conn.Do("rpush", redisKey, jsonStr)
			}
		}
		conn.Do("expire", redisKey, e.VALIDTime)
	}
	return singles
}

// @Desc 获取单页信息一条
// @Param navId int 导航ID
// @Param clientType int 客户端类型
// @Param tag string tag标签
// @Param redisKey string 缓存key
func RedisGetSingleByOne(navId, clientType int, tag, redisKey string) Single.SingleData {
	conn := e.PoolConnect()
	defer conn.Close()

	redisKey = e.REDISKey + redisKey
	var single Single.SingleData

	if exists, _ := redis.Bool(conn.Do("exists", redisKey)); exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		_ = redis.ScanStruct(res, &single)
	} else {
		single = Single.GetSingleByOne(navId, clientType, tag)
		if _, err := conn.Do("hmset", redis.Args{redisKey}.AddFlat(single)...); err == nil {
			conn.Do("expire", redisKey, e.VALIDTime)
		}
	}
	return single
}
