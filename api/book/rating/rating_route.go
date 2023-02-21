package rating

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	book := r.Group("/api/v1/rating")
	{
		book.POST("", RateBook)
	}
}
