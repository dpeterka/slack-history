package funfacts

import (
	"time"

	"github.com/dpeterka/history-slackbot/internal/blobby"
	"github.com/dpeterka/history-slackbot/internal/camping"
	"github.com/dpeterka/history-slackbot/internal/emo"
	"github.com/dpeterka/history-slackbot/internal/foodtakes"
	"github.com/dpeterka/history-slackbot/internal/gardening"
	"github.com/dpeterka/history-slackbot/internal/hottub"
	"github.com/dpeterka/history-slackbot/internal/jokes"
	"github.com/dpeterka/history-slackbot/internal/printing3d"
	"github.com/dpeterka/history-slackbot/internal/quotes"
	"github.com/dpeterka/history-slackbot/internal/rotation"
	"github.com/dpeterka/history-slackbot/internal/rss"
	"github.com/dpeterka/history-slackbot/internal/wikihow"
	"github.com/dpeterka/history-slackbot/internal/wikihowquizzes"
	"github.com/dpeterka/history-slackbot/internal/wikipedia"
)

// FunFact represents a fun fact of any type
type FunFact struct {
	Text          string
	Type          string // "emo", "blobby", "wikihow", "wikihow_quizzes", "quote", "holidays", "hottub", "gardening", "printing3d", "camping", "joke", "foodtakes", "people", "events"
	Category      string
	URL           string             // For wikihow articles
	Author        string             // For quotes
	Holidays      []rss.Holiday      // For holidays (with links)
	NotablePeople []wikipedia.Person // For people (births/deaths)
	ShowEvents    bool               // For events - signals to show historical events
}

// GetRandomFunFact returns a random fun fact from all available types
// For holidays, pass them via GetRandomFunFactWithData
func GetRandomFunFact(includeEmo, includeBlobby, includeWikiHow, includeWikiHowQuizzes, includeQuote, includeHotTub, includeGardening, includePrinting3D, includeCamping, includeJoke, includeFoodTakes, includePeople, includeEvents bool) *FunFact {
	return GetRandomFunFactWithData(includeEmo, includeBlobby, includeWikiHow, includeWikiHowQuizzes, includeQuote, includeHotTub, includeGardening, includePrinting3D, includeCamping, includeJoke, includeFoodTakes, includePeople, includeEvents, false, nil, nil, 0, 0)
}

