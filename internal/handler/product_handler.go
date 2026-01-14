package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/baihakhi/product-search/internal/model"
	"github.com/baihakhi/product-search/internal/response"
	"github.com/baihakhi/product-search/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q == "" {
		log.Println("Empty query parameter 'q'")
		response.WriteJSONError(w, http.StatusBadRequest, "query parameter 'q' is required")
		return
	}
	sizeParam := r.URL.Query().Get("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil || size <= 0 {
		size = 20
	}

	products, err := h.service.SearchProducts(ctx, size, q, "products")
	if err != nil {
		log.Printf("Error searching products: %v", err)
		response.WriteJSONError(w, http.StatusInternalServerError, "failed to search products")
		return
	}
	response := model.ResponseProduct{
		Results: products,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
