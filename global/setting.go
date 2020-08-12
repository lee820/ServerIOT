package global

import (
	"github.com/lee820/ServerIOT/pkg/logger"
	"github.com/lee820/ServerIOT/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
