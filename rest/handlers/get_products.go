package handlers

import (
	"encoding/json"
	"go-server/database" // Adjust path based on your go.mod module name
	"log"
	"net/http"
)

// sendProducts serializes the list of products into a JSON response.
func sendProducts(w http.ResponseWriter, products []database.Product) {
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
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch/define the data (in a real app, this comes from a database)
	var products = []database.Product{
		{ID: 1, Name: "Product1", Price: 45.66},
		{ID: 2, Name: "Product2", Price: 75.99}, // Changed price for variety
		{ID: 3, Name: "Product3", Price: 12.50}, // Changed price for variety
	}

	// 2. Send the data to the client using the helper function
	sendProducts(w, products)
}