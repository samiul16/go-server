package handlers

import (
	"fmt"
	"go-server/database"
	"go-server/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)

	if err != nil {
		fmt.Println("wrong product id")
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			utils.SendProducts(w, database.ProductList)
		}
	}

}