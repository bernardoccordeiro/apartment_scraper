package mock_scraper

import (
	"fmt"

	"github.com/bernardoccordeiro/apartment_scraper/pkg/domain"
)

type MockScraper struct{}

func New() *MockScraper {
	return &MockScraper{}
}

func (s *MockScraper) Scrape() ([]domain.ApartmentListing, error) {
	fmt.Println("Mock Scraping")
	return []domain.ApartmentListing{
		{Title: "Cheap Apartment", Description: "A nice place to stay", Price: 1200.00, URL: "http://www.google.com"},
	}, nil
}
