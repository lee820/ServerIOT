package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/service"
	"github.com/lee820/ServerIOT/pkg/app"
	"github.com/lee820/ServerIOT/pkg/convert"
	"github.com/lee820/ServerIOT/pkg/errcode"
)

type DevRouter struct{}

//NewDevRouter 创建设备路由
func NewDevRouter() DevRouter {
	return DevRouter{}
}

//CreateDevice 添加设备路由
func (d *DevRouter) CreateDevice(c *gin.Context) {
	param := service.CreateDeviceRequest{}
	response := app.NewResponse(c)
	//参数校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "createDevice bindAndValid ers: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	//查询用户是否存在
	userInfo, err := svc.GetUserInfoByID(param.UID)
	if err != nil {
		global.Logger.Errorf(c, "CreateDevice svc.QueryUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotFound)
		return
	}
	//查询用户设备数量是否到达上限
	if userInfo.DevCount >= global.AppSetting.NormalUserDeviceUplimit {
		//达到上限，不能添加设备，返回
		response.ToErrorResponse(errcode.ErrorDeviceUpperLimit)
		return
	}
	//创建设备
	svc.UpdateUserDevCount(param.UID, userInfo.DevCount+1) //用户设备数加1
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

//UpdateDevice 修改设备信息
func (d *DevRouter) UpdateDevice(c *gin.Context) {
	param := service.UpdateDeviceRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "UpdateDevice qpp.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	//修改设备信息
	err := svc.UpdateDevice(&param)
	if err != nil {
		global.Logger.Errorf(c, "UpdateDevice fail: %v", err)
		response.ToErrorResponse(errcode.ErrorDeviceUpdateFail)
		return
	}

	response.ToResponse(gin.H{
		"code": http.StatusOK,
		"msg":  "创建更新成功",
		"data": "",
	})
}

//DeleteDevice 删除设备
func (d *DevRouter) DeleteDevice(c *gin.Context) {
	param := service.DeleteDeviceRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "DeleteDevice qpp.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteDevice(&param)
	if err != nil {
		global.Logger.Errorf(c, "DeleteDevice fail: %v", err)
		response.ToErrorResponse(errcode.ErrorDeviceDelFail)
		return
	}

	response.ToResponse(gin.H{
		"code": http.StatusOK,
		"msg":  "创建删除成功",
		"data": "",
	})
}

//GetDeviceList 获取用户的设备列表
func (d *DevRouter) GetDeviceList(c *gin.Context) {
	param := service.ListDeviceRequest{UID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "GetDeviceList qpp.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	//pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	//获取设备总数
	totalRows, err := svc.CountDevice(&service.CountDeviceRequest{UID: param.UID})
	if err != nil {
		global.Logger.Errorf(c, "CountDevice fail: %v", err)
		response.ToErrorResponse(errcode.ErrorDeviceCountFail)
		return
	}

	//获取设备列表
	devices, err := svc.ListDevice(&param)
	if err != nil {
		global.Logger.Errorf(c, "ListDevice fail: %v", err)
		response.ToErrorResponse(errcode.ErrorDeviceListFail)
		return
	}

	response.ToResponseList(devices, totalRows)
}
