package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"olx_scraper/models"
)

// WriteItems writes items to a CSV file
func (fs *FileStorage) WriteItems(filename string, items map[string]models.Item) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = fs.writeCSVRecords(file, items)
	if err != nil {
		return err
	}
	return nil
}

// writeCSVRecords writes items to the CSV file
func (fs *FileStorage) writeCSVRecords(file *os.File, items map[string]models.Item) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range items {
		record := fs.itemToRecord(item)
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	return nil
}

// itemToRecord converts an Item into a CSV record
func (fs *FileStorage) itemToRecord(item models.Item) []string {
	record := []string{
		item.OlxID,
		item.Category,
		item.Title,
		strconv.Itoa(item.Price),
		item.Location,
		item.DatePosted.Format(time.RFC3339),
		"",
	}
	if item.DateSold != nil {
		record[6] = item.DateSold.Format(time.RFC3339)
	}
	return record
}
