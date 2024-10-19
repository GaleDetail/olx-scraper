package utils

import (
	"fmt"
	"time"

	"olx_scraper/models"
)

// GenerateStatistics calculates and prints statistics for sold items
func GenerateStatistics(items map[string]models.Item) {
	categoryStats := calculateCategoryStatistics(items)
	printStatistics(categoryStats)
}

// calculateCategoryStatistics calculates durations for each category
func calculateCategoryStatistics(items map[string]models.Item) map[string][]time.Duration {
	categoryStats := make(map[string][]time.Duration)
	for _, item := range items {
		if item.DateSold != nil {
			duration := item.DateSold.Sub(item.DatePosted)
			categoryStats[item.Category] = append(categoryStats[item.Category], duration)
		}
	}
	return categoryStats
}

// printStatistics prints the average sale duration for each category
func printStatistics(categoryStats map[string][]time.Duration) {
	for category, durations := range categoryStats {
		var totalDuration time.Duration
		for _, duration := range durations {
			totalDuration += duration
		}
		averageDuration := totalDuration / time.Duration(len(durations))
		fmt.Printf("Category: %s\n", category)
		fmt.Printf("Items Sold: %d\n", len(durations))
		fmt.Printf("Average Sale Duration: %.2f hours\n\n", averageDuration.Hours())
	}
}
