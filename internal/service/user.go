package service

import "github.com/lee820/ServerIOT/internal/model"

//CreateUserRequest service层创建用户业务接口校验
type CreateUserRequest struct {
	Name     string `form:"username" binding:"min=3,max=16"`
	Password string `form:"password" binding:"min=6,max=16"`
	Phone    string `form:"phone" binding:"len=11"`
}

//UpdateUserNameRequest service层更新用户名业务接口校验
type UpdateUserNameRequest struct {
	ID   uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"min=3,max=16"`
}

//UpdateUserPasswordRequest service层更新用户密码业务接口校验
type UpdateUserPasswordRequest struct {
	ID       uint32 `form:"id" binding:"required,gte=1"`
	Password string `form:"password" binding:"min=6,max=16"`
}

//GetUserInfoRequest 获取用户信息接口校验
type GetUserInfoRequestByPhone struct {
	Phone string `form:"phone" binding:"len=11"`
}

type GetUserInfoByNameRequest struct {
	Username string `form:"username" binding:"min=3,max=16"`
	Password string `form:"password" binding:"min=6,max=16"`
}

//CreateUser service层创建用户
func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Name, param.Password, param.Phone)
}

//UpdateUserName service层更新用户名
func (svc *Service) UpdateUserName(param *UpdateUserNameRequest) error {
	return svc.dao.UpdateUserName(param.ID, param.Name)
}

//UpdateUserPasswordRequest service层更新用户密码
func (svc *Service) UpdateUserPasswordRequest(param *UpdateUserPasswordRequest) error {
	return svc.dao.UpdateUserPassword(param.ID, param.Password)
}

//GetUserInfo service层使用手机号查询用户信息
func (svc *Service) GetUserInfoByPhone(param *GetUserInfoRequestByPhone) (model.User, error) {
	return svc.dao.GetUserInfoByPhone(param.Phone)
}

//GetUserInfo service层使用用户名用户信息
func (svc *Service) GetUserInfoByUsername(param *GetUserInfoByNameRequest) (model.User, error) {
	return svc.dao.GetUserInfoByUserName(param.Username)
}
