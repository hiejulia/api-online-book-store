package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/auth"
	"github.com/hiejulia/api-online-book-store/api/book"
	"github.com/hiejulia/api-online-book-store/api/cart"
	"github.com/hiejulia/api-online-book-store/api/order"
	"github.com/hiejulia/api-online-book-store/api/user"
	_ "github.com/hiejulia/api-online-book-store/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// Run the API on the provided address.
func Run() (err error) {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	// docs route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// set utils for gin
	auth.SetupMiddleware(r)
	user.SetupPrivacy()

	user.AddRoutes(r)

	r.Use(auth.Authorize)
	book.AddRoutes(r)
	cart.AddRoutes(r)
	order.AddRoutes(r)

	addr := fmt.Sprintf("%s:%s", os.Getenv("API_HOST"), os.Getenv("API_PORT"))
	fmt.Println("Serving at", addr)
	err = r.Run(addr)
	return
}
