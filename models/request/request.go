package request

import (
	"pAppServer/models/response"

	"github.com/gin-gonic/gin"
)

func Vertify(param interface{}, c *gin.Context) error {
	if err := c.ShouldBind(param); err != nil {
		response.Failed("illegal param", c)
		return err
	}
	return nil
}
