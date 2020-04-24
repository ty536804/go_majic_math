package Router

import (
	backend "elearn100/backend/Controller"
	v1 "elearn100/backend/Controller/Admin"
	v3 "elearn100/backend/Controller/Article"
	v2 "elearn100/backend/Controller/Banner"
	m "elearn100/backend/Controller/Message"
	nav "elearn100/backend/Controller/Nav"
	frontend "elearn100/frontend/Controller"
	"elearn100/middleware/jwt"
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())

	dir := e.GetDir()
	//加载后端js、样式文件
	r.StaticFS("static",http.Dir(dir+"/resources/Public"),)
	//加载后端文件
	r.LoadHTMLGlob("resources/View/**/*")
	//首页
	r.GET("/",frontend.Index)

	//backend
	//login
	r.GET("/admin", func(c *gin.Context) {
		isOk, _:= e.GetVal("token")
		_html := "admin/login.html"
		if isOk {
			_html = "admin/home.html"
		}
		c.HTML(e.SUCCESS,_html,gin.H{
			"title": "易学教育",
		})
	})
	r.POST("/login", v1.Login)
	r.GET("/show", v1.Show)

	apiv1 :=r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//上传图片
		apiv1.POST("/upload",backend.UploadFile)
		//admin
		apiv1.GET("/welcome", v1.Welcome)
		apiv1.GET("/powerShow",v1.PowerShow)
		//用户列表
		apiv1.GET("/userList", v1.UserList)
		apiv1.GET("/userData",v1.UserData)//用户列表API
		apiv1.POST("/AddUser",v1.AddUser)//用户列表API
		apiv1.POST("/GetUser",v1.GetUser)//查看当用户API
		apiv1.GET("/logout",v1.LogOut)//查看当用户API
		//站点信息
		apiv1.GET("/siteInfo",v1.SiteInfo)//查看站点信息API
		apiv1.POST("/addSite",v1.AddSite)//添加站点信息API
		apiv1.GET("/getSite",v1.GetSite)//获取站点信息API
		//banner
		apiv1.GET("/bannerList", v2.List)
		apiv1.GET("/getBanners", v2.GetBanners)
		apiv1.GET("/bannerDetail", v2.Detail)
		apiv1.POST("/AddBanner", v2.AddBanner)
		apiv1.GET("/getBanner", v2.GetBanner)
		//message
		apiv1.GET("/message",m.List)
		apiv1.GET("/messageData",m.ListData)
		//导航
		apiv1.GET("/getNavs",nav.GetNavs)//获取多条导航API
		apiv1.POST("/getNav",nav.GetNav)//获取一条导航API
		apiv1.GET("/navList",nav.Show)//导航列表展示
		apiv1.POST("/addNav",nav.AddNav)//添加导航API
		r.GET("/getNavList",nav.GetNavList)//添加导航API
		//文章
		apiv1.GET("/article",v3.Show)//文章列表
		apiv1.POST("/articleList",v3.ShowList)//文章列表API
		apiv1.GET("/articleDetail",v3.Detail)//文章列表
		apiv1.POST("/getArticle",v3.GetArticle)//文章详情API
		apiv1.POST("/addArticle",v3.AddArticle)//文章详情API
	}
	return r
}
