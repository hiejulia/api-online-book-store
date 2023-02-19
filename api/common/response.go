package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response message format.
type Response struct {
	Error  string      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// Error response.
func Error(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, &Response{
		Error: err.Error(),
	})
}

// SuccessJSON response message.
func SuccessJSON(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, &Response{
		Result: v,
	})
}
