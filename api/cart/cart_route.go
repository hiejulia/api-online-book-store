package cart

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes for the module.
func AddRoutes(r *gin.Engine) {
	cart := r.Group("/cart")
	{
		cart.POST("/:cartId", AddItemToCart)
		cart.GET("/:cartId", GetItemsByCartId)
		cart.GET("/cart/:cartId/book/:bookId", AddQuantityFromCart)
		cart.GET("/cart/:cartId/book/:bookId", RemoveQuantityFromCart)

	}
}
