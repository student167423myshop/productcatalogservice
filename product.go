package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Products struct {
	Products []Product `json:"Products"`
}

type Product struct {
	ProductId   string   `json:"ProductId,omitempty"`
	Name        string   `json:"Name,omitempty"`
	Description string   `json:"Description"`
	PictureUrl  string   `json:"PictureUrl"`
	Price       Price    `json:"Price,omitempty"`
	Categories  []string `json:"Categories"`
}

type Price struct {
	Units int `json:"Units"`
	Nanos int `json:"Nanos"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json, err := ioutil.ReadFile("db/products.json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
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
		if productId == p.ProductId {
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
