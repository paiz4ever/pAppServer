package initialize

import (
	"pAppServer/global"

	"github.com/go-redis/redis/v8"
)

func runRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: global.Config.Redis.Addr,
	})
	global.Rdb = rdb
}
