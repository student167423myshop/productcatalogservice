package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/products", getMuxProducts).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/product/{productId}", getMuxProduct).Methods(http.MethodGet)
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
