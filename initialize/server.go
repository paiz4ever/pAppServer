package initialize

import (
	"fmt"
	"pAppServer/global"
	"pAppServer/middleware"

	"github.com/gin-gonic/gin"
)

func runServer() {
	r := gin.Default()

	r.Use(middleware.Cors())

	consumer := r.Group("/consumer")
	consumer.GET("/test0", func(c *gin.Context) {
		c.String(200, "hello world0000")
	})
	consumer.GET("/test1", func(c *gin.Context) {
		c.String(200, "hello world1111")
	})
	consumer.Use(middleware.JwtAuth())
	consumer.GET("/test2", func(c *gin.Context) {
		c.String(200, "hello world2222")
	})

	// business := r.Group("/business")

	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
