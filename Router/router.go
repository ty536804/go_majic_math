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
	"elearn100/Middleware/jwt"
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdin)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	dir := e.GetDir()
	//加载后端js、样式文件
	r.StaticFS("static", http.Dir(dir+"/Resources/Public"))
	//加载后端文件
	r.LoadHTMLGlob("Resources/View/**/*")
	//首页
	r.GET("/", frontend.Index)
	r.GET("/about", frontend.About)
	r.GET("/aboutData", frontend.AboutData)
	r.GET("/index", frontend.FrontEnd)
	r.GET("/subject", frontend.Subject)
	r.GET("/research", frontend.Research)
	r.GET("/learn", frontend.Learn)
	r.GET("/news", frontend.News)
	r.GET("/newList", frontend.NewList)
	r.GET("/detail", frontend.NewDetail)
	r.GET("/newDetail", frontend.NewDetailData)
	r.GET("/join", frontend.Authorize)
	r.GET("/joinData", frontend.JoinData)
	r.GET("/omo", frontend.Omo)
	r.GET("/campus", frontend.Campus) //全国校区
	r.GET("/down", frontend.Down)
	//移动端
	r.GET("/wap", frontend.WapIndex)
	//Backend
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//上传图片
		apiv1.POST("/upload", backend.UploadFile)
		//admin
		r.GET("/admin", func(c *gin.Context) {
			uuid, uOk := c.Request.Cookie("uuid")
			isOk, _ := e.GetVal("token")
			_html := "admin/home.html"
			if uOk != nil || !isOk || len(uuid.Value) == 0 {
				Services.LogOut(c)
				_html = "admin/login.html"
			}
			c.HTML(e.SUCCESS, _html, gin.H{
				"title": "易学教育",
			})
		})
		apiv1.GET("/powerShow", v1.PowerShow)
		apiv1.GET("/show", v1.Show)
		r.POST("/login", v1.Login)
		//用户列表
		apiv1.GET("/index", v1.BackEndIndex)
		apiv1.GET("/userList", v1.UserList)
		apiv1.GET("/userData", v1.UserData)       //用户列表API
		apiv1.POST("/AddUser", v1.AddUser)        //用户列表API
		apiv1.POST("/GetUser", v1.GetUser)        //查看当用户API
		apiv1.GET("/logout", v1.LogOut)           //查看当用户API
		apiv1.POST("/editUser", v1.UpdateUser)    //获取站点信息API
		apiv1.GET("/detailsUser", v1.DetailsUser) //查看当用户API
		//站点信息
		apiv1.GET("/siteInfo", v1.SiteInfo) //查看站点信息API
		apiv1.POST("/addSite", v1.AddSite)  //添加站点信息API
		apiv1.GET("/getSite", v1.GetSite)   //获取站点信息API
		//banner
		apiv1.GET("/bannerList", v2.List)
		apiv1.GET("/getBanners", v2.GetBanners)
		apiv1.GET("/bannerDetail", v2.Detail)
		apiv1.POST("/AddBanner", v2.AddBanner)
		apiv1.GET("/getBanner", v2.GetBanner)
		apiv1.POST("/delBanner", v2.DelBanner)
		//message
		apiv1.GET("/messageList", m.List)
		apiv1.GET("/messageData", m.ListData)
		r.POST("/AddMessage", m.AddMessage)
		//导航
		apiv1.GET("/getNavs", nav.GetNavs)    //获取多条导航API
		apiv1.POST("/getNav", nav.GetNav)     //获取一条导航API
		apiv1.GET("/navList", nav.Show)       //导航列表展示
		apiv1.POST("/addNav", nav.AddNav)     //添加导航API
		r.POST("/getNavList", nav.GetNavList) //添加导航API
		//文章
		apiv1.GET("/article", v3.Show)           //文章列表
		apiv1.POST("/articleList", v3.ShowList)  //文章列表API
		apiv1.GET("/articleDetail", v3.Detail)   //文章列表
		apiv1.POST("/getArticle", v3.GetArticle) //文章详情API
		apiv1.POST("/addArticle", v3.AddArticle) //文章详情API
		//单页
		apiv1.GET("/single", v4.List)               //文章列表
		apiv1.POST("/singleList", v4.ListData)      //文章列表API
		apiv1.GET("/list", v4.List)                 //文章列表
		apiv1.POST("/getSingle", v4.GetSingle)      //文章详情API
		apiv1.POST("/addSingle", v4.AddSingle)      //添加单页详情API
		apiv1.GET("/singleDetail", v4.DetailSingle) //文章详情API
		//校区
		apiv1.GET("/campus", campus.Index)               //校区首页
		apiv1.POST("/campusList", campus.GetCampus)      //校区列表API
		apiv1.POST("/campusDetail", campus.DetailCampus) //校区详情
		r.POST("/groupCampuses", campus.GroupCampuses)   //校区列表API 带缓冲区的
		r.POST("/campusData", campus.GetCampuses)        //校区列表API 带缓冲区的
		apiv1.POST("/addCampuses", campus.AddCampuses)   //添加校区
	}

	return r
}
