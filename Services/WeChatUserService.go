package Services

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddWeChatUser(c *gin.Context) {
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println("获取用户信息失败", err)
	}

	mobile := com.StrTo(c.PostForm("mobile")).String()
	openId := com.StrTo(c.PostForm("open_id")).String()
	nickName := com.StrTo(c.PostForm("nick_name")).String()
	avatarUrl := com.StrTo(c.PostForm("avatar_url")).String()
	gender := com.StrTo(c.PostForm("gender")).MustInt()
	province := com.StrTo(c.PostForm("province")).String()
	city := com.StrTo(c.PostForm("city")).String()
	country := com.StrTo(c.PostForm("country")).String()
	subscribe := com.StrTo(c.PostForm("subscribe")).String()
	isBack := com.StrTo(c.PostForm("is_back")).MustInt()
	lookCode := com.StrTo(c.PostForm("look_code")).String()

	valid := validation.Validation{}
	valid.Required(openId, "open_id").Message("请关注微信公众号")

	data := make(map[string]interface{})
	if !valid.HasErrors() {
		data["mobile"] = mobile
		data["open_id"] = openId
		data["nick_name"] = nickName
		data["avatar_url"] = avatarUrl
		data["gender"] = gender
		data["province"] = province
		data["city"] = city
		data["country"] = country
		data["subscribe"] = subscribe
		data["is_back"] = isBack
		data["look_code"] = lookCode
	}
}
