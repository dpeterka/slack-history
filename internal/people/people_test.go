package people

import (
	"testing"

	"github.com/dpeterka/history-slackbot/internal/rss"
)

func TestExtractPeople(t *testing.T) {
	events := []rss.HistoricalEvent{
		{
			Title:       "Albert Einstein Born",
			Year:        "1879",
			Description: "Physicist Albert Einstein was born in Germany",
		},
		{
			Title:       "Death of Leonardo da Vinci",
			Year:        "1519",
			Description: "Renaissance artist Leonardo da Vinci died in France",
		},
		{
			Title:       "World War II Begins",
			Year:        "1939",
			Description: "Germany invades Poland",
		},
	}

	people := ExtractPeople(events)

	if len(people) != 2 {
		t.Errorf("Expected 2 people, got %d", len(people))
	}

	// Check for birth
	foundBirth := false
	for _, p := range people {
		if p.Type == "birth" && p.Year == "1879" {
			foundBirth = true
		}
	}
	if !foundBirth {
		t.Error("Expected to find Einstein's birth")
	}

	// Check for death
	foundDeath := false
	for _, p := range people {
		if p.Type == "death" && p.Year == "1519" {
			foundDeath = true
		}
	}
	if !foundDeath {
		t.Error("Expected to find Leonardo's death")
	}
}

func TestFilterNotablePeople(t *testing.T) {
	people := []Person{
		{
			Name:        "Albert Einstein Born",
			Year:        "1879",
			Type:        "birth",
			Description: "Physicist Albert Einstein was born",
		},
		{
			Name:        "John Doe Born",
			Year:        "1900",
			Type:        "birth",
			Description: "Random person was born",
		},
	}

	notable := FilterNotablePeople(people)

	if len(notable) != 1 {
		t.Errorf("Expected 1 notable person, got %d", len(notable))
	}

	if notable[0].Name != "Albert Einstein Born" {
		t.Errorf("Expected Albert Einstein, got %s", notable[0].Name)
	}
}

func TestSeparateBirthsAndDeaths(t *testing.T) {
	people := []Person{
		{Type: "birth", Name: "Person 1"},
		{Type: "death", Name: "Person 2"},
		{Type: "birth", Name: "Person 3"},
	}

	births, deaths := SeparateBirthsAndDeaths(people)

	if len(births) != 2 {
		t.Errorf("Expected 2 births, got %d", len(births))
	}

	if len(deaths) != 1 {
		t.Errorf("Expected 1 death, got %d", len(deaths))
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		substrs  []string
		expected bool
	}{
		{"Match found", "hello world", []string{"world", "foo"}, true},
		{"No match", "hello world", []string{"foo", "bar"}, false},
		{"Empty substrs", "hello world", []string{}, false},
		{"Case insensitive", "Hello World", []string{"world"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := containsAny(tc.str, tc.substrs)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
