package product

import (
	"encoding/json"
	"fmt"
	"go-server/domain"
	"go-server/utils"
	"net/http"
)

type ReqProduct struct {
	ID    int     `json:"id"`
	Title  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *Handler) CreateProducts(w http.ResponseWriter, request *http.Request) {

	var newProduct ReqProduct
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid JSON", 400)
		return
	}

	createdProduct, err := h.svc.Create(&domain.Product{
		ID: 1,
		Title: newProduct.Title,
		Price: newProduct.Price,
	})

	if err != nil {
		http.Error(w, "Could not create Product", http.StatusBadRequest)
	}

	
	utils.SendProducts(w, createdProduct)

}
