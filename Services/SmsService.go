package Services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

var (
	_account  = "短信" //用户名是登录用户中心->国际短信->产品总览->APIID
	_password = "短信" //查看密码请登录用户中心->国际短信->产品总览->APIKEY
	_now      = strconv.FormatInt(time.Now().Unix(), 10)
	_url      = "https://106.ihuyi.com/webservice/sms.php?method=Submit&format=json"
)

// @Summer 发送验证码
func SendSms(mobile, msg string) {
	v := url.Values{}
	//fmt.Printf(_now)

	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+mobile+msg+_now))
	v.Set("mobile", mobile)
	v.Set("content", msg)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", _url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("看下发送的结构 %+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}

func SendSmsToClient(area, name, tel string) {
	msg := "我们已收到您的留言。我们的招商经理会在24小时内联系您，请您注意接听来自北京的电话，谢谢。"
	SendSms(tel, msg)
	msg = area + "的" + name + "留言了。联系" + tel
	SendSms("13811384847", msg)
	SendSms("13811221394", msg)
}
