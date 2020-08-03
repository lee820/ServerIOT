package dao

import "github.com/jinzhu/gorm"

//Dao 数据方位接口
type Dao struct {
	engine *gorm.DB
}

//New 新建Dao实例
func New(engin *gorm.DB) *Dao {
	return &Dao{engine: engin}
}
