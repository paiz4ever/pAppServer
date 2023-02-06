package config

type Config struct {
	Server `mapstructure:"server"`
	Mysql  `mapstructure:"mysql"`
	Redis  `mapstructure:"redis"`
	Jwt    `mapstructure:"jwt"`
}

// 服务配置
type Server struct {
	Port int32 `mapstructure:"port"`
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

// Jwt配置
type Jwt struct {
	SignedKey string `mapstructure:"signedKey"`
	MaxAge    int32  `mapstructure:"maxAge"`
}
