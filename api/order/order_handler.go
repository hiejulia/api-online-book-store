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

// GetAllOrdersByUser godoc
// @Summary Get all orders by user
// @Description Get user all/history orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /orders [post]
func GetAllOrdersByUser(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	userId := c.Param("userId")
	ordersResp := map[string][]models.OrderItem{}
	order := models.Order{UserID: userId}
	orders := make([]models.Order, 0)
	if err := db.Find(&order, &orders); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	for _, ord := range orders {
		orderId := ord.ID
		orderItem := models.OrderItem{OrderID: orderId}
		orderItems := make([]models.OrderItem, 0)
		if err := db.Find(&orderItem, &orderItems); err != nil {
			common.Error(c, http.StatusInternalServerError, err)
			return
		}
		ordersResp[orderId] = orderItems
	}

	common.SuccessJSON(c, ordersResp)
}

// CreateOrder godoc
// @Summary User create an order
// @Description user create an order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	cartId := c.Param("cartId")
	userId := c.Param("userId")

	cartItem := models.CartItem{CartID: cartId}
	cartItems := make([]models.CartItem, 0)
	if err := db.Find(&cartItem, &cartItems); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	totalPrice := 0.0
	for _, item := range cartItems {
		book := models.Book{ID: item.BookID}
		if err := db.First(&book); err != nil {
			common.Error(c, http.StatusInternalServerError, err)
			return
		}

		totalPrice += book.Price * float64(item.Qty)
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