// DetermineContentType returns what content type will be selected for today without fetching data
// This allows the caller to only fetch data for the selected type
func DetermineContentType(includeEmo, includeBlobby, includeWikiHow, includeWikiHowQuizzes, includeQuote, includeHotTub, includeGardening, includePrinting3D, includeCamping, includeJoke, includeFoodTakes, includePeople, includeEvents, includeHolidays bool, seed int) string {
	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	// Count how many types are enabled (without requiring data)
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
	if includeWikiHowQuizzes {
		enabledTypes = append(enabledTypes, "wikihow_quizzes")
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
	if includePrinting3D {
		enabledTypes = append(enabledTypes, "printing3d")
	}
	if includeCamping {
		enabledTypes = append(enabledTypes, "camping")
	}
	if includeJoke {
		enabledTypes = append(enabledTypes, "joke")
	}
	if includeFoodTakes {
		enabledTypes = append(enabledTypes, "foodtakes")
	}
	if includePeople {
		enabledTypes = append(enabledTypes, "people")
	}
	if includeEvents {
		enabledTypes = append(enabledTypes, "events")
	}
	if includeHolidays {
		enabledTypes = append(enabledTypes, "holidays")
	}

	// If none enabled, return empty
	if len(enabledTypes) == 0 {
		return ""
	}

	// Rotate between enabled types based on weekday count (the bot only
	// posts Mon-Fri; rotating on raw days would starve types whenever the
	// type count is a multiple of 7)
	return enabledTypes[rotation.WeekdaysSinceEpoch(seed)%len(enabledTypes)]
}

// GetRandomFunFactWithData returns a random fun fact with optional holidays and people data, and seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
// Rotates between enabled types based on day
func GetRandomFunFactWithData(includeEmo, includeBlobby, includeWikiHow, includeWikiHowQuizzes, includeQuote, includeHotTub, includeGardening, includePrinting3D, includeCamping, includeJoke, includeFoodTakes, includePeople, includeEvents, includeHolidays bool, holidays []rss.Holiday, people []wikipedia.Person, maxPeople, seed int) *FunFact {
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
	if includeWikiHowQuizzes {
		enabledTypes = append(enabledTypes, "wikihow_quizzes")
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
	if includePrinting3D {
		enabledTypes = append(enabledTypes, "printing3d")
	}
	if includeCamping {
		enabledTypes = append(enabledTypes, "camping")
	}
	if includeJoke {
		enabledTypes = append(enabledTypes, "joke")
	}
	if includeFoodTakes {
		enabledTypes = append(enabledTypes, "foodtakes")
	}
	if includePeople {
		enabledTypes = append(enabledTypes, "people")
	}
	if includeEvents {
		enabledTypes = append(enabledTypes, "events")
	}
	if includeHolidays {
		enabledTypes = append(enabledTypes, "holidays")
	}

	// If none enabled, return nil
	if len(enabledTypes) == 0 {
		return nil
	}

	// Rotate between enabled types based on weekday count (must match
	// DetermineContentType)
	selectedType := enabledTypes[rotation.WeekdaysSinceEpoch(seed)%len(enabledTypes)]

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
			Text: article.Title,
			Type: "wikihow",
			URL:  article.URL,
		}
	case "wikihow_quizzes":
		quiz := wikihowquizzes.GetRandomQuizWithSeed(seed)
		return &FunFact{
			Text: quiz.Title,
			Type: "wikihow_quizzes",
			URL:  quiz.URL,
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
	case "printing3d":
		printing3dTip := printing3d.GetRandomTipWithSeed(seed)
		return &FunFact{
			Text:     printing3dTip.Text,
			Type:     "printing3d",
			Category: printing3dTip.Category,
		}
	case "camping":
		campingTip := camping.GetRandomTipWithSeed(seed)
		return &FunFact{
			Text:     campingTip.Text,
			Type:     "camping",
			Category: campingTip.Category,
		}
	case "joke":
		joke := jokes.GetRandomJokeWithSeed(seed)
		return &FunFact{
			Text: joke.Text,
			Type: "joke",
		}
	case "foodtakes":
		take := foodtakes.GetRandomTakeWithSeed(seed)
		return &FunFact{
			Text:     take.Text,
			Type:     "foodtakes",
			Category: take.Category,
		}
	case "people":
		if len(people) == 0 {
			return nil
		}
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
		if len(holidays) == 0 {
			return nil
		}
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
	case "wikihow_quizzes":
		return "🧠 WikiHow Quiz of the Day"
	case "quote":
		return "💭 Quote of the Day"
	case "hottub":
		return "🛁 Hot Tub Care Tip"
	case "gardening":
		return "🌱 Gardening Tip"
	case "printing3d":
		return "🖨️ 3D Printing Tip"
	case "camping":
		return "🏕️ Camping Tip of the Day"
	case "foodtakes":
		return "🌭 Unhinged Food Take of the Day"
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

// aiRefreshablePacks maps text-based pack types to a description used when
// asking the generator for a fresh item in that pack's style.
var aiRefreshablePacks = map[string]string{
	"emo":        "a dry, darkly funny observation about work or life (think tired office philosopher)",
	"blobby":     "an absurd but plausible-sounding 'fact' about Mr Blobby, the chaotic pink British TV character",
	"joke":       "a groan-worthy dad joke or one-liner",
	"camping":    "a blunt, irreverent camping tip that is actually useful",
	"hottub":     "a practical hot tub care tip delivered with personality",
	"gardening":  "a practical vegetable or hydroponic gardening tip with a bit of wit",
	"printing3d": "a practical 3D printing tip with hard-earned-wisdom energy",
	"foodtakes":  "a deliberately controversial but harmless food opinion designed to start friendly arguments",
}

// IsAIRefreshable reports whether a pack type's content can be freshly
// generated fresh each day (only plain-text packs; link/data types can't).
func IsAIRefreshable(packType string) bool {
	_, ok := aiRefreshablePacks[packType]
	return ok
}

// PackDescription returns the generation prompt description for a refreshable pack.
func PackDescription(packType string) string {
	return aiRefreshablePacks[packType]
}

// StyleExamples returns a few items from the pack to calibrate the generator's
// tone and length. Nearby day seeds are used so examples differ from today's pick.
func StyleExamples(packType string, seed int) []string {
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	examples := []string{}
	for i := 1; i <= 3; i++ {
		s := seed + i
		switch packType {
		case "emo":
			examples = append(examples, emo.GetRandomCommentWithSeed(s).Text)
		case "blobby":
			examples = append(examples, blobby.GetRandomFactWithSeed(s).Text)
		case "joke":
			examples = append(examples, jokes.GetRandomJokeWithSeed(s).Text)
		case "camping":
			examples = append(examples, camping.GetRandomTipWithSeed(s).Text)
		case "hottub":
			examples = append(examples, hottub.GetRandomTipWithSeed(s).Text)
		case "gardening":
			examples = append(examples, gardening.GetRandomTipWithSeed(s).Text)
		case "printing3d":
			examples = append(examples, printing3d.GetRandomTipWithSeed(s).Text)
		case "foodtakes":
			examples = append(examples, foodtakes.GetRandomTakeWithSeed(s).Text)
		}
	}
	return examples
}
