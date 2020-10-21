package Article

import (
	"bufio"
	"elearn100/Backend/Controller"
	db "elearn100/Database"
	"elearn100/Model/Nav"
	"elearn100/Pkg/setting"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Article struct {
	ID   int     `gorm:"primary_key" json:"id"`
	Navs Nav.Nav `json:"nav" gorm:"FOREIGNKEY:NavId;ASSOCIATION_FOREIGNKEY:ID"`

	Title     string `json:"title" gorm:"type:varchar(190);not null;default '';comment:'标题'"`
	Summary   string `json:"summary" gorm:"type:varchar(255);not null;default '';comment:'摘要'"`
	ThumbImg  string `json:"thumb_img" gorm:"type:varchar(190);not null;default '';comment:'缩率图'"`
	Admin     string `json:"admin" gorm:"type:varchar(190);not null;default '';comment:'发布者'"`
	Com       string `json:"com" gorm:"type:varchar(190);not null;default '';comment:'来源'"`
	IsShow    int    `json:"is_show" gorm:"not null;default '1';comment:'是否展示 1展示 2不展示'"`
	Content   string `json:"content" gorm:"type:text;not null;default '';comment:'内容'"`
	Hot       int    `json:"hot" gorm:"not null;default '2';comment:'是否热点 1是 2否'"`
	Sort      int    `json:"sort" gorm:"not null;default '0';comment:'优先级 数字越大，排名越前'"`
	NavId     int    `json:"nav_id" gorm:"default '0';comment:'栏目ID'"`
	CreatedAt string `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

// @Summer 添加文章
func AddArticle(data map[string]interface{}) bool {
	article := db.Db.Create(&Article{
		Title:    data["title"].(string),
		Summary:  data["summary"].(string),
		ThumbImg: data["thumb_img"].(string),
		Admin:    data["admin"].(string),
		Com:      data["com"].(string),
		IsShow:   data["is_show"].(int),
		Content:  data["content"].(string),
		Hot:      data["hot"].(int),
		Sort:     data["sort"].(int),
		NavId:    data["nav_id"].(int),
	})

	if article.Error != nil {
		fmt.Print("添加文章失败", article)
		return false
	}
	return true
}

// @Summer 编辑文章
func EditArticle(id int, data interface{}) bool {
	edit := db.Db.Model(&Article{}).Where("id = ?", id).Update(data)
	if edit.Error != nil {
		fmt.Print("编辑文章失败", edit)
		return false
	}
	return true
}

// @Summer 获取所有文章
func GetArticles(page int, data interface{}) (articleS []Article) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * setting.PageSize
	}
	db.Db.Where(data).Offset(offset).Limit(setting.PageSize).Order("created_at desc").Find(&articleS)
	return
}

// @Summer 获取单篇文章
func GetArticle(id int) (article Article) {
	db.Db.Preload("Navs").Where("id = ?", id).First(&article)
	return
}

// @Summer 统计
func GetArticleTotal() (count int) {
	db.Db.Model(&Article{}).Count(&count)
	return
}

// @Summer 当前时间减去1年
func SubTime(timesTr int64) string {
	nowTime := time.Unix(timesTr, 0)
	return nowTime.AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
}

// @Summer下载图片
func DownImg(imgUrl string) string {
	fileName := path.Base(imgUrl)
	if strings.Contains("map", imgUrl) {
		return "mp4"
	}
	if strings.Contains("video_player", imgUrl) {
		return ""
	}
	h := http.Client{
		Timeout: time.Duration(86400 * time.Second),
	}
	res, err := h.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!:", err)
		return ""
	}

	defer res.Body.Close()
	reader := bufio.NewReader(res.Body)

	imgPath, fileErr := Controller.UploadDir()
	if !fileErr {
		fmt.Println("目录创建失败")
		return ""
	}

	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}

	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	writerLen, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("下载图片失败", err)
		return ""
	}
	if writerLen > 0 {
		return RenameImg(fileName)
	}
	return fileName
}

// @Summer 图片重命名
func RenameImg(imgName string) string {
	if imgName == "" {
		fmt.Println("图片重命名：空的图片")
		return ""
	}

	dir, fileErr := Controller.UploadDir()
	if !fileErr {
		fmt.Println("目录创建失败")
		return ""
	}
	fileName := ""
	if ThumbImg(imgName) != "" {
		ot := time.Now().UnixNano() / int64(time.Millisecond)
		_S := strconv.FormatInt(ot, 10)
		fileName = _S + ThumbImg(imgName)
	}

	_, err := os.Stat(dir + imgName)
	if err == nil {
		filePath := dir + fileName
		err = os.Rename(dir+imgName, filePath)
		if err != nil {
			fmt.Println("reName Error", err)
			return ""
		}
		return fileName
	}
	return ""
}

// @Summer 剔除url尾部字符串
func TrimUrl(prefix, url string) string {
	subUrl := strings.TrimRight(url, prefix) + ThumbImg(url)
	downImg := DownImg(subUrl)
	if downImg == "mp4" {
		return url
	}
	if strings.Contains("video_player", url) {
		return url
	}
	imgUrl := ""
	if reader, err := os.Open(downImg); err != nil {
		defer reader.Close()
		_, _, err := image.DecodeConfig(reader)
		if err != nil {
			_newImg := Controller.QiNiu(downImg)
			if _newImg != "" {
				imgUrl = "http://img.cdn.brocaedu.com/" + _newImg
			}
		} else {
			return url
		}
	}

	return imgUrl
}

// @Summer 替换图片地址
func ReplaceUrl(content, url, oldUrl string) string {
	res := strings.Replace(content, oldUrl, url, -1)
	return res
}

// @替换文章中的图片地址
func ReplaceContent(content string) string {
	reg := regexp.MustCompile(`data-src="(.*?)"`)
	res := reg.FindAllStringSubmatch(content, -1)
	for _, v := range res {
		png := ThumbImgType(v[1]) //原有图片地址
		_newUrl := TrimUrl(png, v[1])
		if _newUrl != "" || !strings.Contains("mp4", png) {
			content = ReplaceUrl(content, _newUrl, v[1])
		}
	}
	resCon := strings.Replace(content, "data-src", "src", -1)
	return resCon
}

// @Summer 返回图片结尾的参数
func ThumbImgType(ImgSrc string) string {
	png := ""
	if strings.HasSuffix(ImgSrc, "?wx_fmt=png") {
		png = "?wx_fmt=png"
	}
	if strings.HasSuffix(ImgSrc, "?wx_fmt=jpeg") {
		png = "?wx_fmt=jpeg"
	}
	if strings.HasSuffix(ImgSrc, "?wx_fmt=jpg") {
		png = "?wx_fmt=jpg"
	}
	if strings.HasSuffix(ImgSrc, "?wx_fmt=gif") {
		png = "?wx_fmt=gif"
	}
	return png
}

// @Summer 返回图片结尾的参数
func ThumbImg(ImgSrc string) string {
	png := ""
	if strings.HasSuffix(ImgSrc, "png") {
		png = ".png"
	}
	if strings.HasSuffix(ImgSrc, "jpeg") {
		png = ".jpeg"
	}
	if strings.HasSuffix(ImgSrc, "jpg") {
		png = ".jpg"
	}
	if strings.HasSuffix(ImgSrc, "gif") {
		png = ".gif"
	}
	if strings.HasSuffix(ImgSrc, "mp4") {
		png = ""
	}
	return png
}
