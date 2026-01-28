package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/database"
	"go-server/utils"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, request *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid JSON", 400)
		return
	}

	database.Store(newProduct)
	utils.SendProducts(w, database.List())

}
