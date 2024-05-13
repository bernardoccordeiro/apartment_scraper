package scrapers

import (
	"github.com/bernardoccordeiro/apartment_scraper/pkg/domain"
)

type Scraper interface {
	Scrape() ([]domain.ApartmentListing, error)
}