package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/service"
	"github.com/lee820/ServerIOT/pkg/app"
	"github.com/lee820/ServerIOT/pkg/errcode"
)

//Login 登录接口
type Login struct {
}

//NewLogin 返回登录结构体
func NewLogin() Login {
	return Login{}
}

//UserLogin 用户登录
func (l Login) UserLogin(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("qpp.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//UserRegister 用户注册接口
func (l Login) UserRegister(c *gin.Context) {

}

//UserLogout 用户退出登录
func (l Login) UserLogout(c *gin.Context) {

}
