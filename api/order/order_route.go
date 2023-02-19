package order

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes for the module.
func AddRoutes(r *gin.Engine) {
	order := r.Group("/orders")
	{
		order.GET("/user/:userId", GetAllOrdersByUser)
		order.POST("/cart/:cartId", CreateOrder)
	}
}
