package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"olx_scraper/models"
)

// ReadItems reads items from a CSV file and returns a map of items
func (fs *FileStorage) ReadItems(filename string) (map[string]models.Item, error) {
	items := make(map[string]models.Item)

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist; return empty items map
			return items, nil
		}
		return nil, err
	}
	defer file.Close()

	records, err := fs.readCSVRecords(file)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if len(record) < 7 {
			continue
		}
		item, err := fs.parseRecord(record)
		if err != nil {
			// Log the error and skip the record
			// You can use a logger or handle it as needed
			continue
		}
		items[item.OlxID] = item
	}

	return items, nil
}

// readCSVRecords reads all CSV records from the file
func (fs *FileStorage) readCSVRecords(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// parseRecord parses a single CSV record into an Item
func (fs *FileStorage) parseRecord(record []string) (models.Item, error) {
	price, err := strconv.Atoi(record[3])
	if err != nil {
		price = 0
	}

	datePosted, err := time.Parse(time.RFC3339, record[5])
	if err != nil {
		datePosted = time.Now()
	}

	var dateSold *time.Time
	if record[6] != "" {
		ds, err := time.Parse(time.RFC3339, record[6])
		if err == nil {
			dateSold = &ds
		}
	}

	item := models.Item{
		OlxID:      record[0],
		Category:   record[1],
		Title:      record[2],
		Price:      price,
		Location:   record[4],
		DatePosted: datePosted,
		DateSold:   dateSold,
	}
	return item, nil
}
