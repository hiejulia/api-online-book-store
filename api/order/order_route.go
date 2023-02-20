package order

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes for the module.
func AddRoutes(r *gin.Engine) {
	order := r.Group("/api/v1/orders")
	{
		order.GET("/user/:userId", GetAllOrdersByUser) // Get user history orders
		order.POST("/cart/:cartId", CreateOrder)
	}
}
