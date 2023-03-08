package config

type Config struct {
	App    `mapstructure:"app"`
	Server `mapstructure:"server"`
	Mysql  `mapstructure:"mysql"`
	Redis  `mapstructure:"redis"`
	Jwt    `mapstructure:"jwt"`
	Mail   `mapstructure:"mail"`
	WX     `mapstructure:"wx"`
}

// app配置
type App struct {
	Name string `mapstructure:"name"`
}

// 服务配置
type Server struct {
	Port int32 `mapstructure:"port"`
}

// MySQL配置
type Mysql struct {
	UserName string `mapstructure:"username"`
	PassWord string `mapstructure:"password"`
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

// 邮箱配置
type Mail struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	UserName  string `mapstructure:"username"`
	Alias     string `mapstructure:"alias"`
	PassWord  string `mapstructure:"password"`
	CloseTime int    `mapstructure:"closeTime"`
}

// 邮箱配置
type WX struct {
	AppId  string `mapstructure:"appid"`
	Secret string `mapstructure:"secret"`
}
