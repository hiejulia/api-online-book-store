package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"net/http"
)

// GetCartById
func GetItemsByCartId(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	cartId := c.Param("cartId")

	cartItem := models.CartItem{CartID: cartId}
	cartItems := make([]models.CartItem, 0)
	if err := db.Find(&cartItem, &cartItems); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	common.SuccessJSON(c, cartItems)
}

// AddItemToCart
func AddItemToCart(c *gin.Context) {
	req := new(CartItemRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}
	cartId := c.Param("cartId")
	db := c.MustGet("db").(*clients.SQL)

	if cartId == "" {
		cartId = utils.ID()
	}

	cart := models.Cart{ID: cartId}
	carts := make([]models.Cart, 0)
	if err := db.Find(&cart, &carts); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	if len(carts) == 0 {
		err := db.Create(&cart)
		if err != nil {
			common.Error(c, http.StatusInternalServerError, err)
			return
		}
	}

	cartItem := models.CartItem{
		ID:     utils.ID(),
		CartID: cartId,
		BookID: req.BookID,
		Qty:    req.Qty,
	}
	err := db.Create(&cartItem)
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	common.SuccessJSON(c, cartId)
}

// AddQuantityFromCart
func AddQuantityFromCart(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	cartId := c.Param("cartId")
	bookId := c.Param("bookId")

	cartItem := models.CartItem{CartID: cartId, BookID: bookId}
	cartItems := make([]models.CartItem, 0)
	if err := db.Find(&cartItem, &cartItems); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	cartItem.Qty += 1
	if err := db.Update(&cartItem); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	common.SuccessJSON(c, cartItems)
}

// RemoveQuantityFromCart
func RemoveQuantityFromCart(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	cartId := c.Param("cartId")
	bookId := c.Param("bookId")

	cartItem := models.CartItem{CartID: cartId, BookID: bookId}
	cartItems := make([]models.CartItem, 0)
	if err := db.Find(&cartItem, &cartItems); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	cartItem.Qty -= 1
	if cartItem.Qty == 0 {
		if err := db.Delete(&cartItem); err != nil {
			common.SuccessJSON(c, cartId)
			return
		}
	}
	if err := db.Update(&cartItem); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	common.SuccessJSON(c, cartItems)
}
