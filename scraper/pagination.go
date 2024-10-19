package scraper

import (
	"log"

	"github.com/gocolly/colly"
)

// setupPaginationCallbacks sets up callbacks for pagination
func (s *Scraper) setupPaginationCallbacks() {
	s.ItemCollector.OnHTML("a[data-cy='pagination-forward']", s.onPagination)
}

// onPagination handles pagination by visiting the next page
func (s *Scraper) onPagination(e *colly.HTMLElement) {
	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
	log.Println("Moving to next page:", nextPage)
	e.Request.Visit(nextPage)
}
