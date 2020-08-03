package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	Notfound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	ErrorCreateUserFail       = NewError(20000001, "创建用户失败")
	ErrorUpdateUserNameFail   = NewError(20000002, "更新用吗名失败")
	ErrorUpdateUserPwdFail    = NewError(20000003, "更新用户密码失败")
	ErrorGetUserInfoFail      = NewError(20000004, "获取用户信息失败")
)
