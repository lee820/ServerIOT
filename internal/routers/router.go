package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/middleware"
	"github.com/lee820/ServerIOT/internal/pages"
	v1 "github.com/lee820/ServerIOT/internal/routers/api/v1"
	"github.com/lee820/ServerIOT/pkg/limiter"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

//NewRouter 新增api路由
func NewRouter(r *gin.Engine) {
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	login := v1.NewLogin()
	dev := v1.NewDevRouter()
	//r.GET("/auth", api.GetAuth)
	//用户注册和登录不需要token验证
	r.POST("/register", login.UserRegister)
	r.POST("/auth", login.UserLogin)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//apiv1.DELETE("/login/:id", login.UserLogout)
		apiv1.POST("/device", dev.CreateDevice)
		apiv1.POST("/device/:id", dev.UpdateDevice)
		apiv1.DELETE("/device/:id", dev.DeleteDevice)
		apiv1.GET("/device/:id", dev.GetDeviceList)
	}
}

//NewPageRouter 新增页面路由
func NewPageRouter(r *gin.Engine) {
	//配置静态文件目录
	r.Static("/css", "./dist/css")
	r.Static("/fonts", "./dist/fonts")
	r.Static("/js", "./dist/js")
	r.LoadHTMLFiles("./dist/index.html")

	r.GET("/login", pages.LoginPage)
}
