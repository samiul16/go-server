package utils

import (
	"encoding/json"
	"go-server/database"
	"net/http"
)

func SendProducts(w http.ResponseWriter, products []database.Product) {
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}