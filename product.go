package main

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
