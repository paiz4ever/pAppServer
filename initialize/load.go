package initialize

import (
	"fmt"
	"log"
	"pAppServer/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 加载配置文件
func loadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error resources file: %s \n", err.Error())
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("unable to decode into struct %s \n", err.Error())
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO 重新加载配置
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}
