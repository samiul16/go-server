package product

import (
	"fmt"
	"go-server/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)

	if err != nil {
		fmt.Println("wrong product id")
	}

	product, err := h.porductRepo.Get(pId)

	utils.SendProduct(w, *product)

}