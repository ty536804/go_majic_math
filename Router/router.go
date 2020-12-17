package Router

import (
	backend "elearn100/Backend/Controller"
	v1 "elearn100/Backend/Controller/Admin"
	v3 "elearn100/Backend/Controller/Article"
	v2 "elearn100/Backend/Controller/Banner"
	campus "elearn100/Backend/Controller/Campus"
	m "elearn100/Backend/Controller/Message"
	nav "elearn100/Backend/Controller/Nav"
	v4 "elearn100/Backend/Controller/Single"
	frontend "elearn100/Frontend/Controller"
	"elearn100/Frontend/Controller/Wap"
	"elearn100/Middleware/jwt"
	"elearn100/Pkg/e"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"os"
)

func unescaped(x string) interface{} { return template.HTML(x) }

func subYear(x string) interface{} {
	return x[0:4]
}

func subDate(x string) interface{} {
	return x[5:10]
}

func totalItem(x []interface{}) int {
	return len(x)
}

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdin)

	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"unescaped": unescaped,
		"subYear":   subYear,
		"subDate":   subDate,
		"totalItem": totalItem,
	})

	r.Use(gin.Logger(), gin.Recovery())

	dir := e.GetDir()
	//加载后端js、样式文件
	r.StaticFS("static", http.Dir(dir+"/Resources/Public"))
	r.LoadHTMLFiles()
	//加载后端文件
	r.LoadHTMLGlob("Resources/View/**/*")
	//首页
	r.GET("/", frontend.Index)
	r.GET("/about", frontend.About)
	r.GET("/subject", frontend.Subject)
	r.GET("/research", frontend.Research)
	r.GET("/learn", frontend.Learn)
	r.GET("/news", frontend.News)
	r.GET("/newList", frontend.NewList)
	r.GET("/detail", frontend.NewDetail)
	r.GET("/join", frontend.Authorize)
	r.GET("/omo", frontend.Omo)
	r.GET("/campus", frontend.Campus) //全国校区
	r.GET("/down", frontend.Down)
	r.GET("/weChat1", frontend.GetWeChat)
	r.GET("/groupCampuses", campus.GroupCampuses) //校区列表API 带缓冲区的
	r.POST("/campusData", campus.GetCampuses)     //校区列表API 带缓冲区的
	//移动端
	r.GET("/wap", Wap.Index)
	r.GET("/sub", Wap.Subject)
	r.GET("/le", Wap.Learn)
	r.GET("/om", Wap.Omo)
	r.GET("/authorize", Wap.Authorize)
	r.POST("/AddMessage", m.AddMessage)
	//Backend
	r.GET("/login", func(c *gin.Context) {
		c.HTML(e.SUCCESS, "admin/login.html", gin.H{
			"title": "登录",
		})
	})
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//上传图片
		apiv1.POST("/upload", backend.UploadFile)
		apiv1.POST("/menus", v1.GetPowers)
		apiv1.POST("/newList", frontend.NewList)
		apiv1.POST("/login", v1.Login)
		//用户列表
		apiv1.POST("/userData", v1.UserData)       //用户列表API
		apiv1.POST("/AddUser", v1.AddUser)         //用户列表API
		apiv1.POST("/GetUser", v1.GetUser)         //查看当用户API
		apiv1.GET("/logout", v1.LogOut)            //查看当用户API
		apiv1.POST("/editUser", v1.UpdateUser)     //获取站点信息API
		apiv1.POST("/detailsUser", v1.DetailsUser) //查看当用户API
		//站点信息
		apiv1.POST("/addSite", v1.AddSite) //添加站点信息API
		apiv1.POST("/getSite", v1.GetSite) //获取站点信息API
		//banner
		apiv1.POST("/getBanners", v2.GetBanners)
		apiv1.POST("/AddBanner", v2.AddBanner)
		apiv1.POST("/getBanner", v2.GetBanner)
		apiv1.POST("/delBanner", v2.DelBanner)
		//message
		apiv1.POST("/messageData", m.ListData)
		//导航
		apiv1.POST("/getNavs", nav.GetNavs)     //获取多条导航API
		apiv1.POST("/getNav", nav.GetNav)       //获取一条导航API
		apiv1.POST("/addNav", nav.AddNav)       //添加导航API
		apiv1.POST("/menuList", nav.GetNavList) //添加导航API
		//文章
		apiv1.POST("/articleList", v3.ShowList)  //文章列表API
		apiv1.POST("/getArticle", v3.GetArticle) //文章详情API
		apiv1.POST("/addArticle", v3.AddArticle) //编辑文章
		//单页
		apiv1.POST("/singleList", v4.ListData) //单页列表详情API
		apiv1.POST("/getSingle", v4.GetSingle) //单页详情Api
		apiv1.POST("/addSingle", v4.AddSingle) //编辑单页
		//校区
		apiv1.POST("/campusList", campus.GetCampus)      //校区列表API
		apiv1.POST("/campusDetail", campus.DetailCampus) //校区详情
		apiv1.POST("/addCampuses", campus.AddCampuses)   //添加校区
	}

	return r
}
