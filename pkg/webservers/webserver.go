package webservers

import (
	"log"
	"net/http"
	"os"

	"github.com/bernardoccordeiro/apartment_scraper/pkg/config"
	"github.com/bernardoccordeiro/apartment_scraper/pkg/handlers"
	"github.com/bernardoccordeiro/apartment_scraper/pkg/services/apartments"
	"github.com/bernardoccordeiro/apartment_scraper/pkg/services/scrapers/providers/mock_scraper"

	"github.com/gorilla/mux"
)

func Run() error {
	cfg := config.LoadConfig()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Define scraper providers to use
	mock_scraper := mock_scraper.New()

	providers := []apartments.ScraperProvider{{
		Name: "mock_scraper", Scraper: mock_scraper,
	}}
	service := apartments.NewApartmentService(providers)
	h := handlers.New(logger, service)

	r := mux.NewRouter()

	r.HandleFunc("/apartments", h.GetApartments).Methods("GET")
	r.Use(loggingMiddleware(logger))
	r.Use(recoverMiddleware(logger))
    r.Use(corsMiddleware())

	logger.Printf("Starting server on port: %s", cfg.ServerAddress)
	
	server := http.Server{
		Addr: cfg.ServerAddress,
		Handler: r,
	}
	return server.ListenAndServe()
}
