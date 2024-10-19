package scraper

import (
	"log"
	"time"

	"github.com/gocolly/colly"
	"olx_scraper/models"
)

// Scraper struct holds the state of the scraper
type Scraper struct {
	Collector     *colly.Collector
	ItemCollector *colly.Collector
	Categories    []string
	ExistingItems map[string]models.Item
	FoundOlxIDs   map[string]bool
}

// NewScraper creates a new Scraper instance
func NewScraper(existingItems map[string]models.Item) *Scraper {
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ua", "olx.ua"),
		colly.UserAgent("Mozilla/5.0 ..."),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*olx.ua*",
		RandomDelay: 2 * time.Second,
	})

	s := &Scraper{
		Collector:     c,
		ItemCollector: c.Clone(),
		ExistingItems: existingItems,
		FoundOlxIDs:   make(map[string]bool),
		Categories:    []string{},
	}

	// Set up callbacks
	s.setupCallbacks()

	return s
}

// setupCallbacks sets up all the callbacks
func (s *Scraper) setupCallbacks() {
	s.setupErrorCallbacks()
	s.setupCategoryCallbacks()
	s.setupItemCallbacks()
	s.setupPaginationCallbacks()
}

// Scrape starts the scraping process
func (s *Scraper) Scrape() {
	s.visitMainPage()
	s.scrapeCategories()
}

// visitMainPage visits the main page and waits for category collection
func (s *Scraper) visitMainPage() {
	err := s.Collector.Visit("https://www.olx.ua/")
	if err != nil {
		log.Fatal("Failed to visit main page:", err)
	}
	s.Collector.Wait()
	log.Println("Found categories:", len(s.Categories))
}

// scrapeCategories visits each category page
func (s *Scraper) scrapeCategories() {
	for _, categoryURL := range s.Categories {
		log.Println("Scraping category:", categoryURL)
		err := s.ItemCollector.Visit(categoryURL)
		if err != nil {
			log.Println("Error visiting category:", err)
			continue
		}
	}
	s.ItemCollector.Wait()
}

// UpdateSoldItems updates the status of items that are no longer found
func (s *Scraper) UpdateSoldItems() {
	for olxID, item := range s.ExistingItems {
		if !s.FoundOlxIDs[olxID] && item.DateSold == nil {
			now := time.Now()
			item.DateSold = &now
			s.ExistingItems[olxID] = item
		}
	}
}
