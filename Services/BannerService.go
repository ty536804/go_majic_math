package Services

import (
	"elearn100/Model/Banner"
	"elearn100/Pkg/e"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strings"
	"time"
)

// @Desc 添加/编辑图片
func AddBanner(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	bName := com.StrTo(c.PostForm("bname")).String()
	bPosition := com.StrTo(c.PostForm("bposition")).MustInt()
	imgUrl := com.StrTo(c.PostForm("imgurl")).String()
	targetLink := com.StrTo(c.PostForm("target_link")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	info := com.StrTo(c.PostForm("info")).String()
	sort := com.StrTo(c.PostForm("sort")).MustInt()
	tag := com.StrTo(c.PostForm("tag")).String()
	webType := com.StrTo(c.PostForm("type")).MustInt()

	if strings.HasPrefix(imgUrl, "/static/upload/") {
		imgUrl = strings.Replace(imgUrl, "/static/upload/", "", -1)
	}

	if code, err := validBanner(bName, imgUrl, bPosition, isShow); code == e.ERROR {
		return code, err
	}

	startTime := time.Now().Add(100 * time.Hour)
	banner := Banner.Banner{
		TargetLink: targetLink,
		Info:       info,
		Bname:      bName,
		Bposition:  bPosition,
		Imgurl:     imgUrl,
		IsShow:     isShow,
		Sort:       sort,
		Tag:        tag,
		Type:       webType,
		Province:   "10000",
		City:       "0",
		Area:       "0",
		BeginTime:  startTime,
		EndTime:    startTime,
	}
	isOK := false
	if id < 1 {
		isOK = Banner.AddBanner(banner)
	} else {
		isOK = Banner.EditBanner(id, banner)
	}
	if isOK {
		return e.ReSuccess()
	}
	return e.ReSuccess()
}

// @Desc 数据验证
func validBanner(bName, imgUrl string, bPosition, isShow int) (code int, msg string) {
	valid := validation.Validation{}
	valid.Required(bName, "bname").Message("名称不能为空")
	valid.Required(bPosition, "bposition").Message("展示位置必须选择")
	valid.Required(imgUrl, "imgurl").Message("上传图片")
	valid.Required(isShow, "is_show").Message("状态必须选择")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Desc 删除banner
func DelBanner(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	if Banner.DelBanner(id) {
		return e.ReSuccess()
	}
	return e.ReError()
}

func GetBanner(tit string) (banner []Banner.Banner) {
	return Banner.GetBannerList(tit)
}

// @Desc 缓存中获取内容
// @Param bPosition int 栏目ID
// @Param clientType int 客户端类型
// @Param tag string 标签
// @Param redisKey string 键
func RedisGetBannerList(bPosition, clientType int, tag, redisKey string) []Banner.Banner {
	conn := e.PoolConnect()
	defer conn.Close()

	redisKey = e.REDISKey + redisKey
	var banners []Banner.Banner

	if isOk, _ := redis.Bool(conn.Do("exists", redisKey)); isOk {
		values, _ := redis.Values(conn.Do("lrange", redisKey, 0, -1))
		var banner Banner.Banner
		for _, v := range values {
			if err := json.Unmarshal(v.([]byte), &banner); err == nil {
				banners = append(banners, banner)
			}
		}
	} else {
		banners = Banner.GetBannerData(bPosition, clientType, tag)
		for _, v := range banners {
			if jsonStr, err := json.Marshal(v); err == nil {
				conn.Do("rpush", redisKey, jsonStr)
			}
		}
		conn.Do("expire", redisKey, e.VALIDTime)
	}
	return banners
}

// @Desc 缓存中获取一条内容
// @Param bPosition int 栏目ID
// @Param clientType int 客户端类型
// @Param tag string 标签
// @Param redisKey string 键
func RedisGetOneBanner(bPosition, clientType int, tag, redisKey string) Banner.BannerData {
	conn := e.PoolConnect()
	defer conn.Close()

	redisKey = e.REDISKey + redisKey
	var banner Banner.BannerData
	if exists, _ := redis.Bool(conn.Do("exists", redisKey)); exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		redis.ScanStruct(res, &banner)
	} else {
		banner = Banner.GetOneBanner(bPosition, clientType, tag)
		if banner.ID > 0 {
			if _, err := conn.Do("hmset", redis.Args{redisKey}.AddFlat(banner)...); err == nil {
				conn.Do("expire", redisKey, e.VALIDTime)
			}
		}
	}
	return banner
}
