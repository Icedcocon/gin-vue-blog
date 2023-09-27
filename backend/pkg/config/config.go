package config

var GlobalConfig Configuration

type Configuration struct {
	DB     DBConfig     `mapstructure:"DB"`
	LOG    LogConfig    `mapstructure:"LOG"`
	SERVER ServerConfig `mapstructure:"SERVER"`
}

type DBConfig struct {
	Dbname   string `yaml:"DBNAME"`   // 数据库名
	Host     string `yaml:"HOST"`     // 服务器地址
	Port     int    `yaml:"PORT"`     // 端口
	Username string `yaml:"USERNAME"` // 数据库用户名
	Password string `yaml:"PASSWORD"` // 数据库密码
	Config   string `yaml:"CONFIG"`   // 高级配置
	LogMode  string `yaml:"LOGMODE"`  // 日志级别
}

type LogConfig struct {
	Level  string `yaml:"LEVEL"`
	Prefix string `yaml:"PREFIX"`
}

type ServerConfig struct {
	Host     string `yaml:"HOST"`
	BackPort int    `yaml:"BACKPORT"`
	Env      string `yaml:"ENV"`
}
