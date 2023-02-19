package user

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes for the module.
func AddRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/register", PostPublicRegister)
		user.POST("/login", PostPublicLogin)
	}
}
