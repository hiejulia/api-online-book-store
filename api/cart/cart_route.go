package cart

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes for the module.
func AddRoutes(r *gin.Engine) {
	cart := r.Group("/api/v1/cart")
	{
		cart.POST("/:cartId/users/:userId", AddItemToCart)
		cart.GET("/:cartId/users/:userId", GetItemsByCartId)

	}
}
