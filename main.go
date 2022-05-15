package main

import (
	"github.com/gofiber/fiber/v2"
)

func getApp() *fiber.App {
	app := fiber.New()
	app.Get("/api/v1/products", getProducts)
	return app
}

func main() {
	app := getApp()
	app.Listen(":3550")
}
