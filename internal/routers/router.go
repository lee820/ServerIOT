package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/internal/middleware"
	v1 "github.com/lee820/ServerIOT/internal/routers/api/v1"
)

//NewRouter 新增路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	user := v1.NewUser()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/user", user.Create)
		apiv1.DELETE("/user/:id", user.Delete)
		apiv1.PUT("/user/:id", user.Update)
		apiv1.GET("/user/:id", user.Retrieve)
	}

	return r
}
