package initialize

import (
	"pAppServer/manager"
)

func Run() {
	// 读取配置
	loadConfig()
	// 链接数据库
	runMysql()
	// 链接redis
	runRedis()
	// 启动各个管理进程
	manager.Run()
	// 启动服务
	runServer()
}
