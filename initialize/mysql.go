package initialize

import (
	"fmt"
	"log"
	"pAppServer/global"
	"pAppServer/models/table"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func runMysql() {
	mc := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@%s", mc.UserName, mc.PassWord, mc.DSN)
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
	global.DB = db
	db.AutoMigrate(&table.CUserAuth{})
	db.AutoMigrate(&table.CUser{})
	// a := table.CUserAuth{
	// 	IdentityType: 2,
	// 	Identifier:   "1111",
	// }
	// r := db.First(&a).Error
	// fmt.Println("+++++", r, a)
	// cauth := table.CUserAuth{
	// 	IdentityType: constant.WX,
	// 	Identifier:   "",
	// }
	// db.AutoMigrate(&table.CUserAuth{})
	// if err := db.First(&cauth).Error; err != nil {
	// 	fmt.Println("用户不存在 创建", err.Error())
	// 	cauth.CUserUUID = uuid.New().String()
	// 	db.Create(&cauth)
	// } else {
	// 	fmt.Println("用户存在", cauth)
	// cauth := table.CUserAuth{
	// 	IdentityType: constant.WX,
	// 	Identifier:   "",
	// }
	// cauth.CUserUUID = "123456"
	// r := db.Create(&table.CUserAuth{
	// 	IdentityType: 2,
	// 	Identifier:   "",
	// 	CUserUUID:    "123456",
	// }).Error
	// fmt.Print("++++++++++", r)
	// }
	// db.AutoMigrate(&table.CUser{})
	// db.AutoMigrate(&models.ShopStaff{})
	// db.Omit(clause.Associations)
	// errors.Is()
	// db.Model().Preload()
	// db.Create(&models.Shop{
	// 	Address: "区",
	// 	Model: gorm.Model{
	// 		ID: 222,
	// 	},
	// })
	// var s models.Shop
	// var s1 models.Shop
	// db.First(&s, 222)
	// db.First(&s, "tels = ?", "dasdad")
	// db.First(&s1, 333)
	// fmt.Println(s)
	// fmt.Println(s1)
	// db.Model(&models.Shop{
	// 	Model: gorm.Model{
	// 		ID: 5555555,
	// 	},
	// }).Update("Id", "3123")
}
