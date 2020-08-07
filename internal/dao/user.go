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

//UpdateUserName dao层更新用户名
func (d *Dao) UpdateUserName(id uint32, name string) error {
	usr := model.User{
		Name:  name,
		Model: &model.Model{ID: id},
	}
	return usr.Update(d.engine, name)
}

//UpdateUserPassword dao层更新用户密码
func (d *Dao) UpdateUserPassword(id uint32, pwd string) error {
	usr := model.User{
		Password: pwd,
		Model:    &model.Model{ID: id},
	}
	return usr.Update(d.engine, pwd)
}

//GetUserInfo dao层使用手机号查询用户信息
func (d *Dao) GetUserInfoByPhone(phone string) (model.User, error) {
	user := model.User{
		Phone: phone,
	}
	return user.Query(d.engine)
}

//GetUserInfo dao层用户名查询用户信息
func (d *Dao) GetUserInfoByUserName(name string) (model.User, error) {
	user := model.User{
		Name: name,
	}
	return user.Query(d.engine)
}
