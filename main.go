package main

import (
	"github.com/hiejulia/api-online-book-store/api"
	"github.com/hiejulia/api-online-book-store/clients"
)

// @title          	Book store shop
// @version         1.0
// @description     A book online store service API in Go using Gin framework.
// @termsOfService

// @contact.name   Hien Nguyen
// @contact.url
// @contact.email

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
func main() {
	clients.SetUpMain()
	// API
	if err := api.Run(); err != nil {
		panic(err)
	}

}
