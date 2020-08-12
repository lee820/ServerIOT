package dao

import "github.com/lee820/ServerIOT/internal/model"

//GetAuth dao层获取认证密码
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
