package model

import "github.com/jinzhu/gorm"

//Device 设备表模型
type Device struct {
	*Model
	UserId  uint32 `json:"user_id"`
	Name    string `json:"name"`
	Place   string `json:"place"`
	Running uint8  `json:"running"`
}

//TableName 获取设备表的表名
func (d *Device) TableName() string {
	return "Device"
}

// Create 创建设备
func (d *Device) Create(db *gorm.DB) error {
	return db.Create(d).Error
}

// Update 更新设备信息
func (d *Device) Update(db *gorm.DB, values interface{}) error {
	return db.Model(d).Where("id = ? AND is_del = ?", d.ID, 0).Update(values).Error
}

//Delete 删除设备
func (d *Device) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", d.ID, 0).Delete(d).Error
}

//Count 查询设备的总数
func (d *Device) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(d).Where("user_id = ? AND is_del = ?", d.UserId, 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

//List 查询用户设备信息
func (d *Device) List(db *gorm.DB, pageOffset, pageSize int) ([]*Device, error) {
	var devices []*Device
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	//查找有的数据
	err = db.Where("user_id = ?AND is_del = ?", d.UserId, 0).Find(&devices).Error
	if err != nil {
		return nil, err
	}
	return devices, nil
}
