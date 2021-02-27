package Services

import (
	"bytes"
	db "elearn100/Database"
	"elearn100/Model/Article"
	"elearn100/Pkg/e"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var wt sync.WaitGroup

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

	if err = json.NewDecoder(resp.Body).Decode(&jMap); err != nil {
		return "", errors.New("request token response json parse err :" + err.Error())
	}

	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
		accessToken, _ := jMap["access_token"].(string)

		conn := e.PoolConnect()
		defer conn.Close()

		if _, err := conn.Do("set", "access_token", accessToken); err == nil {
			conn.Do("expire", "access_token", 2*time.Hour) //设置缓存
		}
		return accessToken, nil
	}
	errcode := jMap["errcode"].(string)
	errmsg := jMap["errmsg"].(string)
	err = errors.New(errcode + ":" + errmsg)
	return "", err
}

type BatChGetMaterial struct {
	Item []struct {
		MediaId string `json:"media_id"`
		Content struct {
			NewsItem []struct {
				Title              string `json:"title"`
				Author             string `json:"author"`
				Digest             string `json:"digest"`
				Content            string `json:"content"`
				ContentSourceUrl   string `json:"content_source_url"`
				ThumbMediaId       string `json:"thumb_media_id"`
				ShowCoverPic       int    `json:"show_cover_pic"`
				Url                string `json:"url"`
				ThumbUrl           string `json:"thumb_url"`
				NeedOpenComment    int    `json:"need_open_comment"`
				OnlyFansCanComment int    `json:"only_fans_can_comment"`
			} `json:"news_item"`
			CreateTime int64 `json:"create_time"`
			UpdateTime int64 `json:"update_time"`
		}
		UpdateTime int `json:"update_time"`
	}
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
}

// @Summer 微信获取文章
func GetArticle(begin, count int) {
	result, err := ResolveUrl(begin, count)
	var article = make(map[string]interface{})

	if err != nil {
		fmt.Printf("read resp.body failed,err:%v\n", err)
	} else {
		stu := &BatChGetMaterial{}
		res := json.Unmarshal(result, &stu)
		if res == nil {
			for _, item := range stu.Item {
				res := item.Content.NewsItem[0]
				if res.Title != "" {
					tit := strings.TrimSpace(res.Title)
					currentTime := time.Unix(item.Content.CreateTime, 0).Format("2006-01-02 15:04:05")
					if strings.Contains("练脑时刻", tit) {
						currentTime = Article.SubTime(item.Content.UpdateTime)
					}
					// url尾部字符串
					imgType := Article.ThumbImgType(res.ThumbUrl)
					thumbImg := Article.TrimUrl(imgType, res.ThumbUrl)
					article["title"] = res.Title
					article["summary"] = res.Digest
					article["thumb_img"] = res.ThumbUrl
					article["admin"] = res.Author
					article["com"] = "weChat"
					article["is_show"] = 1
					article["content"] = Article.ReplaceContent(res.Content)
					article["hot"] = 0
					article["sort"] = 0
					article["nav_id"] = 8
					article["created_at"] = currentTime
					if thumbImg != "" {
						wt.Add(1)
						go WeAddArticle(article)
					}
				}
			}
			wt.Wait()
		}
	}
}

func ResolveUrl(offset, count int) ([]byte, error) {
	conn := e.PoolConnect()
	defer conn.Close()

	accessToken, isOk := redis.String(conn.Do("get", "access_token"))
	if isOk == nil {
		token, err := GetToken()
		if err != nil {
			panic(err)
		}
		accessToken = token
	}
	data := make(map[string]interface{})
	data["type"] = "news"
	data["offset"] = offset
	data["count"] = count
	url := "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=" + accessToken

	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bytes.NewReader(bytesData)

	rep, err := http.NewRequest("POST", url, reader)
	resp, err := http.DefaultClient.Do(rep)

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// @Summer 添加文章
func WeAddArticle(data map[string]interface{}) bool {
	defer wt.Done()

	UpdatedAt := time.Now().Format("2006-01-02 15:04:05")
	article := db.Db.Create(&Article.Article{
		Title:     data["title"].(string),
		Summary:   data["summary"].(string),
		ThumbImg:  data["thumb_img"].(string),
		Admin:     data["admin"].(string),
		Com:       data["com"].(string),
		IsShow:    data["is_show"].(int),
		Content:   data["content"].(string),
		Hot:       data["hot"].(int),
		Sort:      data["sort"].(int),
		NavId:     data["nav_id"].(int),
		CreatedAt: data["created_at"].(string),
		UpdatedAt: UpdatedAt,
	})

	if article.Error != nil {
		fmt.Print("添加文章失败", article)
		return false
	}
	return true
}

func GetArt() {
	total := Article.GetArticleTotal()
	GetArticle(total+1, 1)
}
