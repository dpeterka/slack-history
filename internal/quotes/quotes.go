package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// Quote represents a quote of the day
type Quote struct {
	Text   string
	Author string
	Source string
}

// Fetcher handles fetching quotes
type Fetcher struct {
	client *http.Client
}

// NewFetcher creates a new quote fetcher
func NewFetcher() *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// ZenQuotesResponse represents the response from ZenQuotes API
type ZenQuotesResponse struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
	HTML   string `json:"h"`
}

// FetchQuoteOfTheDay fetches a quote of the day
// Uses ZenQuotes API (free, no API key required)
func (f *Fetcher) FetchQuoteOfTheDay() (*Quote, error) {
	// Try ZenQuotes API first
	quote, err := f.fetchFromZenQuotes()
	if err == nil {
		return quote, nil
	}

	// Fallback to hardcoded quotes if API fails
	return f.getFallbackQuote(), nil
}

// fetchFromZenQuotes fetches a quote from ZenQuotes API
func (f *Fetcher) fetchFromZenQuotes() (*Quote, error) {
	resp, err := f.client.Get("https://zenquotes.io/api/today")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quote: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var quotes []ZenQuotesResponse
	if err := json.Unmarshal(body, &quotes); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quotes in response")
	}

	return &Quote{
		Text:   quotes[0].Quote,
		Author: quotes[0].Author,
		Source: "ZenQuotes",
	}, nil
}

// getFallbackQuote returns a hardcoded quote as fallback
func (f *Fetcher) getFallbackQuote() *Quote {
	fallbackQuotes := []Quote{
		{
			Text:   "The only way to do great work is to love what you do.",
			Author: "Steve Jobs",
			Source: "Fallback",
		},
		{
			Text:   "Innovation distinguishes between a leader and a follower.",
			Author: "Steve Jobs",
			Source: "Fallback",
		},
		{
			Text:   "The future belongs to those who believe in the beauty of their dreams.",
			Author: "Eleanor Roosevelt",
			Source: "Fallback",
		},
		{
			Text:   "It is during our darkest moments that we must focus to see the light.",
			Author: "Aristotle",
			Source: "Fallback",
		},
		{
			Text:   "The only impossible journey is the one you never begin.",
			Author: "Tony Robbins",
			Source: "Fallback",
		},
		{
			Text:   "Life is what happens when you're busy making other plans.",
			Author: "John Lennon",
			Source: "Fallback",
		},
		{
			Text:   "The way to get started is to quit talking and begin doing.",
			Author: "Walt Disney",
			Source: "Fallback",
		},
		{
			Text:   "Don't let yesterday take up too much of today.",
			Author: "Will Rogers",
			Source: "Fallback",
		},
		{
			Text:   "You learn more from failure than from success. Don't let it stop you. Failure builds character.",
			Author: "Unknown",
			Source: "Fallback",
		},
		{
			Text:   "If you are working on something that you really care about, you don't have to be pushed. The vision pulls you.",
			Author: "Steve Jobs",
			Source: "Fallback",
		},
	}

	// Seed random with current date to get consistent quote for the day
	now := time.Now()
	seed := now.Year()*10000 + int(now.Month())*100 + now.Day()
	r := rand.New(rand.NewSource(int64(seed)))

	return &fallbackQuotes[r.Intn(len(fallbackQuotes))]
}
