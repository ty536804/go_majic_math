package Services

import (
	"elearn100/Model/Message"
	"elearn100/Model/Site"
	"elearn100/MqQueue/Mq"
	"elearn100/Pkg/e"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"regexp"
	"strings"
	"time"
)

type Info struct {
	MName   string
	Area    string
	Tel     string
	Client  string
	Ip      string
	Uid     string
	MsgType int
}

// @Title提交留言
func AddMessage(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}
	tel := com.StrTo(c.PostForm("tel")).String()
	re := regexp.MustCompile(`^1\d{10}$`)

	fmt.Println(re.MatchString(tel), tel, len(tel))
	if !re.MatchString(tel) || len(tel) < 11 {
		return e.SUCCESS, "请填写有效的手机号码"
	}

	ip := c.Request.RemoteAddr
	initTime := time.Now().Format("2006-01-02")
	total := Message.GetTotalMessage(ip, initTime+" 00:00:00", initTime+" 23:59:59")
	if total >= 5 {
		return e.SUCCESS, "提交成功"
	}

	mname := TrimHtml(com.StrTo(c.PostForm("mname")).String())
	area := TrimHtml(com.StrTo(c.PostForm("area")).String())
	webType := com.StrTo(c.PostForm("client")).String()

	valid := validation.Validation{}
	valid.Required(mname, "mname").Message("姓名不能为空")
	valid.Required(area, "area").Message("地区不能为空")
	valid.Required(tel, "tel").Message("选择是否展示")
	if !valid.HasErrors() {
		uid := strings.Split(strings.Replace(ip, ".", "", -1), ":")[0]
		var word Info
		word.MName = mname
		word.Area = area
		word.Tel = tel
		word.Client = webType
		word.Ip = ip
		word.Uid = uid
		word.MsgType = 1
		jsonData, _ := json.Marshal(word)
		Mq.PublishEx("mofashuxue", "fanout", "", string(jsonData))
		sendSms(tel, area, mname) //发送短信
		return e.SUCCESS, "提交成功"
	}
	return ViewErr(valid)
}

type CodeInfo struct {
	Msg string
	Tel string
}

// @Title 发送短信
// @Param tel  string	电话
// @Param area string	区域
// @Param name string	名字
func sendSms(tel, area, name string) {
	site := Site.GetSite()
	var telList = strings.Split(strings.TrimSpace(site.AdminTel), ",")
	telList = append(telList, tel)

	for k, v := range telList { //发送短信
		var code CodeInfo
		msg := ""
		if (k + 1) == len(telList) {
			msg = "我们已收到您的留言。我们的招商经理会在24小时内联系您，请您注意接听来自北京的电话，谢谢。"
		} else {
			msg = area + "的" + name + "留言了。联系" + tel + "留言来源魔法数学"
		}
		code.Tel = v
		code.Msg = msg
		dataJson, _ := json.Marshal(code)
		Mq.Publish("", "smsinfo", string(dataJson))
	}
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
