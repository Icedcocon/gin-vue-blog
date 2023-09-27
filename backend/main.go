package main

import (
	"backend/pkg/config"
	"backend/pkg/utils"
	"fmt"
)

func main() {
	// 初始化配置信息
	utils.InitViper()
	db := utils.InitMariaDB()

	fmt.Printf("Config is %v", config.GlobalConfig)
	fmt.Printf("Config is %v", db)
}
