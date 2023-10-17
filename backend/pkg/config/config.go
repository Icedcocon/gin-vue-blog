package config

var GlobalConfig Configuration

type Configuration struct {
	DB      DBConfig      `mapstructure:"DB"`
	LOG     LogConfig     `mapstructure:"LOG"`
	SERVER  ServerConfig  `mapstructure:"SERVER"`
	UPLOAD  UpLoadConfig  `mapstructure:"UPLOAD"`
	SESSION SessionConfig `mapstructure:"SESSION"`
	REDIS   RedisConfig   `mapstructure:"REDIS"`
	JWT     JWTConfig     `mapstructure:"JWT"`
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
	Level    string `yaml:"LEVEL"`
	Prefix   string `yaml:"PREFIX"`
	FileName string `yaml:"FILENAME"`
}

type ServerConfig struct {
	Mode     string `yaml:"MODE"`
	Host     string `yaml:"HOST"`
	BackPort int    `yaml:"BACKPORT"`
}

type UpLoadConfig struct {
	OssType     string `yaml:"OSSTYPE"`
	Path        string `yaml:"PATH"`
	StorePath   string `yaml:"STOREPATH"`
	MdPath      string `yaml:"MDPATH"`
	MdStorePath string `yaml:"MDSTOREPATH"`
}

type SessionConfig struct {
	Name   string `yaml:"NAME"`
	Salt   string `yaml:"SALT"`
	MaxAge int    `yaml:"MAXAGE"`
}

type RedisConfig struct {
	Addr     string `yaml:"ADDR"`
	Password string `yaml:"PASSWORD"`
	DB       int    `yaml:"DB"`
}

type JWTConfig struct {
	Secret string `yaml:"SECRET"`
	Expire int    `yaml:"EXPIRE"`
	Issuer string `yaml:"ISSUER"`
}
