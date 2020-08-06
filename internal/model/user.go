package model

import "github.com/jinzhu/gorm"

//User 用户模型
type User struct {
	*Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

//TableName 获取用户表的表名
func (u User) TableName() string {
	return "user"
}

//Create 创建用户
func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

//Update 更新用户信息
func (u User) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&u).Where("id = ?", u.ID).Update(values).Error
}

//Delete 删除用
func (u User) Delete(db *gorm.DB) error {
	return db.Where("id = ?", u.ID).Delete(&u).Error
}

//Query 使用手机号查询用户是否存在
func (u User) Query(db *gorm.DB, param interface{}) (User, error) {
	var usr User
	//使用结构体，gorm自动查询有值的字段
	err := db.Where(param).Find(&usr).Error
	return usr, err
}
