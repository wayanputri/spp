package main

import (
	"project/app/config"
	"project/app/database"
	"project/app/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.InitConfig()
	postgres := database.InitPostgres(cfg)
	app := fiber.New()
	router.InitRouter(app, postgres)
	app.Listen(":8080")
}
