package main

import (
	"pembiayaan/features/pembiayaan/SPP/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/api/spp", handler.Get)
	// ADD THIS
	app.Post("/api/spp", handler.Create)
	app.Listen(":8080")
}
