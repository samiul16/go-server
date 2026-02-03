package product

import (
	"encoding/json"
	"fmt"
	"go-server/repo"
	"log"
	"net/http"
)

// sendProducts serializes the list of products into a JSON response.
func sendProducts(w http.ResponseWriter, products []*repo.Product) {
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

// getProducts handles the GET /products request.
func  (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch/define the data (in a real app, this comes from a database)
	 products, error := h.porductRepo.List()

	 if error != nil {
		fmt.Println("Error occured on getting products")
	 }

	// 2. Send the data to the client using the helper function
	sendProducts(w, products)
}