package v1


import (
	v1 "elearn100/backend/Controller/Admin"
	v2 "elearn100/backend/Controller/Banner"
	m "elearn100/backend/Controller/Message"
	"elearn100/middleware/jwt"
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())

	dir := e.GetDir()
	r.StaticFS("static",http.Dir(dir+"/backend/Public"))

	r.LoadHTMLGlob("backend/View/**/*")
	//login
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(e.SUCCESS,"admin/login.html",gin.H{
			"title": "易学教育",
		})
	})

	r.Group("/admin")
	{
		r.POST("/login", v1.Login)
	}

	apiv1 := r.GET("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		r.GET("/userData",v1.UserData)//用户列表API
		r.POST("/AddUser",v1.AddUser)//用户列表API
		r.POST("/GetUser",v1.GetUser)//查看当用户信息API
		r.GET("/show", v1.Show)
		r.GET("/welcome", v1.Welcome)
		r.GET("/powerShow",v1.PowerShow)
		r.GET("/userList", v1.UserList)//用户列表
		r.GET("/bannerList", v2.List)
		r.GET("/bannerDetail", v2.Detail)
		r.GET("/message",m.List)
		r.GET("/messageData",m.ListData)
	}
	return r
}
