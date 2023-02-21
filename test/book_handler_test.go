package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/auth"
	"github.com/hiejulia/api-online-book-store/api/book"
	"github.com/hiejulia/api-online-book-store/api/user"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllBooks(t *testing.T) {
	clients.SetUpMain()
	Convey("Test TestGetAllBooks succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.GET("/api/v1/books", book.GetAllBooks)

		req, err := http.NewRequest("GET", "/api/v1/books", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var books []models.Book
		json.Unmarshal(w.Body.Bytes(), &books)
		assert.Equal(t, http.StatusOK, w.Code)
		// assert.Empty(t, books)
		So(err, ShouldBeNil)
	})
	Convey("Test TestGetAllBooks fail if header token is not exisit... ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.GET("/api/v1/books", book.GetAllBooks)

		req, err := http.NewRequest("GET", "/api/v1/book", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var books []models.Book
		json.Unmarshal(w.Body.Bytes(), &books)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		// assert.Empty(t, books)
		So(err, ShouldNotBeNil)
	})
}

func TestCreateBook(t *testing.T) {
	Convey("TestCreateBook success", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.POST("/api/v1/books", book.GetAllBooks)

		book := models.Book{
			ID:    utils.ID(),
			Name:  "Book1",
			Price: 10.0,
			Qty:   100,
		}
		jsonValue, _ := json.Marshal(book)
		req, err := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}
