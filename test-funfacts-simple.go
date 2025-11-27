package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dpeterka/history-slackbot/internal/config"
	"github.com/dpeterka/history-slackbot/internal/wikipedia"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Test Wikipedia links for people
	fmt.Println("\n=== Testing People Links ===")
	people, err := wikipedia.FetchBirthsAndDeaths(12, 4, cfg.MaxPeople)
	if err != nil {
		log.Fatalf("Failed to fetch people: %v", err)
	}

	for _, person := range people {
		displayName := person.Name
		if commaIndex := len(person.Name); commaIndex > 0 {
			for i, c := range person.Name {
				if c == ',' {
					displayName = person.Name[:i]
					break
				}
			}
		}

		if person.WikiURL != "" {
			fmt.Printf("<%s|*%s*> (%d) - %s\n", person.WikiURL, displayName, person.Year, person.Type)
		} else {
			fmt.Printf("*%s* (%d) - %s (NO LINK)\n", person.Name, person.Year, person.Type)
		}
	}

	// Test Wikipedia search for event
	fmt.Println("\n=== Testing Event Link ===")
	eventTitle := "Tsar Alexander I"
	url, err := wikipedia.SearchWikipediaURL(eventTitle)
	if err != nil {
		fmt.Printf("Failed to find Wikipedia URL for '%s': %v\n", eventTitle, err)
	} else {
		fmt.Printf("Event: <%s|*%s*>\n", url, eventTitle)
	}

	time.Sleep(1 * time.Second)
}
