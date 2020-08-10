package service

import "github.com/lee820/ServerIOT/internal/model"

//CreateDeviceRequest 创建设备请求
type CreateDeviceRequest struct {
	UserID      uint32 `form:"userid" binding:"required,gte=1"`
	DeviceName  string `form:"devicename" binding:"min=3,max=32"`
	DevicePlace string `form:"deviceplace" binding:"min=1,max=32"`
	Running     uint8  `form:"running" bingding:"oneof= 0 1"`
}

//UpdateDeviceRequest 更新设备信息请求
type UpdateDeviceRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	DeviceName string `form:"devicename" binding:"min=3,max=32"`
	Running    uint8  `form:"running" bingding:"oneof= 0 1"`
}

//DeleteDeviceRequest 删除设备请求
type DeleteDeviceRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//CountDeviceRequest 查询设备总数请求
type CountDeviceRequest struct {
	UserID uint32 `form:"userid" binding:"required,gte=1"`
}

//ListDeviceRequest 查询用户设备信息请求
type ListDeviceRequest struct {
	UserID   uint32 `form:"userid" binding:"required,gte=1"`
	Page     int    `form:"page" binding:"required,gte=0"`
	PageSize int    `form:"pagesize" binding:"required, gte=1"`
}

//CreateDevice service层处理创建设备请求
func (svc *Service) CreateDevice(param *CreateDeviceRequest) error {
	return svc.dao.CreateDevice(param.UserID, param.DeviceName, param.DevicePlace, param.Running)
}

//UpdateDevice service层处理更新设备信息请求
func (svc *Service) UpdateDevice(param *UpdateDeviceRequest) error {
	return svc.dao.UpdateDevice(param.ID, param.DeviceName, param.Running)
}

//DeleteDevice service层处理删除设备请求
func (svc *Service) DeleteDevice(param *DeleteDeviceRequest) error {
	return svc.dao.DeleteDevice(param.ID)
}

//CountDevice service层处理统计设备数量请求
func (svc *Service) CountDevice(param *CountDeviceRequest) (int, error) {
	return svc.dao.CountDevice(param.UserID)
}

//ListDevice service层处理查询设备信息请求
func (svc *Service) ListDevice(param *ListDeviceRequest) ([]*model.Device, error) {
	return svc.dao.ListDevices(param.UserID, param.Page, param.PageSize)
}
