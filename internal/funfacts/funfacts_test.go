package funfacts

import (
	"testing"
)

func TestGetRandomFunFact(t *testing.T) {
	testCases := []struct {
		name                  string
		includeEmo            bool
		includeBlobby         bool
		includeWikiHow        bool
		includeWikiHowQuizzes bool
		includeQuote          bool
		expectNil             bool
	}{
		{"All enabled", true, true, true, true, true, false},
		{"Emo, Blobby, WikiHow", true, true, true, false, false, false},
		{"Only emo", true, false, false, false, false, false},
		{"Only blobby", false, true, false, false, false, false},
		{"Only wikihow", false, false, true, false, false, false},
		{"Only wikihow_quizzes", false, false, false, true, false, false},
		{"Only quote", false, false, false, false, true, false},
		{"None enabled", false, false, false, false, false, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fact := GetRandomFunFact(tc.includeEmo, tc.includeBlobby, tc.includeWikiHow, tc.includeWikiHowQuizzes, tc.includeQuote, false, false, false, false, false, false, false)

			if tc.expectNil {
				if fact != nil {
					t.Error("Expected nil fact when nothing is enabled")
				}
				return
			}

			if fact == nil {
				t.Error("Expected non-nil fact")
				return
			}

			if fact.Text == "" {
				t.Error("Expected non-empty text")
			}

			if fact.Type == "" {
				t.Error("Expected non-empty type")
			}
		})
	}
}

func TestGetRandomFunFactConsistency(t *testing.T) {
	// Should return same fact for same day
	fact1 := GetRandomFunFact(true, true, true, false, false, false, false, false, false, false, false, false)
	fact2 := GetRandomFunFact(true, true, true, false, false, false, false, false, false, false, false, false)

	if fact1 == nil || fact2 == nil {
		t.Fatal("Expected non-nil facts")
	}

	if fact1.Text != fact2.Text {
		t.Error("Expected same fact for the same day")
	}

	if fact1.Type != fact2.Type {
		t.Error("Expected same type for the same day")
	}
}

func TestGetDisplayTitle(t *testing.T) {
	testCases := []struct {
		factType      string
		expectedTitle string
	}{
		{"emo", "💭 Today's Thought"},
		{"blobby", "🎀 Mr Blobby Fact of the Day"},
		{"wikihow", "📚 Helpful WikiHow Article"},
		{"wikihow_quizzes", "🧠 WikiHow Quiz of the Day"},
		{"camping", "🏕️ Camping Tip of the Day"},
		{"joke", "💡 Fun Fact"},
		{"unknown", "💡 Fun Fact"},
	}

	for _, tc := range testCases {
		t.Run(tc.factType, func(t *testing.T) {
			fact := FunFact{Type: tc.factType}
			title := fact.GetDisplayTitle()

			if title != tc.expectedTitle {
				t.Errorf("Expected title '%s', got '%s'", tc.expectedTitle, title)
			}
		})
	}
}

func TestShouldDisplayAsItalic(t *testing.T) {
	testCases := []struct {
		factType       string
		expectedItalic bool
	}{
		{"emo", true},
		{"blobby", false},
		{"unknown", false},
	}

	for _, tc := range testCases {
		t.Run(tc.factType, func(t *testing.T) {
			fact := FunFact{Type: tc.factType}
			italic := fact.ShouldDisplayAsItalic()

			if italic != tc.expectedItalic {
				t.Errorf("Expected italic=%v, got %v", tc.expectedItalic, italic)
			}
		})
	}
}
