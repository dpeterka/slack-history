package emo

import (
	"testing"
)

func TestGetRandomComment(t *testing.T) {
	comment := GetRandomComment()

	if comment.Text == "" {
		t.Error("Expected non-empty comment text")
	}

	if comment.Category == "" {
		t.Error("Expected non-empty category")
	}
}

func TestGetRandomCommentConsistency(t *testing.T) {
	// Should return same comment for same day
	comment1 := GetRandomComment()
	comment2 := GetRandomComment()

	if comment1.Text != comment2.Text {
		t.Error("Expected same comment for the same day")
	}

	if comment1.Category != comment2.Category {
		t.Error("Expected same category for the same day")
	}
}

func TestGetRandomCommentByCategory(t *testing.T) {
	categories := []string{"Work", "Life", "Relationships"}

	for _, category := range categories {
		comment := GetRandomCommentByCategory(category)

		if comment.Text == "" {
			t.Errorf("Expected non-empty comment for category %s", category)
		}

		if comment.Category != category {
			t.Errorf("Expected category %s, got %s", category, comment.Category)
		}
	}
}

func TestGetRandomCommentByInvalidCategory(t *testing.T) {
	// Should fall back to any random comment
	comment := GetRandomCommentByCategory("InvalidCategory")

	if comment.Text == "" {
		t.Error("Expected fallback comment")
	}
}

func TestAllCommentsHaveCategory(t *testing.T) {
	comments := getAllComments()

	validCategories := map[string]bool{
		"Work":          true,
		"Life":          true,
		"Relationships": true,
	}

	for i, comment := range comments {
		if comment.Text == "" {
			t.Errorf("Comment %d has empty text", i)
		}

		if !validCategories[comment.Category] {
			t.Errorf("Comment %d has invalid category: %s", i, comment.Category)
		}
	}
}
