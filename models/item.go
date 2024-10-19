package models

import "time"

// Item represents a product item scraped from OLX.ua
type Item struct {
	OlxID      string
	Category   string
	Title      string
	Price      int
	Location   string
	DatePosted time.Time
	DateSold   *time.Time // nil if the item is not sold yet
}
