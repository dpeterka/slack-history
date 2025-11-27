package people

import (
	"strings"

	"github.com/dpeterka/history-slackbot/internal/rss"
)

// Person represents a notable person born or died on this day
type Person struct {
	Name        string
	Year        string
	Type        string // "birth" or "death"
	Description string
}

// ExtractPeople extracts notable births and deaths from historical events
func ExtractPeople(events []rss.HistoricalEvent) []Person {
	var people []Person

	for _, event := range events {
		title := strings.ToLower(event.Title)
		description := strings.ToLower(event.Description)
		combined := title + " " + description

		// Check for birth keywords
		if containsAny(combined, []string{"born", "birth of", "birthday"}) &&
			!containsAny(combined, []string{"died", "death", "killed", "assassination"}) {
			person := Person{
				Name:        event.Title,
				Year:        event.Year,
				Type:        "birth",
				Description: event.Description,
			}
			people = append(people, person)
		}

		// Check for death keywords
		if containsAny(combined, []string{"died", "death of", "dies", "passed away", "killed", "assassination"}) &&
			!containsAny(combined, []string{"born", "birth"}) {
			person := Person{
				Name:        event.Title,
				Year:        event.Year,
				Type:        "death",
				Description: event.Description,
			}
			people = append(people, person)
		}
	}

	return people
}

// FilterNotablePeople filters people by notability keywords
func FilterNotablePeople(people []Person) []Person {
	var notable []Person

	notabilityKeywords := []string{
		"president", "king", "queen", "emperor", "prime minister",
		"scientist", "inventor", "artist", "composer", "writer", "author",
		"actor", "actress", "director", "musician", "singer",
		"philosopher", "poet", "painter", "sculptor",
		"general", "admiral", "commander",
		"physicist", "chemist", "biologist", "mathematician",
		"nobel", "award", "prize",
		"founder", "pioneer", "revolutionary",
	}

	for _, person := range people {
		combined := strings.ToLower(person.Name + " " + person.Description)
		if containsAny(combined, notabilityKeywords) {
			notable = append(notable, person)
		}
	}

	return notable
}

// containsAny checks if a string contains any of the substrings (case-insensitive)
func containsAny(s string, substrs []string) bool {
	s = strings.ToLower(s)
	for _, substr := range substrs {
		if strings.Contains(s, strings.ToLower(substr)) {
			return true
		}
	}
	return false
}

// SeparateBirthsAndDeaths separates people into births and deaths
func SeparateBirthsAndDeaths(people []Person) (births []Person, deaths []Person) {
	for _, person := range people {
		if person.Type == "birth" {
			births = append(births, person)
		} else if person.Type == "death" {
			deaths = append(deaths, person)
		}
	}
	return births, deaths
}
