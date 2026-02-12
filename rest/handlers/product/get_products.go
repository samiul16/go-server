package product

import (
	"encoding/json"
	"fmt"
	"go-server/domain"
	"go-server/utils"
	"log"
	"net/http"
	"strconv"
)

// sendProducts serializes the list of products into a JSON response.
func sendProducts(w http.ResponseWriter, products any) {
	// Set the Content-Type header to inform the client that the response body is JSON
	w.Header().Set("Content-Type", "application/json")

	// Serialize the products slice to a JSON byte slice
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		// If serialization fails (shouldn't happen here), write an internal server error
		log.Printf("Error encoding products: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

type paginatedData struct {
	Data []*domain.Product `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// getProducts handles the GET /products request.
func  (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch/define the data (in a real app, this comes from a database)
	reqQuery := r.URL.Query()
	fmt.Println("reqQuery", reqQuery)
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")
	fmt.Println("pageAsString", pageAsString)
	page, err := strconv.Atoi(pageAsString)
	limit, err := strconv.Atoi(limitAsString)
	if err != nil {
		page = 1
	}

	if page == 0 || page < 1 {
		page = 1
	}

	if limit == 0 || limit < 1 {
		limit = 10
	}
	
	 products, error := h.svc.List(page, limit)

	 fmt.Println("products", products)

	 if error != nil {
		fmt.Println("Error occured on getting products")
	 }

	 cnt, err := h.svc.Count()

	 if err != nil {
		fmt.Println("Error occured on getting products count")
	}
	
	// 2. Send the data to the client using the helper function
	utils.SendPaginatedData(w, products, page, limit, cnt, (cnt + limit - 1) / limit)
}