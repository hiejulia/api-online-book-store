package main

import (
	"flag"
	"fmt"
	"github.com/hiejulia/api-online-book-store/api"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// @title
// @version 1.0
// @description

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func main() {
	// Env
	env := ".env"
	fenv := flag.String("env", env, "environment variable definition file location")
	flag.Parse()
	if *fenv != "" {
		if !utils.FileExists(*fenv) {
			fmt.Println(*fenv, "does not exist")
			os.Exit(1)
		}
		if err := godotenv.Load(*fenv); err != nil {
			panic(err)
		}
	}

	// DB
	var cfg clients.Config

	time.Sleep(5 * time.Second) // Delay for clients

	cfg = clients.RedisConfig()
	cache := clients.NewRedis(cfg)
	if err := cache.Open(); err != nil {
		panic(err)
	}
	defer cache.Close()

	cfg = clients.SQLConfig()
	db := clients.NewSQL(cfg)
	if err := db.Open(); err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate DB tables
	if err := db.AutoMigrate(
		models.User{},
		models.Book{},
		models.Cart{},
		models.CartItem{},
		models.Order{},
		models.OrderItem{},
	); err != nil {
		panic(err)
	}

	clients.SetCache(cache)
	clients.SetDB(db)

	// API
	if err := api.Run(); err != nil {
		panic(err)
	}

}
