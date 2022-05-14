package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

func getProducts(c *fiber.Ctx) error {
	file, _ := ioutil.ReadFile("products.json")
	products := Products{}
	_ = json.Unmarshal([]byte(file), &products)
	return c.JSON(products)
}
