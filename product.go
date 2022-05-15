package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       Price      `json:"price,omitempty"`
	Categories  []Category `json:"categories"`
}

type Price struct {
	Units int `json:"units"`
	Nanos int `json:"nanos"`
}

type Category struct {
	Name string
}

func getProducts(c *fiber.Ctx) error {
	file, _ := ioutil.ReadFile("db/products.json")
	products := Products{}
	_ = json.Unmarshal([]byte(file), &products)
	return c.JSON(products)
}
