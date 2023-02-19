package order

import (
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"net/http"
)

// Move to a service

// GetAllOrdersByUser Get all history orders of a user
func GetAllOrdersByUser(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	userId := c.Param("userId")

	item := models.Order{UserID: userId}
	items := make([]models.Order, 0)
	if err := db.Find(&item, &items); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	common.SuccessJSON(c, items)
}

// CreateOrder
func CreateOrder(c *gin.Context) {
	userId := c.MustGet("user").(*models.User).ID
	db := c.MustGet("db").(*clients.SQL)
	cartId := c.Param("cartId")

	cartItem := models.CartItem{ID: cartId}
	cartItems := make([]models.CartItem, 0)
	if err := db.Find(&cartItem, &cartItems); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	totalPrice := 0.0
	for _, item := range cartItems {
		book := models.Book{ID: item.BookID}
		books := make([]models.Book, 0)
		if err := db.Find(&book, &books); err != nil {
			common.Error(c, http.StatusInternalServerError, err)
			return
		}
		totalPrice += books[0].Price * float64(item.Qty)
	}

	order := models.Order{
		ID:     utils.ID(),
		UserID: userId,
		Price:  totalPrice,
		Status: "Ordered",
	}
	if err := db.Create(&order); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	for _, item := range cartItems {
		orderItem := models.OrderItem{
			ID:      utils.ID(),
			OrderID: order.ID,
			BookID:  item.BookID,
			Qty:     item.Qty,
		}
		if err := db.Create(&orderItem); err != nil {
			common.Error(c, http.StatusInternalServerError, err)
			return
		}
	}

	common.SuccessJSON(c, order)
}
