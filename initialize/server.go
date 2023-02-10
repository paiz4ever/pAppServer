package initialize

import (
	"fmt"
	"pAppServer/global"
	"pAppServer/initialize/router"
	"pAppServer/middleware"

	"github.com/gin-gonic/gin"
)

func runServer() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Static("/assets", "./assets/img")
	router.InitRounter(r)
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
