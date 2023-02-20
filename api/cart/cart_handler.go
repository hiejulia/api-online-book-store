package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"net/http"
)

// GetItemsByCartId godoc
// @Summary Get all items by cartId
// @Description Get all items by cartId
// @Tags carts
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /carts [post]
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

// AddItemToCart godoc
// @Summary User add item to cart
// @Description user add item to cart
// @Tags carts
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /carts [post]
func AddItemToCart(c *gin.Context) {
	req := new(CartItemRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}
	cartId := c.Param("cartId")
	userId := c.Param("userId")

	db := c.MustGet("db").(*clients.SQL)

	if cartId == "new" {
		cartId = utils.ID()
	}

	cart := models.Cart{ID: cartId}
	carts := make([]models.Cart, 0)
	if err := db.Find(&cart, &carts); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	cart.UserID = userId
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
	common.SuccessJSON(c, gin.H{"cartId": cartId, "status": "OK"})
}
