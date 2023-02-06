package initialize

import (
	"fmt"
	"log"
	"pAppServer/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func runMysql() {
	mc := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@%s", mc.Username, mc.Password, mc.DSN)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("mysql error: %s", err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("mysql error: %s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	global.Db = db
}
