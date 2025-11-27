package blobby

import (
	"testing"
)

func TestGetRandomFact(t *testing.T) {
	fact := GetRandomFact()

	if fact.Text == "" {
		t.Error("Expected non-empty fact text")
	}

	if fact.Category == "" {
		t.Error("Expected non-empty category")
	}
}

func TestGetRandomFactConsistency(t *testing.T) {
	// Should return same fact for same day
	fact1 := GetRandomFact()
	fact2 := GetRandomFact()

	if fact1.Text != fact2.Text {
		t.Error("Expected same fact for the same day")
	}

	if fact1.Category != fact2.Category {
		t.Error("Expected same category for the same day")
	}
}

func TestGetRandomFactByCategory(t *testing.T) {
	categories := []string{"Origins", "Music", "Theme Park", "TV", "Legacy"}

	for _, category := range categories {
		fact := GetRandomFactByCategory(category)

		if fact.Text == "" {
			t.Errorf("Expected non-empty fact for category %s", category)
		}

		if fact.Category != category {
			t.Errorf("Expected category %s, got %s", category, fact.Category)
		}
	}
}

func TestGetRandomFactByInvalidCategory(t *testing.T) {
	// Should fall back to any random fact
	fact := GetRandomFactByCategory("InvalidCategory")

	if fact.Text == "" {
		t.Error("Expected fallback fact")
	}
}

func TestAllFactsHaveCategory(t *testing.T) {
	facts := getAllFacts()

	validCategories := map[string]bool{
		"Origins":         true,
		"Music":           true,
		"Theme Park":      true,
		"TV":              true,
		"Legacy":          true,
		"Controversies":   true,
		"Commercial":      true,
		"Modern":          true,
		"Design":          true,
		"Academic":        true,
		"Almost Happened": true,
		"Influence":       true,
		"Creator":         true,
	}

	for i, fact := range facts {
		if fact.Text == "" {
			t.Errorf("Fact %d has empty text", i)
		}

		if !validCategories[fact.Category] {
			t.Errorf("Fact %d has invalid category: %s", i, fact.Category)
		}
	}
}

func TestMinimumNumberOfFacts(t *testing.T) {
	facts := getAllFacts()

	// Should have at least 30 facts
	if len(facts) < 30 {
		t.Errorf("Expected at least 30 facts, got %d", len(facts))
	}
}
