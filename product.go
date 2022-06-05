package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       Price    `json:"price,omitempty"`
	Categories  []string `json:"categories"`
}

type Price struct {
	Units int `json:"units"`
	Nanos int `json:"nanos"`
}

func getMuxProducts(w http.ResponseWriter, r *http.Request) {
	json, _ := ioutil.ReadFile("db/products.json")
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getMuxProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	file, err := ioutil.ReadFile("db/products.json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	products := Products{}
	var product Product
	founded := false
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	for _, p := range products.Products {
		if productId == p.Id {
			product = p
			founded = true
		}
	}

	if !founded {
		w.WriteHeader(http.StatusNotFound)
	}

	json, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
