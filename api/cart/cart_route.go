package cart

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes for the module.
func AddRoutes(r *gin.Engine) {
	cart := r.Group("/api/v1/cart")
	{
		cart.POST("/:cartId", AddItemToCart)
		cart.GET("/:cartId", GetItemsByCartId)
		cart.POST("/cart/:cartId/book/:bookId", AddQuantityFromCart)
		cart.POST("/cart/:cartId/book/:bookId", RemoveQuantityFromCart)

	}
}
