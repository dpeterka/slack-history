package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// EventHistory tracks posted events to prevent repeats
type EventHistory struct {
	filePath string
	history  *History
}

// History represents the cache file structure
type History struct {
	Events []EventRecord `json:"events"`
}

// EventRecord represents a single posted event
type EventRecord struct {
	Title       string    `json:"title"`
	Year        string    `json:"year"`
	PostedAt    time.Time `json:"posted_at"`
	Description string    `json:"description"`
}

// NewEventHistory creates a new event history tracker
func NewEventHistory(cacheDir string) (*EventHistory, error) {
	if cacheDir == "" {
		cacheDir = ".cache"
	}

	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	filePath := filepath.Join(cacheDir, "event_history.json")

	eh := &EventHistory{
		filePath: filePath,
		history:  &History{Events: []EventRecord{}},
	}

	// Load existing history
	if err := eh.load(); err != nil {
		// If file doesn't exist, that's okay - we'll create it on first save
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load history: %w", err)
		}
	}

	return eh, nil
}

// load loads the history from disk
func (eh *EventHistory) load() error {
	data, err := os.ReadFile(eh.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, eh.history)
}

// save saves the history to disk
func (eh *EventHistory) save() error {
	data, err := json.MarshalIndent(eh.history, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal history: %w", err)
	}

	return os.WriteFile(eh.filePath, data, 0644)
}

// WasRecentlyPosted checks if an event was posted within the specified number of weeks
func (eh *EventHistory) WasRecentlyPosted(title, year string, withinWeeks int) bool {
	cutoffDate := time.Now().AddDate(0, 0, -7*withinWeeks)

	for _, event := range eh.history.Events {
		if event.Title == title && event.Year == year && event.PostedAt.After(cutoffDate) {
			return true
		}
	}

	return false
}

// AddEvent adds an event to the history
func (eh *EventHistory) AddEvent(title, year, description string) error {
	record := EventRecord{
		Title:       title,
		Year:        year,
		Description: description,
		PostedAt:    time.Now(),
	}

	eh.history.Events = append(eh.history.Events, record)

	return eh.save()
}

// AddEvents adds multiple events to the history
func (eh *EventHistory) AddEvents(events []EventRecord) error {
	eh.history.Events = append(eh.history.Events, events...)
	return eh.save()
}

// Cleanup removes old events beyond the retention period
func (eh *EventHistory) Cleanup(retentionWeeks int) error {
	cutoffDate := time.Now().AddDate(0, 0, -7*retentionWeeks)

	// Filter out old events
	filtered := []EventRecord{}
	for _, event := range eh.history.Events {
		if event.PostedAt.After(cutoffDate) {
			filtered = append(filtered, event)
		}
	}

	eh.history.Events = filtered

	return eh.save()
}

// GetRecentEvents returns events posted within the specified number of weeks
func (eh *EventHistory) GetRecentEvents(withinWeeks int) []EventRecord {
	cutoffDate := time.Now().AddDate(0, 0, -7*withinWeeks)

	recent := []EventRecord{}
	for _, event := range eh.history.Events {
		if event.PostedAt.After(cutoffDate) {
			recent = append(recent, event)
		}
	}

	return recent
}

// Count returns the total number of events in history
func (eh *EventHistory) Count() int {
	return len(eh.history.Events)
}
