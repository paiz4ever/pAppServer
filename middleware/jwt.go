package middleware

import (
	"pAppServer/response"
	"pAppServer/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		if !utils.VerifyToken(token) {
			response.Failed("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
