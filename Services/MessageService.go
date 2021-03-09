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
	"strings"
)

// @Title提交留言
func AddMessage(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return e.ERROR, "网络请求失败，请稍后再试"
	}

	tel := com.StrTo(c.PostForm("tel")).String()
	if isOk := e.CheckPhone(tel); !isOk {
		return e.SUCCESS, "请填写有效的手机号码"
	}

	ip := e.GetIpAddress(c.Request.RemoteAddr)
	if total := Message.GetTotalMessage(ip); total >= 5 {
		return e.SUCCESS, "提交成功"
	}

	MName := e.TrimHtml(com.StrTo(c.PostForm("mname")).String())
	area := e.TrimHtml(com.StrTo(c.PostForm("area")).String())
	webType := com.StrTo(c.PostForm("client")).String()
	msgType := com.StrTo(c.PostForm("msg_type")).MustInt()
	webCom := com.StrTo(c.PostForm("com")).String()
	if code, err := validMessage(MName, area, tel); code == e.ERROR {
		return code, err
	}

	SendMessageForMq(MName, area, tel, webType, ip, webCom, msgType)
	sendSms(tel, area, MName) //发送短信
	return e.SUCCESS, "提交成功"
}

// @Desc 数据校验
func validMessage(MName, area, tel string) (int, string) {
	valid := validation.Validation{}
	valid.Required(MName, "mname").Message("姓名不能为空")
	valid.Required(area, "area").Message("地区不能为空")
	valid.Required(tel, "tel").Message("选择是否展示")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

type Info struct {
	MName   string
	Area    string
	Tel     string
	Client  string
	Ip      string
	Uid     string
	Com     string
	MsgType int
}

// @Desc 表单提交到队列
func SendMessageForMq(MName, area, tel, webType, ip, webCom string, msgType int) {
	word := new(Info)
	word.MName = MName
	word.Area = area
	word.Tel = tel
	word.Client = webType
	word.Ip = ip
	word.Com = webCom
	word.Uid = strings.Split(strings.Replace(ip, ".", "", -1), ":")[0]
	word.MsgType = msgType
	if jsonData, err := json.Marshal(word); err == nil {
		fmt.Println("添加留言")
		Mq.PublishEx("mofashuxue", "fanout", "", string(jsonData))
	} else {
		fmt.Println("json序列化失败")
	}
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
		msg := ""
		if (k + 1) == len(telList) {
			msg = "我们已收到您的留言。我们的招商经理会在24小时内联系您，请您注意接听来自北京的电话，谢谢。"
		} else {
			msg = area + "的" + name + "留言了。联系" + tel + "留言来源魔法数学"
		}
		code := new(CodeInfo)
		code.Tel = v
		code.Msg = msg
		dataJson, err := json.Marshal(code)
		if err == nil {
			Mq.Publish("", "smsinfo", string(dataJson))
		} else {
			fmt.Println("短信队列出错")
		}
	}
}
