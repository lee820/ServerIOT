package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/service"
	"github.com/lee820/ServerIOT/pkg/app"
	"github.com/lee820/ServerIOT/pkg/errcode"
)

type DevRouter struct{}

//NewDevRouter 创建设备路由
func NewDevRouter() DevRouter {
	return DevRouter{}
}

func (d *DevRouter) CreateDevice(c *gin.Context) {
	param := service.CreateDeviceRequest{}
	response := app.NewResponse(c)
	//参数校验
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "createDevice bindAndValid ers: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	//查询用户是否存在
	userInfo, err := svc.GetUserInfoByID(param.UserID)
	if err != nil {
		global.Logger.Errorf(c, "CreateDevice svc.QueryUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotFound)
		return
	}
	//查询用户设备数量是否到达上限
	if userInfo.Devcount >= global.AppSetting.NormalUserDeviceUplimit {
		//达到上限，不能添加设备，返回
		response.ToErrorResponse(errcode.ErrorDeviceUpperLimit)
		return
	}
	//创建设备
	svc.UpdateUserDevCount(param.UserID, userInfo.Devcount+1) //用户设备数加1
	err = svc.CreateDevice(&param)
	//创建设备失败
	if err != nil {
		global.Logger.Errorf(c, "CreateDevice svc.CreateDevice err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeviceCreateFail)
		return
	}
	//创建成功返回
	response.ToResponse(gin.H{
		"code": http.StatusOK,
		"msg":  "创建设备成功",
		"data": "",
	})
}
