package Services

import (
	"elearn100/Model/Site"
	"elearn100/Pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加/编辑站点信息
func AddSite(c *gin.Context) (code int, err string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return e.ReError()
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	siteTitle := com.StrTo(c.PostForm("site_title")).String()
	SiteDesc := com.StrTo(c.PostForm("site_desc")).String()
	SiteKeyboard := com.StrTo(c.PostForm("site_keyboard")).String()
	SiteCopyright := com.StrTo(c.PostForm("site_copyright")).String()
	SiteTel := com.StrTo(c.PostForm("site_tel")).String()
	LandLine := com.StrTo(c.PostForm("land_line")).String()
	ClientTel := com.StrTo(c.PostForm("client_tel")).String()
	SiteEmail := com.StrTo(c.PostForm("site_email")).String()
	SiteAddress := com.StrTo(c.PostForm("site_address")).String()
	RecordNumber := com.StrTo(c.PostForm("record_number")).String()
	adminTel := com.StrTo(c.PostForm("admin_tel")).String()

	if code, err := validSite(siteTitle, SiteDesc, SiteKeyboard, SiteCopyright, SiteEmail, SiteAddress); code == e.ERROR {
		return code, err
	}

	if err := e.ValidTel(SiteTel, LandLine, ClientTel); err {
		return e.ERROR, "电话联系方式，必须填写一项"
	}

	site := Site.Site{
		SiteTitle:     siteTitle,
		SiteDesc:      SiteDesc,
		SiteKeyboard:  SiteKeyboard,
		SiteCopyright: SiteCopyright,
		SiteTel:       SiteTel,
		SiteEmail:     SiteEmail,
		SiteAddress:   SiteAddress,
		LandLine:      LandLine,
		ClientTel:     ClientTel,
		RecordNumber:  RecordNumber,
		AdminTel:      adminTel,
	}

	isOk := false
	if id < 1 {
		isOk = Site.AddSite(site)
	} else {
		isOk = Site.EditSite(id, site)
	}
	if isOk {
		return e.ReSuccess()
	}
	return e.ReError()
}

// @Desc 数据验证
func validSite(siteTitle, SiteDesc, SiteKeyboard, SiteCopyright, SiteEmail, SiteAddress string) (int, string) {
	valid := validation.Validation{}
	valid.Required(siteTitle, "site_title").Message("网站标题不能为空")
	valid.Required(SiteDesc, "site_desc").Message("网站描述不能为空")
	valid.Required(SiteKeyboard, "site_keyboard").Message("关键字不能为空")
	valid.Required(SiteCopyright, "site_copyright").Message("版权不能为空")
	valid.Required(SiteEmail, "site_email").Message("邮箱不能为空")
	valid.Required(SiteAddress, "site_address").Message("地址不能为空")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Desc 获取站点信息
func GetSite() Site.WebSite {
	conn := e.PoolConnect()
	defer conn.Close()

	redisKey := e.REDISKey + "site"
	var site Site.WebSite

	if exists, _ := redis.Bool(conn.Do("exists", redisKey)); exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		_ = redis.ScanStruct(res, &site)
	} else {
		site = Site.GetWebSite()
		conn.Do("hmset", redis.Args{redisKey}.AddFlat(site)...)
	}
	return site
}
