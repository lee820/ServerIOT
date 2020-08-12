package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/internal/middleware"
	"github.com/lee820/ServerIOT/internal/routers/api"
	v1 "github.com/lee820/ServerIOT/internal/routers/api/v1"
)

//NewRouter 新增路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	r.GET("/auth", api.GetAuth)
	login := v1.NewLogin()
	dev := v1.NewDevRouter()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/login", login.UserRegister)
		apiv1.GET("/login", login.UserLogin)
		//apiv1.DELETE("/login/:id", login.UserLogout)
		apiv1.POST("/device", dev.CreateDevice)
		apiv1.POST("/device/:id", dev.UpdateDevice)
		apiv1.DELETE("/device/:id", dev.DeleteDevice)
		apiv1.GET("/device/:id", dev.GetDeviceList)
	}

	return r
}
