package service

import (
	"context"

	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/dao"
)

//Service service层业务结构体
type Service struct {
	ctx context.Context
	dao *dao.Dao
}

//New service层创建业务结构体
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
