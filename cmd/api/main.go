package main

import (
	"golang-merchant/apps/auth"
	"golang-merchant/apps/product"
	"golang-merchant/apps/transaction"
	"golang-merchant/external/database"
	infrafiber "golang-merchant/infra/fiber"
	"golang-merchant/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	router.Use(infrafiber.Trace())

	auth.Init(router, db)
	product.Init(router, db)
	transaction.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
