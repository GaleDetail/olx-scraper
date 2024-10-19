package main

import (
	"log"

	"olx_scraper/scraper"
	"olx_scraper/storage"
	"olx_scraper/utils"
)

func main() {
	dataFile := "items.csv"

	// Initialize storage
	storage := storage.NewFileStorage()

	// Read existing items from file
	existingItems, err := storage.ReadItems(dataFile)
	if err != nil {
		log.Println("Error reading items:", err)
		// Handle error as needed
	}
	// Initialize scraper
	scraper := scraper.NewScraper(existingItems)

	// Start scraping process
	scraper.Scrape()

	// Update sold items
	scraper.UpdateSoldItems()

	// Write updated items to file
	if err != nil {
		log.Println("Error reading items:", err)
		// Handle error as needed
	}

	// Generate statistics
	utils.GenerateStatistics(scraper.ExistingItems)

	log.Println("Scraping completed!")
}
