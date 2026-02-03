package utils

import (
	"encoding/json"
	"go-server/repo"
	"net/http"
)

func SendProducts(w http.ResponseWriter, products interface{}) {
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func SendProduct(w http.ResponseWriter, product repo.Product) {
	err := json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func SendData(w http.ResponseWriter, data interface{} ) {
 err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}