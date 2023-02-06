package initialize

import (
	"fmt"
	"pAppServer/global"

	"github.com/gin-gonic/gin"
)

func runServer() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Post))
}
