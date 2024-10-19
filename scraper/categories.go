package scraper

import (
	"github.com/gocolly/colly"
)

// setupCategoryCallbacks sets up callbacks for category collection
func (s *Scraper) setupCategoryCallbacks() {
	s.Collector.OnHTML("div[data-testid='home-categories-menu-row'] a", s.onCategoryLink)
}

// onCategoryLink collects category URLs from the main page
func (s *Scraper) onCategoryLink(e *colly.HTMLElement) {
	categoryURL := e.Request.AbsoluteURL(e.Attr("href"))
	s.Categories = append(s.Categories, categoryURL)
}
