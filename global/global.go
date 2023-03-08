package global

import (
	"pAppServer/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
)
