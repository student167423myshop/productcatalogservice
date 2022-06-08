package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", getProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{productId}", getProduct).Methods(http.MethodGet)
	return r
}

func main() {
	r := getRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":3550",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
