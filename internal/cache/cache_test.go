package cache

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewEventHistory(t *testing.T) {
	tmpDir := t.TempDir()

	eh, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create event history: %v", err)
	}

	if eh == nil {
		t.Fatal("Expected event history, got nil")
	}

	if eh.Count() != 0 {
		t.Errorf("Expected 0 events, got %d", eh.Count())
	}
}

func TestAddAndCheckEvent(t *testing.T) {
	tmpDir := t.TempDir()

	eh, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create event history: %v", err)
	}

	// Add an event
	err = eh.AddEvent("Moon Landing", "1969", "Apollo 11 lands on the moon")
	if err != nil {
		t.Fatalf("Failed to add event: %v", err)
	}

	// Check it was added
	if eh.Count() != 1 {
		t.Errorf("Expected 1 event, got %d", eh.Count())
	}

	// Check it's marked as recently posted
	if !eh.WasRecentlyPosted("Moon Landing", "1969", 1) {
		t.Error("Expected event to be marked as recently posted")
	}

	// Check different event is not marked as posted
	if eh.WasRecentlyPosted("Different Event", "2000", 1) {
		t.Error("Expected different event to not be marked as posted")
	}
}

func TestPersistence(t *testing.T) {
	tmpDir := t.TempDir()

	// Create and add event
	eh1, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create event history: %v", err)
	}

	err = eh1.AddEvent("Moon Landing", "1969", "Apollo 11 lands on the moon")
	if err != nil {
		t.Fatalf("Failed to add event: %v", err)
	}

	// Create new instance and check event persisted
	eh2, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create second event history: %v", err)
	}

	if eh2.Count() != 1 {
		t.Errorf("Expected 1 event after reload, got %d", eh2.Count())
	}

	if !eh2.WasRecentlyPosted("Moon Landing", "1969", 1) {
		t.Error("Expected event to persist after reload")
	}
}

func TestCleanup(t *testing.T) {
	tmpDir := t.TempDir()

	eh, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create event history: %v", err)
	}

	// Add events with different timestamps
	oldEvent := EventRecord{
		Title:       "Old Event",
		Year:        "1900",
		Description: "This is old",
		PostedAt:    time.Now().AddDate(0, 0, -60), // 60 days ago
	}

	recentEvent := EventRecord{
		Title:       "Recent Event",
		Year:        "2020",
		Description: "This is recent",
		PostedAt:    time.Now(),
	}

	eh.AddEvents([]EventRecord{oldEvent, recentEvent})

	if eh.Count() != 2 {
		t.Errorf("Expected 2 events before cleanup, got %d", eh.Count())
	}

	// Cleanup events older than 4 weeks
	err = eh.Cleanup(4)
	if err != nil {
		t.Fatalf("Failed to cleanup: %v", err)
	}

	if eh.Count() != 1 {
		t.Errorf("Expected 1 event after cleanup, got %d", eh.Count())
	}

	// Check the correct event remains
	if !eh.WasRecentlyPosted("Recent Event", "2020", 1) {
		t.Error("Expected recent event to remain after cleanup")
	}

	if eh.WasRecentlyPosted("Old Event", "1900", 1) {
		t.Error("Expected old event to be removed after cleanup")
	}
}

func TestGetRecentEvents(t *testing.T) {
	tmpDir := t.TempDir()

	eh, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create event history: %v", err)
	}

	// Add events with different timestamps
	oldEvent := EventRecord{
		Title:       "Old Event",
		Year:        "1900",
		Description: "This is old",
		PostedAt:    time.Now().AddDate(0, 0, -60),
	}

	recentEvent := EventRecord{
		Title:       "Recent Event",
		Year:        "2020",
		Description: "This is recent",
		PostedAt:    time.Now(),
	}

	eh.AddEvents([]EventRecord{oldEvent, recentEvent})

	recent := eh.GetRecentEvents(4)
	if len(recent) != 1 {
		t.Errorf("Expected 1 recent event, got %d", len(recent))
	}

	if recent[0].Title != "Recent Event" {
		t.Errorf("Expected 'Recent Event', got '%s'", recent[0].Title)
	}
}

func TestCacheFileCreation(t *testing.T) {
	tmpDir := t.TempDir()

	eh, err := NewEventHistory(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create event history: %v", err)
	}

	err = eh.AddEvent("Test Event", "2024", "Test description")
	if err != nil {
		t.Fatalf("Failed to add event: %v", err)
	}

	// Check file exists
	cacheFile := filepath.Join(tmpDir, "event_history.json")
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		t.Error("Expected cache file to be created")
	}
}
