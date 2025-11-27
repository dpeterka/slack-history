package holidays

import (
	"time"
)

// MajorHoliday represents a major holiday in a country
type MajorHoliday struct {
	Name    string
	Month   time.Month
	Day     int
	Country string
}

// GetMajorHolidays returns a list of major holidays for India, Canada, and the US
func GetMajorHolidays() []MajorHoliday {
	return []MajorHoliday{
		// United States
		{Name: "New Year's Day", Month: time.January, Day: 1, Country: "US"},
		{Name: "Independence Day", Month: time.July, Day: 4, Country: "US"},
		{Name: "Christmas Day", Month: time.December, Day: 25, Country: "US"},

		// Canada
		{Name: "New Year's Day", Month: time.January, Day: 1, Country: "Canada"},
		{Name: "Canada Day", Month: time.July, Day: 1, Country: "Canada"},
		{Name: "Christmas Day", Month: time.December, Day: 25, Country: "Canada"},

		// India
		{Name: "Republic Day", Month: time.January, Day: 26, Country: "India"},
		{Name: "Independence Day", Month: time.August, Day: 15, Country: "India"},
		{Name: "Gandhi Jayanti", Month: time.October, Day: 2, Country: "India"},
	}
}

// GetTodaysMajorHoliday checks if today is a major holiday in India, Canada, or the US
// Returns the holiday name and true if today is a major holiday, empty string and false otherwise
func GetTodaysMajorHoliday() (string, bool) {
	now := time.Now()
	currentMonth := now.Month()
	currentDay := now.Day()

	holidays := GetMajorHolidays()

	// Collect all matching holidays for today
	var matchingHolidays []string
	holidayMap := make(map[string]bool) // To deduplicate holidays

	for _, holiday := range holidays {
		if holiday.Month == currentMonth && holiday.Day == currentDay {
			// Avoid duplicates (e.g., New Year's Day and Christmas are in multiple countries)
			if !holidayMap[holiday.Name] {
				matchingHolidays = append(matchingHolidays, holiday.Name)
				holidayMap[holiday.Name] = true
			}
		}
	}

	// If we found any holidays, return the first one
	// (Since holidays like New Year's and Christmas are shared, we only need to return one)
	if len(matchingHolidays) > 0 {
		return matchingHolidays[0], true
	}

	return "", false
}
