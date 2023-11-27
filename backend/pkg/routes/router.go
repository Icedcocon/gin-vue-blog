package routes

import (
	"backend/pkg/config"
	"backend/pkg/dao"
	"backend/pkg/utils"
	"fmt"
	"net/http"
	"time"
)

// 初始化全局变量
func InitGlobalVariable() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitGLogger()
	// 初始化数据库 DB
	dao.DB = utils.InitMariaDB() // 需要先导入 gvb.sql
	// dao.DB = utils.InitSQLiteDB("gorm.db") // TODO: 默认无数据，暂时无法使用
	// 初始化 Redis
	utils.InitRedis()
}

// 后台服务
func BackendServer() *http.Server {
	logger := utils.GLogger

	backPort := config.GlobalConfig.SERVER.BackPort
	host := config.GlobalConfig.SERVER.Host
	logger.Infof("后台服务启动于 %d 端口", backPort)
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, backPort),
		Handler:      BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
