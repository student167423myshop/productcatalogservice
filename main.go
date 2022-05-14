package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/api/v1/products", getProducts)
	app.Listen(":3550")
}
