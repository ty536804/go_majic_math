package Services

import (
	"elearn100/Model/Article"
	"elearn100/Pkg/e"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加文章
func AddArticle(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return e.ReError()
	}
	id := com.StrTo(c.PostForm("id")).MustInt()
	title := com.StrTo(c.PostForm("title")).String()
	summary := com.StrTo(c.PostForm("summary")).String()
	admin := com.StrTo(c.PostForm("admin")).String()
	content := com.StrTo(c.PostForm("content")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	sort := com.StrTo(c.PostForm("sort")).MustInt()
	hot := com.StrTo(c.PostForm("hot")).MustInt()
	thumbImg := com.StrTo(c.PostForm("thumb_img")).String()
	articleCom := com.StrTo(c.PostForm("com")).String()
	navId := com.StrTo(c.PostForm("nav_id")).MustInt()

	if code, err := validArt(title, summary, content, admin, hot, isShow, navId); code == e.ERROR {
		return code, err
	}

	article := Article.Article{
		Title:    title,
		Summary:  summary,
		IsShow:   isShow,
		Hot:      hot,
		Sort:     sort,
		ThumbImg: thumbImg,
		Admin:    admin,
		Com:      articleCom,
		NavId:    navId,
		Content:  content,
	}

	isOk := false

	if id < 1 {
		isOk = Article.AddArticle(article)
	} else {
		isOk = Article.EditArticle(id, article)
	}
	if isOk {
		return e.ReSuccess()
	}
	return e.ReError()
}

// @Desc 验证
func validArt(title, summary, content, admin string, hot, isShow, navId int) (code int, msg string) {
	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(summary, "summary").Message("摘要不能为空")
	valid.Required(isShow, "is_show").Message("选择是否展示")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(hot, "hot").Message("选择是否热点")
	valid.Required(admin, "admin").Message("发布者不能为空")
	valid.Required(navId, "nav_id").Message("栏目不能为空")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Desc 获取缓存中的文章
func RedisGetArticles(page, pageNum int, redisKey string) []Article.Article {
	conn := e.PoolConnect()
	defer conn.Close()

	redisKey = e.REDISKey + redisKey
	var articles []Article.Article

	if exists, _ := redis.Bool(conn.Do("exists", redisKey)); exists {
		values, _ := redis.Values(conn.Do("lrange", redisKey, 0, -1))
		var article Article.Article
		for _, v := range values {
			if err := json.Unmarshal(v.([]byte), &article); err == nil {
				articles = append(articles, article)
			}
		}
	} else {
		articles = Article.GetArticles(page, pageNum, make(map[string]interface{}))
		for _, v := range articles {
			if jsonStr, err := json.Marshal(v); err == nil {
				conn.Do("rpush", redisKey, jsonStr)
			}
		}
		conn.Do("expire", redisKey, e.VALIDTime)
	}
	return articles
}
