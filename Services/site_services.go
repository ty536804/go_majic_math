package Services

import (
	"bytes"
	"elearn100/Model/Site"
	"elearn100/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"io/ioutil"
)

// @Summer 添加/编辑站点信息
func AddSite(c *gin.Context) (code int,err string)  {
	buf := make([]byte,3072)
	n, _ := c.Request.Body.Read(buf)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))

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

	valid := validation.Validation{}
	valid.Required(siteTitle,"site_title").Message("网站标题不能为空")
	valid.Required(SiteDesc,"site_desc").Message("网站描述不能为空")
	valid.Required(SiteKeyboard,"site_keyboard").Message("关键字不能为空")
	valid.Required(SiteCopyright,"site_copyright").Message("版权不能为空")
	//valid.Required(SiteTel,"site_tel").Message("电话不能为空")
	valid.Required(SiteEmail,"site_email").Message("邮箱不能为空")
	valid.Required(SiteAddress,"site_address").Message("地址不能为空")

	data := make(map[string]interface{})
	isOk := false

	if !valid.HasErrors() {
		if err :=validTel(SiteTel,LandLine,ClientTel);err{
			return e.ERROR,"电话联系方式，必须填写一项"
		}
		data["site_title"] = siteTitle
		data["site_desc"] = SiteDesc
		data["site_keyboard"] = SiteKeyboard
		data["site_copyright"] = SiteCopyright
		if SiteTel != "" {
			data["site_tel"] = SiteTel
		}
		if LandLine != "" {
			data["land_line"] = LandLine
		}
		if ClientTel != "" {
			data["client_tel"] = ClientTel
		}
		data["site_email"] = SiteEmail
		data["site_address"] = SiteAddress
		data["record_number"] = RecordNumber
		if id < 1 {
			isOk = Site.AddSite(data)
		} else {
			isOk = Site.EditSite(id,data)
		}
		if isOk {
			return e.SUCCESS,"操作成功"
		}
	}
	return ViewErr(valid)
}

// @Summer 获取站点信息
func GetSite() (sites Site.Site){
	return Site.GetSite()
}

func validTel(tel,land,client string) bool  {
	if tel == "" && land == "" && client == "" {
		return true
	}
	return false
}