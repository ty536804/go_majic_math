package Services

import (
	"elearn100/Model/Nav"
	"elearn100/Pkg/e"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Desc 获取导航列表
func RedisGetNavList() []Nav.Nav {
	conn := e.PoolConnect()
	defer conn.Close()

	var navs []Nav.Nav

	if exists, _ := redis.Bool(conn.Do("exists", e.MENUKey)); exists {
		values, _ := redis.Values(conn.Do("lrange", e.MENUKey, 0, -1))
		var nav Nav.Nav
		for _, v := range values {
			if err := json.Unmarshal(v.([]byte), &nav); err == nil {
				navs = append(navs, nav)
			}
		}
	} else {
		var navWhere = make(map[string]interface{})
		navWhere["is_show"] = 1
		navs = Nav.Navs(navWhere)
		for _, v := range navs {
			if jsonStr, err := json.Marshal(v); err == nil {
				conn.Do("rpush", e.MENUKey, jsonStr)
			}
		}
		conn.Do("expire", e.MENUKey, e.VALIDTime)
	}
	return navs
}

// @Desc 获取一条导航列表
func GetNav(c *gin.Context) (navs Nav.Nav) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	return Nav.GetNav(id)
}

// @Desc 添加/编辑导航
func AddNav(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	Name := com.StrTo(c.PostForm("name")).String()
	BaseUrl := com.StrTo(c.PostForm("base_url")).String()
	IsShow := com.StrTo(c.PostForm("is_show")).MustInt64()
	id := com.StrTo(c.PostForm("id")).MustInt()

	if code, err := validNav(Name, BaseUrl, IsShow); code == e.ERROR {
		return code, err
	}

	nav := Nav.Nav{
		Name:    Name,
		BaseUrl: BaseUrl,
		IsShow:  IsShow,
	}

	isOK := false
	if id < 1 {
		isOK = Nav.AddNav(nav)
	} else {
		isOK = Nav.EditNav(id, nav)
	}
	if isOK {
		conn := e.PoolConnect()
		defer conn.Close()

		conn.Do("expire", e.MENUKey, -1)
		return e.SUCCESS, "操作失败"
	}
	return e.ERROR, "操作失败"
}

// @Desc 数据校验
func validNav(Name, BaseUrl string, IsShow int64) (int, string) {
	valid := validation.Validation{}
	valid.Required(Name, "bname").Message("名称不能为空")
	valid.Required(BaseUrl, "base_url").Message("跳转地址不能为空")
	valid.Required(IsShow, "is_show").Message("是否展示必须选择")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// 获取缓存中的导航
func GetMenu() (navs []Nav.Nav) {
	return RedisGetNavList()
}
