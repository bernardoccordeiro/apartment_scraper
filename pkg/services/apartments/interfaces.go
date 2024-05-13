package apartments

import (
	"github.com/bernardoccordeiro/apartment_scraper/pkg/domain"
)

type ApartmentService interface {
	GetApartments() ([]domain.ApartmentListing, error)
}