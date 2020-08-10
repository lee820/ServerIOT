package dao

import (
	"github.com/lee820/ServerIOT/internal/model"
)

//CreateUser dao层创建用户数据
func (d *Dao) CreateUser(name, password, phone string) error {
	usr := model.User{
		Name:     name,
		Password: password,
		Phone:    phone,
	}
	return usr.Create(d.engine)
}

//UpdateUser dao层更新用户信息
func (d *Dao) UpdateUser(id uint32, name, pwd string, devCount int) error {
	usr := model.User{
		Model: &model.Model{ID: id},
	}

	newInfo := map[string]interface{}{}
	if name != "" {
		newInfo["name"] = name
	}
	if pwd != "" {
		newInfo["password"] = pwd
	}
	if devCount != -1 {
		newInfo["dev_count"] = devCount
	}
	return usr.Update(d.engine, newInfo)
}

//GetUserInfo dao层查询用户信息
func (d *Dao) GetUserInfo(id uint32, name, phone string) (model.User, error) {
	user := model.User{}
	if id != 0 {
		user.Model = &model.Model{ID: id}
	}
	if name != "" {
		user.Name = name
	}
	if phone != "" {
		user.Phone = phone
	}
	return user.Query(d.engine)
}
