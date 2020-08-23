package global

import (
	"github.com/lee820/ServerIOT/pkg/logger"
	"github.com/lee820/ServerIOT/pkg/setting"
	"github.com/opentracing/opentracing-go"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	Logger          *logger.Logger
	Tracer          opentracing.Tracer
)
