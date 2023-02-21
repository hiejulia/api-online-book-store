package clients

import (
	"flag"
	"fmt"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func SetUpMain() {
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
	var cfg Config

	time.Sleep(5 * time.Second) // Delay for clients

	cfg = RedisConfig()
	cache := NewRedis(cfg)
	if err := cache.Open(); err != nil {
		panic(err)
	}
	defer cache.Close()

	cfg = SQLConfig()
	db := NewSQL(cfg)
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
		models.Rating{},
	); err != nil {
		panic(err)
	}

	SetCache(cache)
	SetDB(db)
}
