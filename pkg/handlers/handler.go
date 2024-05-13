package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bernardoccordeiro/apartment_scraper/pkg/services/apartments"
)

type Handler struct {
	Service apartments.ApartmentService
	Logger *log.Logger
}

func New(logger *log.Logger, apartmentService apartments.ApartmentService) *Handler {
	return &Handler{Service: apartmentService, Logger: logger}
}

func (h *Handler) GetApartments(w http.ResponseWriter, r *http.Request) {
	h.Logger.Print("Entered fnction")
	listings, err := h.Service.GetApartments()
	if err != nil {
		http.Error(w, "Failed to get apartment listings", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(listings)
}
