package birthday

import (
	"testing"
	"time"
)

func TestIsBotBirthday(t *testing.T) {
	// This test will only pass on January 8th
	now := time.Now()
	result := IsBotBirthday()
	expected := now.Month() == time.January && now.Day() == 8

	if result != expected {
		t.Errorf("IsBotBirthday() = %v, expected %v (today is %s)", result, expected, now.Format("January 2"))
	}
}

func TestGetBotAge(t *testing.T) {
	age := GetBotAge()
	currentYear := time.Now().Year()
	expectedAge := currentYear - 2026

	if age != expectedAge {
		t.Errorf("GetBotAge() = %d, expected %d", age, expectedAge)
	}
}

func TestGetBirthdayMessage(t *testing.T) {
	// Test without Giphy API key
	msg := GetBirthdayMessage("")

	if msg == nil {
		t.Fatal("GetBirthdayMessage() returned nil")
	}

	if msg.Title == "" {
		t.Error("Birthday message title is empty")
	}

	if msg.Message == "" {
		t.Error("Birthday message text is empty")
	}

	if msg.YouTubeURL == "" {
		t.Error("Birthday message YouTube URL is empty")
	}

	if msg.BirthdayAge < 0 {
		t.Errorf("Birthday age is negative: %d", msg.BirthdayAge)
	}

	// GiphyURL can be empty if no API key provided
	t.Logf("Birthday message: %s", msg.Message)
	t.Logf("YouTube URL: %s", msg.YouTubeURL)
	t.Logf("Giphy URL: %s", msg.GiphyURL)
	t.Logf("Bot age: %d", msg.BirthdayAge)
}
