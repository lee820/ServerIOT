package main

import (
	"fmt"
	"time"

	"github.com/lee820/ServerIOT/global"
	"github.com/lee820/ServerIOT/internal/model"
	"github.com/lee820/ServerIOT/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Printf("init setting error: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		fmt.Printf("init db fail. %v", err)
	}
}

func main() {
	fmt.Printf("run mode: %s", global.ServerSetting.RunMode)
	fmt.Printf("log save path: %s", global.AppSetting.LogSavePath)
	fmt.Printf("db host: %s", global.DatabaseSetting.Host)
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
