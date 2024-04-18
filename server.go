package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8000", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := criaEstoque()
	json.NewEncoder(w).Encode(products)
}
