package config

type Config struct {
	Server `mapstructure:"server"`
	Mysql  `mapstructure:"mysql"`
	Redis  `mapstructure:"redis"`
}

// 服务配置
type Server struct {
	Post int32 `mapstructure:"post"`
}

// MySQL配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DSN      string `mapstructure:"dsn"`
}

// Redis配置
type Redis struct {
	Addr string `mapstructure:"addr"`
}
