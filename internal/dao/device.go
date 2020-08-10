package dao

import (
	"github.com/lee820/ServerIOT/internal/model"
	"github.com/lee820/ServerIOT/pkg/app"
)

//CreateDevice dao层创建设备
func (d *Dao) CreateDevice(userId uint32, name, place string, status uint8) error {
	dev := model.Device{
		UserId:  userId,
		Name:    name,
		Place:   place,
		Running: status,
	}
	return dev.Create(d.engine)
}

//UpdateDevice dao层更新设备信息
func (d *Dao) UpdateDevice(id uint32, name string, status uint8) error {
	dev := model.Device{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"running": status,
	}
	if name != "" {
		values["name"] = name
	}
	return dev.Update(d.engine, values)
}

//DeleteDevice 根据设备id 删除设备
func (d *Dao) DeleteDevice(id uint32) error {
	dev := model.Device{
		Model: &model.Model{ID: id},
	}
	return dev.Delete(d.engine)
}

//CountDevice dao层查询用户设备总数
func (d *Dao) CountDevice(userId uint32) (int, error) {
	dev := model.Device{
		UserId: userId,
	}
	return dev.Count(d.engine)
}

//ListDevices dao层，根据用户id查询设备
func (d *Dao) ListDevices(userId uint32, page, pageSize int) ([]*model.Device, error) {
	dev := model.Device{
		UserId: userId,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return dev.List(d.engine, pageOffset, pageSize)
}
