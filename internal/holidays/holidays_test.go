package holidays

import (
	"testing"
	"time"
)

func TestGetMajorHolidays(t *testing.T) {
	holidays := GetMajorHolidays()

	// Verify we have at least the expected holidays
	if len(holidays) < 9 {
		t.Errorf("Expected at least 9 major holidays, got %d", len(holidays))
	}

	// Verify specific holidays exist
	expectedHolidays := map[string]bool{
		"New Year's Day":   false,
		"Independence Day": false, // Both US and India
		"Christmas Day":    false,
		"Canada Day":       false,
		"Republic Day":     false,
		"Gandhi Jayanti":   false,
	}

	for _, holiday := range holidays {
		if _, exists := expectedHolidays[holiday.Name]; exists {
			expectedHolidays[holiday.Name] = true
		}
	}

	for name, found := range expectedHolidays {
		if !found {
			t.Errorf("Expected holiday %s not found", name)
		}
	}
}

func TestGetTodaysMajorHoliday(t *testing.T) {
	// This test will vary based on the current date
	// We can test the logic by checking specific dates

	tests := []struct {
		name          string
		month         time.Month
		day           int
		shouldBeEmpty bool
	}{
		{
			name:          "New Year's Day",
			month:         time.January,
			day:           1,
			shouldBeEmpty: false,
		},
		{
			name:          "Random day with no holiday",
			month:         time.March,
			day:           15,
			shouldBeEmpty: true,
		},
		{
			name:          "Christmas Day",
			month:         time.December,
			day:           25,
			shouldBeEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This test just verifies the structure
			// The actual GetTodaysMajorHoliday() function uses time.Now()
			// so we can't directly test it without mocking
			// But we can verify the logic works for the data structure

			holidays := GetMajorHolidays()
			found := false
			for _, holiday := range holidays {
				if holiday.Month == tt.month && holiday.Day == tt.day {
					found = true
					break
				}
			}

			if found == tt.shouldBeEmpty {
				if tt.shouldBeEmpty {
					t.Errorf("Expected no holiday on %v %d, but found one", tt.month, tt.day)
				} else {
					t.Errorf("Expected a holiday on %v %d, but none found", tt.month, tt.day)
				}
			}
		})
	}
}
