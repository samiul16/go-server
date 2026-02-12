package utils

import (
	"encoding/json"
	"net/http"
)

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func SendPaginatedData(w http.ResponseWriter, data any, Page int, Limit int, TotalItems int, TotalPages int) {

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page: Page,
			Limit: Limit,
			TotalItems: TotalItems,
			TotalPages: TotalPages,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedData)
}
