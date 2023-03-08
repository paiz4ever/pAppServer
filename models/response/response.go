package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, "ok", data})
}

func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{http.StatusBadRequest, message, 0})
}
