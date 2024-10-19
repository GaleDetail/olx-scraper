package scraper

import (
	"log"

	"github.com/gocolly/colly"
)

// setupErrorCallbacks sets up error handling callbacks
func (s *Scraper) setupErrorCallbacks() {
	s.Collector.OnError(s.onError)
	s.ItemCollector.OnError(s.onError)
}

// onError handles errors during requests
func (s *Scraper) onError(r *colly.Response, err error) {
	log.Println("Error requesting", r.Request.URL, "-", err)
}
