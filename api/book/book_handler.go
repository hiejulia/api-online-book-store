package book

import (
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	db := c.MustGet("db").(*clients.SQL)
	book := models.Book{}
	books := make([]models.Book, 0)
	if err := db.Find(&book, &books); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	common.SuccessJSON(c, &books)
}

func PostBook(c *gin.Context) {
	req := new(BookRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}
	db := c.MustGet("db").(*clients.SQL)
	book := models.Book{
		ID:    utils.ID(),
		Name:  req.Name,
		Price: req.Price,
		Qty:   req.Qty,
	}

	err := db.Create(&book)
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	common.SuccessJSON(c, &book)
}
