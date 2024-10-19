package storage

import (
	"olx_scraper/models"
)

// Storage interface defines methods for reading and writing items
type Storage interface {
	ReadItems(filename string) (map[string]models.Item, error)
	WriteItems(filename string, items map[string]models.Item) error
}

// FileStorage implements the Storage interface using CSV files
type FileStorage struct{}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage() *FileStorage {
	return &FileStorage{}
}
