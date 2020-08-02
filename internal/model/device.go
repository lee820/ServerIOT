package model

//Device 设备表模型
type Device struct {
	*Model
	Name    string `json:"name"`
	Place   string `json:"place"`
	Running uint8  `json:"running"`
}

//TableName 获取设备表的表名
func (d Device) TableName() string {
	return "Device"
}
