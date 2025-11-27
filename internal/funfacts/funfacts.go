package funfacts

import (
	"time"

	"github.com/dpeterka/history-slackbot/internal/blobby"
	"github.com/dpeterka/history-slackbot/internal/emo"
	"github.com/dpeterka/history-slackbot/internal/gardening"
	"github.com/dpeterka/history-slackbot/internal/hottub"
	"github.com/dpeterka/history-slackbot/internal/quotes"
	"github.com/dpeterka/history-slackbot/internal/rss"
	"github.com/dpeterka/history-slackbot/internal/wikipedia"
	"github.com/dpeterka/history-slackbot/internal/wikihow"
)

// FunFact represents a fun fact of any type
type FunFact struct {
	Text           string
	Type           string              // "emo", "blobby", "wikihow", "quote", "holidays", "hottub", "gardening", "people", "events"
	Category       string
	URL            string              // For wikihow articles
	Author         string              // For quotes
	Holidays       []rss.Holiday       // For holidays (with links)
	NotablePeople  []wikipedia.Person  // For people (births/deaths)
	ShowEvents     bool                // For events - signals to show historical events
}

// GetRandomFunFact returns a random fun fact from all available types
// For holidays, pass them via GetRandomFunFactWithData
func GetRandomFunFact(includeEmo, includeBlobby, includeWikiHow, includeQuote, includeHotTub, includeGardening, includePeople, includeEvents bool) *FunFact {
	return GetRandomFunFactWithData(includeEmo, includeBlobby, includeWikiHow, includeQuote, includeHotTub, includeGardening, includePeople, includeEvents, nil, nil, 0, 0)
}

// GetRandomFunFactWithData returns a random fun fact with optional holidays and people data, and seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
// Rotates between enabled types based on day
func GetRandomFunFactWithData(includeEmo, includeBlobby, includeWikiHow, includeQuote, includeHotTub, includeGardening, includePeople, includeEvents bool, holidays []rss.Holiday, people []wikipedia.Person, maxPeople, seed int) *FunFact {
	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	// Count how many types are enabled
	enabledTypes := []string{}
	if includeEmo {
		enabledTypes = append(enabledTypes, "emo")
	}
	if includeBlobby {
		enabledTypes = append(enabledTypes, "blobby")
	}
	if includeWikiHow {
		enabledTypes = append(enabledTypes, "wikihow")
	}
	if includeQuote {
		enabledTypes = append(enabledTypes, "quote")
	}
	if includeHotTub {
		enabledTypes = append(enabledTypes, "hottub")
	}
	if includeGardening {
		enabledTypes = append(enabledTypes, "gardening")
	}
	if includePeople && len(people) > 0 {
		enabledTypes = append(enabledTypes, "people")
	}
	if includeEvents {
		enabledTypes = append(enabledTypes, "events")
	}
	if len(holidays) > 0 {
		enabledTypes = append(enabledTypes, "holidays")
	}

	// If none enabled, return nil
	if len(enabledTypes) == 0 {
		return nil
	}

	// Rotate between enabled types based on day
	dayOfYear := seed % 1000 // Get just the day portion (MMDD)
	selectedType := enabledTypes[dayOfYear%len(enabledTypes)]

	switch selectedType {
	case "emo":
		emoComment := emo.GetRandomCommentWithSeed(seed)
		return &FunFact{
			Text:     emoComment.Text,
			Type:     "emo",
			Category: emoComment.Category,
		}
	case "blobby":
		blobbyFact := blobby.GetRandomFactWithSeed(seed)
		return &FunFact{
			Text:     blobbyFact.Text,
			Type:     "blobby",
			Category: blobbyFact.Category,
		}
	case "wikihow":
		article := wikihow.GetRandomArticleWithSeed(seed)
		return &FunFact{
			Text:     article.Title,
			Type:     "wikihow",
			URL:      article.URL,
		}
	case "quote":
		fetcher := quotes.NewFetcher()
		quote, err := fetcher.FetchQuoteOfTheDay()
		if err != nil || quote == nil {
			// Fallback if quote fetch fails - try next type
			return nil
		}
		return &FunFact{
			Text:   quote.Text,
			Type:   "quote",
			Author: quote.Author,
		}
	case "hottub":
		hotTubTip := hottub.GetRandomTipWithSeed(seed)
		return &FunFact{
			Text:     hotTubTip.Text,
			Type:     "hottub",
			Category: hotTubTip.Category,
		}
	case "gardening":
		gardeningTip := gardening.GetRandomTipWithSeed(seed)
		return &FunFact{
			Text:     gardeningTip.Text,
			Type:     "gardening",
			Category: gardeningTip.Category,
		}
	case "people":
		return &FunFact{
			Type:          "people",
			NotablePeople: people,
		}
	case "events":
		return &FunFact{
			Type:       "events",
			ShowEvents: true,
		}
	case "holidays":
		return &FunFact{
			Type:     "holidays",
			Holidays: holidays,
		}
	}

	return nil
}

// GetDisplayTitle returns the appropriate display title for the fact type
func (f *FunFact) GetDisplayTitle() string {
	switch f.Type {
	case "emo":
		return "💭 Today's Thought"
	case "blobby":
		return "🎀 Mr Blobby Fact of the Day"
	case "wikihow":
		return "📚 Helpful WikiHow Article"
	case "quote":
		return "💭 Quote of the Day"
	case "hottub":
		return "🛁 Hot Tub Care Tip"
	case "gardening":
		return "🌱 Gardening Tip"
	case "people":
		return "🎂 Notable People"
	case "events":
		return "📜 Historical Event"
	case "holidays":
		return "🎉 Today's Fun Holidays"
	default:
		return "💡 Fun Fact"
	}
}

// ShouldDisplayAsItalic returns whether the fact should be displayed in italic
func (f *FunFact) ShouldDisplayAsItalic() bool {
	return f.Type == "emo"
}
