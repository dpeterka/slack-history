package quotes

import (
	"testing"
)

func TestGetFallbackQuote(t *testing.T) {
	fetcher := NewFetcher()
	quote := fetcher.getFallbackQuote()

	if quote == nil {
		t.Fatal("Expected quote, got nil")
	}

	if quote.Text == "" {
		t.Error("Expected non-empty quote text")
	}

	if quote.Author == "" {
		t.Error("Expected non-empty author")
	}

	if quote.Source != "Fallback" {
		t.Errorf("Expected source 'Fallback', got '%s'", quote.Source)
	}
}

func TestFallbackQuoteConsistency(t *testing.T) {
	fetcher := NewFetcher()

	// Call multiple times, should get same quote for the same day
	quote1 := fetcher.getFallbackQuote()
	quote2 := fetcher.getFallbackQuote()

	if quote1.Text != quote2.Text {
		t.Error("Expected same quote for the same day")
	}

	if quote1.Author != quote2.Author {
		t.Error("Expected same author for the same day")
	}
}
