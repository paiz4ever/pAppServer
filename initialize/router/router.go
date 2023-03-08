package router

import (
	"net/http"
	v1 "pAppServer/api/v1"
	"pAppServer/middleware"

	"github.com/gin-gonic/gin"
)

type Test struct {
	GoodsId string `form:"goodsId" json:"goodsi" binding:"required"`
}

func InitRounter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	csm := r.Group("/csm")
	{
		csm.POST("/login", new(v1.CUserApi).Login)
		csm.Use(middleware.JwtAuth())
		csm.POST("/upload", new(v1.CUserApi).Upload)
		csm.POST("/editinfo", new(v1.CUserApi).EditInfo)
		csm.POST("/bindphone", new(v1.CUserApi).BindPhone)
		csm.GET("/socket/chat", hanleChatSocket)
	}

	{
		// consumer.GET("/collect", func(c *gin.Context) {
		// 	var param Test
		// 	err := c.ShouldBind(&param)
		// 	fmt.Println(err, param.GoodsId)
		// 	c.String(200, "ok")
		// })
		// consumer.GET("/collect1/:goodsId", func(c *gin.Context) {
		// 	var param Test
		// 	err := c.ShouldBind(&param)
		// 	fmt.Println(err, param.GoodsId, c.Param("goodsId"), c.PostForm("goodsId"), c.PostForm("goodsi"))
		// 	c.String(200, "ok")
		// })
		// consumer.PUT("/collect2/:goodsId", func(c *gin.Context) {
		// 	fmt.Println(c.ContentType())
		// 	var param Test
		// 	err := c.ShouldBind(&param)
		// 	fmt.Println(err, param.GoodsId, c.Param("goodsId"), c.PostForm("goodsId"), c.PostForm("goodsi"))
		// 	c.String(200, "ok")
		// })
		// consumer.POST("/collect3/:goodsId/:ppp", func(c *gin.Context) {
		// 	fmt.Println(c.ContentType())
		// 	var param Test
		// 	err := c.ShouldBind(&param)
		// 	fmt.Println(err, param.GoodsId, c.Param("goodsId"), c.Param("ppp"), c.PostForm("goodsId"), c.PostForm("goodsi"))
		// 	c.String(200, "ok")
		// })
		// consumer.POST("/collect4", func(c *gin.Context) {
		// 	fmt.Println(c.PostForm("goodsId"), c.PostForm("goodsi"))
		// 	c.String(200, "ok")
		// })
	}

}
