package llm

import (
	"testing"

	"github.com/dpeterka/history-slackbot/internal/rss"
)

func TestNewSelector(t *testing.T) {
	selector := NewSelector("test-key", "test-model", 2, "test-prompt")

	if selector == nil {
		t.Error("NewSelector() returned nil")
	}
	if selector.apiKey != "test-key" {
		t.Errorf("apiKey = %q, want %q", selector.apiKey, "test-key")
	}
	if selector.model != "test-model" {
		t.Errorf("model = %q, want %q", selector.model, "test-model")
	}
	if selector.maxEvents != 2 {
		t.Errorf("maxEvents = %d, want %d", selector.maxEvents, 2)
	}
}

func TestFormatEventsForPrompt(t *testing.T) {
	selector := NewSelector("test-key", "test-model", 2, "test-prompt")

	events := []rss.HistoricalEvent{
		{
			Year:        "1969",
			Title:       "Apollo 11 Moon Landing",
			Description: "First humans land on the Moon",
			Category:    "Science",
		},
		{
			Year:        "1776",
			Title:       "Declaration of Independence",
			Description: "United States declares independence",
			Category:    "Politics",
		},
	}

	result := selector.formatEventsForPrompt(events)

	if result == "" {
		t.Error("formatEventsForPrompt() returned empty string")
	}

	// Check that key information is included
	tests := []string{
		"1969",
		"Apollo 11 Moon Landing",
		"First humans land on the Moon",
		"Science",
		"1776",
		"Declaration of Independence",
		"United States declares independence",
		"Politics",
	}

	for _, want := range tests {
		if !contains(result, want) {
			t.Errorf("formatEventsForPrompt() missing %q", want)
		}
	}
}

func TestExtractJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Plain JSON",
			input:    `{"events": []}`,
			expected: `{"events": []}`,
		},
		{
			name: "JSON in markdown code block with json tag",
			input: "```json\n{\"events\": []}\n```",
			expected: "\n{\"events\": []}\n",
		},
		{
			name: "JSON in markdown code block without tag",
			input: "```\n{\"events\": []}\n```",
			expected: "\n{\"events\": []}\n",
		},
		{
			name:     "No code blocks",
			input:    "Some text with {\"events\": []} embedded",
			expected: "Some text with {\"events\": []} embedded",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSON(tt.input)
			if result != tt.expected {
				t.Errorf("extractJSON() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestParseSelection(t *testing.T) {
	selector := NewSelector("test-key", "test-model", 2, "test-prompt")

	tests := []struct {
		name        string
		response    string
		expectError bool
		validate    func(t *testing.T, events []SelectedEvent)
	}{
		{
			name: "Valid JSON response",
			response: `{
				"events": [
					{
						"year": "1969",
						"title": "Apollo 11",
						"description": "Moon landing",
						"category": "Science"
					}
				]
			}`,
			expectError: false,
			validate: func(t *testing.T, events []SelectedEvent) {
				if len(events) != 1 {
					t.Errorf("len(events) = %d, want 1", len(events))
				}
				if events[0].Year != "1969" {
					t.Errorf("Year = %q, want %q", events[0].Year, "1969")
				}
				if events[0].Title != "Apollo 11" {
					t.Errorf("Title = %q, want %q", events[0].Title, "Apollo 11")
				}
			},
		},
		{
			name: "Valid JSON in markdown",
			response: "```json\n{\"events\": [{\"year\": \"1776\", \"title\": \"Independence\", \"description\": \"US declares independence\", \"category\": \"Politics\"}]}\n```",
			expectError: false,
			validate: func(t *testing.T, events []SelectedEvent) {
				if len(events) != 1 {
					t.Errorf("len(events) = %d, want 1", len(events))
				}
			},
		},
		{
			name:        "Invalid JSON",
			response:    "not json",
			expectError: true,
		},
		{
			name:        "Empty events array",
			response:    `{"events": []}`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events, err := selector.parseSelection(tt.response)

			if tt.expectError {
				if err == nil {
					t.Error("parseSelection() should return error")
				}
			} else {
				if err != nil {
					t.Errorf("parseSelection() returned unexpected error: %v", err)
				}
				if tt.validate != nil {
					tt.validate(t, events)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && (s[0:len(substr)] == substr || contains(s[1:], substr))))
}
