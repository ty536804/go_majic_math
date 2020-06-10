package Services

import (
	"bytes"
	"elearn100/Pkg/e"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	APPID     = "wxc6fc8246185aa2b8"
	APPSECRET = "fd85ee04d782f48418bb2baaa474106a"
	GRANTTYPE = "client_credential"
)

// @Summer 获取token
func GetToken() (string, error) {
	u, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token")

	if err != nil {
		log.Fatal(err)
	}

	parse := url.Values{}
	parse.Set("grant_type", GRANTTYPE)
	parse.Set("appid", APPID)
	parse.Set("secret", APPSECRET)
	u.RawQuery = parse.Encode()

	resp, err := http.Get(u.String())

	jMap := make(map[string]interface{})

	if err != nil {
		return "", errors.New("request token err :" + err.Error())
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return "", errors.New("request token response json parse err :" + err.Error())
	}

	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
		accessToken, _ := jMap["access_token"].(string)
		e.SetAccessToken(accessToken) //设置缓存
		return accessToken, nil
	} else {
		errcode := jMap["errcode"].(string)
		errmsg := jMap["errmsg"].(string)
		err = errors.New(errcode + ":" + errmsg)
		return "", err
	}
}

func GetArticle() {
	isOk, accessToken := e.GetVal("access_token")
	if !isOk {
		token, err := GetToken()
		if err != nil {
			panic(err)
		}
		accessToken = token
	}

	url := "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=" + accessToken
	data := make(map[string]interface{})
	data["begin_date"] = "2020-04-01"
	data["end_date"] = "2020-04-01"

	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bytes.NewReader(bytesData)

	rep, err := http.NewRequest("POST", url, reader)
	resp, err := http.DefaultClient.Do(rep)

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.body failed,err:%v\n", err)
	} else {
		fmt.Println(string(b))
	}
}
