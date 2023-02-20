package book

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	book := r.Group("/api/v1/books")
	{
		book.GET("", GetAllBooks)
		book.POST("", PostBook)
	}
}
