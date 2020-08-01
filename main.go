package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/model"
	"github.com/lee820/ServerIOT/pkg/logger"
	"github.com/lee820/ServerIOT/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Printf("init setting error: %v", err)

	}

	/*
		err = setupDBEngine()
		if err != nil {
			fmt.Printf("init db fail. %v", err)
		}
	*/
	err = setupLogger()
	if err != nil {
		fmt.Printf("init log fail err: %v", err)
	}
}

// @title 物联网后端操作系统
// @version v1.0
// description lee
//termsOfService https://github.com//lee820
func main() {
	global.Logger.Infof("%s: %d", "lee", 666)
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true},
		"", log.LstdFlags).WithCaller(2)
	return nil
}
