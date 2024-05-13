package main

import (
	"log"

	"github.com/bernardoccordeiro/apartment_scraper/pkg/webservers"
)

func main() {
	if err := webservers.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
