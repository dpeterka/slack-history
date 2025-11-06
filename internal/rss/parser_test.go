package rss

import (
	"testing"
)

func TestCleanHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Remove basic HTML tags",
			input:    "<p>Hello <strong>World</strong></p>",
			expected: "Hello World",
		},
		{
			name:     "Remove br tags and convert to newlines",
			input:    "Line 1<br>Line 2<br/>Line 3",
			expected: "Line 1\nLine 2\nLine 3",
		},
		{
			name:     "Remove multiple tags",
			input:    "<div><p>Test</p><span>Content</span></div>",
			expected: "Test\nContent",
		},
		{
			name:     "Handle empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Handle plain text",
			input:    "Plain text without tags",
			expected: "Plain text without tags",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanHTML(tt.input)
			if result != tt.expected {
				t.Errorf("cleanHTML() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestParseItem(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name     string
		item     Item
		expected HistoricalEvent
	}{
		{
			name: "Parse item with year in title",
			item: Item{
				Title:       "1969: Apollo 11 lands on the Moon",
				Description: "<p>The first human landing on the Moon</p>",
				Link:        "https://example.com/event/1",
				Categories:  []string{"Science", "Space"},
			},
			expected: HistoricalEvent{
				Year:        "1969",
				Title:       "Apollo 11 lands on the Moon",
				Description: "The first human landing on the Moon",
				Category:    "Science",
				Link:        "https://example.com/event/1",
			},
		},
		{
			name: "Parse item without year",
			item: Item{
				Title:       "Apollo 11 lands on the Moon",
				Description: "The first human landing on the Moon",
				Link:        "https://example.com/event/2",
				Categories:  []string{"Science"},
			},
			expected: HistoricalEvent{
				Year:        "",
				Title:       "Apollo 11 lands on the Moon",
				Description: "The first human landing on the Moon",
				Category:    "Science",
				Link:        "https://example.com/event/2",
			},
		},
		{
			name: "Parse item without category",
			item: Item{
				Title:       "1776: Declaration of Independence",
				Description: "The United States declares independence",
				Link:        "https://example.com/event/3",
				Categories:  []string{},
			},
			expected: HistoricalEvent{
				Year:        "1776",
				Title:       "Declaration of Independence",
				Description: "The United States declares independence",
				Category:    "",
				Link:        "https://example.com/event/3",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.parseItem(tt.item)

			if result.Year != tt.expected.Year {
				t.Errorf("Year = %q, want %q", result.Year, tt.expected.Year)
			}
			if result.Title != tt.expected.Title {
				t.Errorf("Title = %q, want %q", result.Title, tt.expected.Title)
			}
			if result.Description != tt.expected.Description {
				t.Errorf("Description = %q, want %q", result.Description, tt.expected.Description)
			}
			if result.Category != tt.expected.Category {
				t.Errorf("Category = %q, want %q", result.Category, tt.expected.Category)
			}
			if result.Link != tt.expected.Link {
				t.Errorf("Link = %q, want %q", result.Link, tt.expected.Link)
			}
		})
	}
}

func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Error("NewParser() returned nil")
	}
	if parser.client == nil {
		t.Error("Parser client is nil")
	}
}
