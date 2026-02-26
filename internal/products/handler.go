// Package products contains all about products
package products

import (
	"log"
	"net/http"

	repo "github.com/EmiliodDev/go-pos/internal/adapters/pgdb/sqlc"
	"github.com/EmiliodDev/go-pos/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.listProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Write(w, http.StatusOK, resp); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productParams repo.CreateProductParams
	if err := json.Read(r, &productParams); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := h.service.createProduct(r.Context(), productParams)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Write(w, http.StatusCreated, resp); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
