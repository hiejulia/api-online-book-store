package rating

import (
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"net/http"
)

// PostBook godoc
// @Summary Create a book
// @Description Create a book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Failure 400
// @Router /books [post]
func RateBook(c *gin.Context) {
	req := new(RateBookRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}
	db := c.MustGet("db").(*clients.SQL)

	// User only can review a book if they have already purchased it
	// search for all orders of user_id and book_id
	// order_id

	//
	//select id from order_items where order_id IN (
	//	select id from orders where user_id = userId
	//)
	//AND book_id = bookID
	//

	// count < 1 ->

	// else ->

	// select count (id) from rating where user_id = userId AND book_id = bookId;

	rating := models.Rating{
		ID:     utils.ID(),
		UserID: req.UserId,
		BookID: req.BookId,
		Rating: req.Rating,
	}

	err := db.Create(&rating)
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	common.SuccessJSON(c, &rating)
}
