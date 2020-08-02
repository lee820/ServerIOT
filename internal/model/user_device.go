package model

//UserDevice 用户和设备关联表模型
type UserDevice struct {
	*Model
	UserID   uint32 `json:"user_id"`
	DeviceID uint32 `json:"device_id"`
}

//TableName 获取用户设备关联表的表名
func (u UserDevice) TableName() string {
	return "user_device"
}
