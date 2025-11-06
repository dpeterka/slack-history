package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds the application configuration
type Config struct {
	// Slack configuration
	SlackWebhookURL string

	// Anthropic Claude API configuration
	ClaudeAPIKey string
	ClaudeModel  string

	// RSS feed URLs
	RSSFeedURLs []string

	// Scheduler configuration
	ScheduleCron string // Cron expression for scheduling
	RunOnce      bool   // Run once and exit (for testing)

	// LLM prompt configuration
	MaxEvents         int // Maximum number of events to select
	EventSelectionPrompt string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		SlackWebhookURL: os.Getenv("SLACK_WEBHOOK_URL"),
		ClaudeAPIKey:    os.Getenv("CLAUDE_API_KEY"),
		ClaudeModel:     getEnvOrDefault("CLAUDE_MODEL", "claude-sonnet-4-5"),
		ScheduleCron:    getEnvOrDefault("SCHEDULE_CRON", "0 9 * * *"), // Default: 9 AM daily
		RunOnce:         getEnvBool("RUN_ONCE", false),
		MaxEvents:       getEnvInt("MAX_EVENTS", 2),
	}

	// RSS feed URLs - support multiple feeds
	feedURL := getEnvOrDefault("RSS_FEED_URL", "https://www.onthisday.com/rss/today-in-history.xml")
	cfg.RSSFeedURLs = []string{feedURL}

	// Default event selection prompt
	cfg.EventSelectionPrompt = getEnvOrDefault("EVENT_SELECTION_PROMPT",
		`You are analyzing historical events that happened on this day. Your task is to select the most interesting, rare, or significant events from the list provided.

Criteria for selection:
- Events that are historically significant or impactful
- Unusual, rare, or surprising events
- Events that would be interesting to a general audience
- Avoid overly common or mundane events
- Prefer events from different time periods and categories for variety

Select exactly %d events from the list and format them for posting to Slack. For each event, provide:
1. The year and a brief, engaging description (2-3 sentences max)
2. Why this event is interesting or significant

Format your response as JSON with the following structure:
{
  "events": [
    {
      "year": "YYYY",
      "title": "Brief event title",
      "description": "Engaging 2-3 sentence description with context and significance",
      "category": "Category of event (e.g., Politics, Science, Arts, etc.)"
    }
  ]
}`)

	// Validate required configuration
	if cfg.SlackWebhookURL == "" {
		return nil, fmt.Errorf("SLACK_WEBHOOK_URL is required")
	}
	if cfg.ClaudeAPIKey == "" {
		return nil, fmt.Errorf("CLAUDE_API_KEY is required")
	}

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

// GetSchedule returns the next scheduled run time
func (c *Config) GetSchedule() (time.Duration, error) {
	// For now, we'll implement a simple daily schedule
	// In a production system, you'd use a cron parser library
	now := time.Now()

	// Parse the cron expression (simplified - assumes "0 9 * * *" format)
	// For a full implementation, use github.com/robfig/cron
	hour := 9
	if c.ScheduleCron != "" {
		// Basic parsing for hour only
		// Format: "minute hour * * *"
		var minute int
		fmt.Sscanf(c.ScheduleCron, "%d %d", &minute, &hour)
	}

	// Calculate next run time
	nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
	if now.After(nextRun) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return nextRun.Sub(now), nil
}
