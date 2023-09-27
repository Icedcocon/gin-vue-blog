package config

// 代码废弃

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ConfigInstance *viper.Viper

func Init(filename string) error {
	// 读取配置文件
	viper.SetConfigFile(filename)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// 验证配置文件
	if err := validateConfig(viper.GetViper()); err != nil {
		return err
	}

	// 配置文件映射到结构体
	viper.Unmarshal(&GlobalConfig)
	ConfigInstance = viper.GetViper()

	// 注册回调函数，配置变化时触发，更新全局配置变量 GlobalConfig
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.Unmarshal(&GlobalConfig)
	})

	// 开启后台线程监听配置变化
	viper.WatchConfig()
	return nil
}

func GetConfig() *Configuration {
	return &GlobalConfig
}

func validateConfig(viper *viper.Viper) error {
	return nil
}
