package v1

import (
	"net/http"

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
// @Summary 用户登录
// @Produce json
// @Param name query string true "用户名"
// @Param password query string true "用户密码"
// @Success 200
// @Failure 400 (object) errcode.Error "请求错误"
// @Router /api/v1/tags [get] func (t Tag)List(c *gin.Conttext){}
func (l Login) UserLogin(c *gin.Context) {
	param := service.GetUserInfoByNameRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("qpp.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	//查询用户是否存在
	getUser, err := svc.GetUserInfoByUsername(&param)
	if err != nil {
		//global.Logger.Errorf("svc.QueryUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotFound)
		return
	}
	//查询用户密码是否正确
	if getUser.Password != param.Password {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorWrongPassword)
		return
	}
	response.ToResponse(gin.H{
		"code": http.StatusOK,
		"Msg":  "登录成功",
		"Data": "token:",
	})

	return
}

//UserRegister 用户注册接口
func (l Login) UserRegister(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("UserRegister app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		c.Abort()
	}
	//服务层实例
	svc := service.New(c.Request.Context())
	//查询用户是否存在
	userInfoByPhone := service.GetUserInfoRequestByPhone{
		Phone: param.Phone,
	}
	getUser, err := svc.GetUserInfoByPhone(&userInfoByPhone)
	if err != nil {
		//查询失败
		global.Logger.Errorf("UserRegister svc.GetUserInfoByPhone err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserInfoFail)
		c.Abort()
	}
	if getUser.Phone != "" {
		//用户手机号已被注册
		response.ToErrorResponse(errcode.ErrorPhoneExist)
		c.Abort()
	}

	//进行注册，创建用户
	err = svc.CreateUser(&param)
	if err != nil {
		//创建失败
		global.Logger.Errorf("UserRegister svc.GetUserInfoByPhone err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		c.Abort()
	}

	//注册成功，创建用户成功。发放token
	response.ToResponse(gin.H{
		"code": http.StatusOK,
		"Msg":  "注册成功",
		"Data": "token:",
	})
}

//UserLogout 用户退出登录
func (l Login) UserLogout(c *gin.Context) {

}
