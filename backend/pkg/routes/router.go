package routes

import (
	"backend/pkg/config"
	"backend/pkg/utils"
	"fmt"
	"net/http"
	"time"
)

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
