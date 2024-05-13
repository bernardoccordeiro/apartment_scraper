package apartments

import (
	"fmt"

	"github.com/bernardoccordeiro/apartment_scraper/pkg/domain"

	"github.com/bernardoccordeiro/apartment_scraper/pkg/services/scrapers"
)

type ScraperProvider struct {
	Name string
	Scraper scrapers.Scraper
}

type ApartmentServiceImpl struct {
	Scrapers []ScraperProvider
}

func NewApartmentService(providers []ScraperProvider) *ApartmentServiceImpl {
	return &ApartmentServiceImpl{Scrapers: providers}
}

func (s *ApartmentServiceImpl) GetApartments() ([]domain.ApartmentListing, error) {
	listings := make([]domain.ApartmentListing, 0)
	for _, scrp := range s.Scrapers {
		scrapedListings, err := scrp.Scraper.Scrape()
		if err != nil {
			return nil, fmt.Errorf("error scraping: %w", err)
		}
		listings = append(listings, scrapedListings...)
	}
	return listings, nil
}
