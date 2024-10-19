package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ExtractDigits extracts digits from a string
func ExtractDigits(input string) string {
	re := regexp.MustCompile(`[^\d]`)
	return re.ReplaceAllString(input, "")
}

// StringToInt converts a string to an integer
func StringToInt(input string) (int, error) {
	return strconv.Atoi(input)
}

// SplitLocationDate splits location and date from the combined text
func SplitLocationDate(locationDate string) (string, string) {
	location := ""
	dateStr := ""
	if locationDate != "" {
		parts := strings.Split(locationDate, " - ")
		if len(parts) >= 1 {
			location = parts[0]
		}
		if len(parts) >= 2 {
			dateStr = parts[1]
		}
	}
	return location, dateStr
}

// ParseDate parses a date string from OLX.ua and returns a time.Time
func ParseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)

	if strings.HasPrefix(dateStr, "Сьогодні") {
		timeStr := strings.TrimPrefix(dateStr, "Сьогодні")
		return parseTodayDate(strings.TrimSpace(timeStr))
	} else if strings.HasPrefix(dateStr, "Вчора") {
		timeStr := strings.TrimPrefix(dateStr, "Вчора")
		return parseYesterdayDate(strings.TrimSpace(timeStr))
	} else {
		return parseSpecificDate(dateStr)
	}
}

// parseTodayDate parses time for "Сьогодні" date string
func parseTodayDate(timeStr string) (time.Time, error) {
	now := time.Now()
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return now, err
	}
	return time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location()), nil
}

// parseYesterdayDate parses time for "Вчора" date string
func parseYesterdayDate(timeStr string) (time.Time, error) {
	now := time.Now()
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return now, err
	}
	yesterday := now.AddDate(0, 0, -1)
	return time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location()), nil
}

// parseSpecificDate parses date strings like "15 жовтня"
func parseSpecificDate(dateStr string) (time.Time, error) {
	now := time.Now()
	months := map[string]time.Month{
		"січня":     time.January,
		"лютого":    time.February,
		"березня":   time.March,
		"квітня":    time.April,
		"травня":    time.May,
		"червня":    time.June,
		"липня":     time.July,
		"серпня":    time.August,
		"вересня":   time.September,
		"жовтня":    time.October,
		"листопада": time.November,
		"грудня":    time.December,
	}

	parts := strings.Split(dateStr, " ")
	if len(parts) != 2 {
		return now, fmt.Errorf("unknown date format")
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return now, err
	}

	monthName := strings.ToLower(parts[1])
	month, ok := months[monthName]
	if !ok {
		return now, fmt.Errorf("unknown month")
	}

	return time.Date(now.Year(), month, day, 0, 0, 0, 0, now.Location()), nil
}

// GetCategoryFromURL extracts the category from a given URL
func GetCategoryFromURL(url string) string {
	parts := strings.Split(url, "/")
	for i, part := range parts {
		if part == "uk" || part == "d" || part == "" {
			continue
		}
		if i > 2 {
			return part
		}
	}
	return "Unknown Category"
}
