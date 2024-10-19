package scraper

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	"olx_scraper/models"
	"olx_scraper/utils"
)

// Rest of the code remains the same...

// SetupItemCallbacks sets up callbacks for item processing
func (s *Scraper) setupItemCallbacks() {
	s.ItemCollector.OnHTML("div[data-cy='l-card']", func(e *colly.HTMLElement) {
		onItem(e, s)
	})
}

// onItem processes each item found on category pages
func onItem(e *colly.HTMLElement, s *Scraper) {
	item, err := extractItemData(e)
	if err != nil {
		log.Println("Error extracting item data:", err)
		return
	}
	s.FoundOlxIDs[item.OlxID] = true
	updateExistingItems(item, s)
}

// extractItemData extracts item data from the HTML element
func extractItemData(e *colly.HTMLElement) (models.Item, error) {
	olxID := e.Attr("id")
	if olxID == "" {
		return models.Item{}, fmt.Errorf("failed to get olxID")
	}
	log.Println("Processing item with ID:", olxID)

	title := e.ChildText("h6")
	priceText := e.ChildText("p[data-testid='ad-price']")
	locationDate := e.ChildText("p[data-testid='location-date']")

	price := parsePrice(priceText)
	location, dateStr := parseLocationDate(locationDate)
	datePosted := parseDatePosted(dateStr)
	category := utils.GetCategoryFromURL(e.Request.URL.String())

	item := models.Item{
		OlxID:      olxID,
		Category:   category,
		Title:      title,
		Price:      price,
		Location:   location,
		DatePosted: datePosted,
		DateSold:   nil,
	}
	return item, nil
}

// parsePrice cleans and converts price text to integer
func parsePrice(priceText string) int {
	priceDigits := utils.ExtractDigits(priceText)
	price, err := utils.StringToInt(priceDigits)
	if err != nil {
		price = 0
	}
	return price
}

// parseLocationDate splits location and date from the combined text
func parseLocationDate(locationDate string) (string, string) {
	return utils.SplitLocationDate(locationDate)
}

// parseDatePosted parses the date string into a time.Time
func parseDatePosted(dateStr string) time.Time {
	datePosted := time.Now()
	if dateStr != "" {
		dateParsed, err := utils.ParseDate(dateStr)
		if err == nil {
			datePosted = dateParsed
		}
	}
	return datePosted
}

// updateExistingItems updates or adds the item in the existing items map
func updateExistingItems(item models.Item, s *Scraper) {
	existingItem, exists := s.ExistingItems[item.OlxID]
	if exists {
		existingItem.Title = item.Title
		existingItem.Price = item.Price
		existingItem.Location = item.Location
		existingItem.Category = item.Category
		if existingItem.DateSold != nil {
			existingItem.DateSold = nil
		}
		s.ExistingItems[item.OlxID] = existingItem
	} else {
		s.ExistingItems[item.OlxID] = item
	}
}
