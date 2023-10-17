package main

import (
	"backend/pkg/config"
	"backend/pkg/utils"
	"fmt"
)

func main() {
	// 初始化配置信息
	utils.InitViper()
	utils.InitGLogger()
	db := utils.InitMariaDB()

	logger := utils.GLogger
	logger.Info("Hello World")
	logger.Infof("Config is %v", config.GlobalConfig)
	logger.Error("Hello World")
	logger.Errorf("Config is %v", config.GlobalConfig)
	logger.Trace("Hello World")
	logger.Tracef("Config is %v", config.GlobalConfig)
	logger.Debug("Hello World")
	logger.Debugf("Config is %v", config.GlobalConfig)
	logger.Warn("Hello World")
	logger.Warnf("Config is %v", config.GlobalConfig)

	fmt.Printf("Config is %v", config.GlobalConfig)
	fmt.Printf("Config is %v", db)
}
