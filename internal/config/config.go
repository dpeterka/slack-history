package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds the application configuration
type Config struct {
	// Slack configuration
	SlackWebhookURL string

	// Anthropic Claude API configuration
	ClaudeAPIKey string
	ClaudeModel  string

	// Giphy API configuration
	GiphyAPIKey string

	// RSS feed URLs
	RSSFeedURLs []string

	// Holiday feed URL
	HolidayFeedURL string

	// Scheduler configuration
	ScheduleCron   string // Cron expression for scheduling
	RunOnce        bool   // Run once and exit (for testing)
	SkipInitialRun bool   // Skip running immediately on startup, only use schedule

	// LLM prompt configuration
	MaxEvents            int // Maximum number of events to select
	MaxHolidays          int // Maximum number of holidays to display
	EventSelectionPrompt string

	// New feature flags
	IncludeQuote          bool   // Include quote of the day
	IncludePeople         bool   // Include notable births/deaths
	IncludeEmoComment     bool   // Include emo/philosophical comment
	IncludeBlobbyFact     bool   // Include Mr Blobby fact
	IncludeWikiHow        bool   // Include WikiHow articles
	IncludeWikiHowQuizzes bool   // Include WikiHow quizzes
	IncludeHotTub         bool   // Include hot tub care tips
	IncludeGardening      bool   // Include gardening tips
	IncludePrinting3D     bool   // Include 3D printing tips
	IncludeCamping        bool   // Include camping tips
	IncludeJoke           bool   // Include daily joke (rendered bare, no title)
	IncludeFoodTakes      bool   // Include unhinged food take
	IncludeEvents         bool   // Include historical events
	MaxPeople             int    // Maximum number of people to display
	CacheDir              string // Directory for event history cache
	ContentRotationWeeks  int    // Number of weeks before repeating content
	TestDateSeed          int    // Override date seed for testing (0 = use current date)
	AIGeneratedContent    bool   // Generate the day's text content fresh daily (falls back to static packs)
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		SlackWebhookURL:       os.Getenv("SLACK_WEBHOOK_URL"),
		ClaudeAPIKey:          os.Getenv("CLAUDE_API_KEY"),
		ClaudeModel:           getEnvOrDefault("CLAUDE_MODEL", "claude-sonnet-4-5"),
		GiphyAPIKey:           os.Getenv("GIPHY_API_KEY"),
		ScheduleCron:          getEnvOrDefault("SCHEDULE_CRON", "0 9 * * *"), // Default: 9 AM daily
		RunOnce:               getEnvBool("RUN_ONCE", false),
		SkipInitialRun:        getEnvBool("SKIP_INITIAL_RUN", false),
		MaxEvents:             getEnvInt("MAX_EVENTS", 1),
		MaxHolidays:           getEnvInt("MAX_HOLIDAYS", 2),
		IncludeQuote:          getEnvBool("INCLUDE_QUOTE", true),
		IncludePeople:         getEnvBool("INCLUDE_PEOPLE", true),
		IncludeEmoComment:     getEnvBool("INCLUDE_EMO_COMMENT", true),
		IncludeBlobbyFact:     getEnvBool("INCLUDE_BLOBBY_FACT", true),
		IncludeWikiHow:        getEnvBool("INCLUDE_WIKIHOW", true),
		IncludeWikiHowQuizzes: getEnvBool("INCLUDE_WIKIHOW_QUIZZES", true),
		IncludeHotTub:         getEnvBool("INCLUDE_HOTTUB", true),
		IncludeGardening:      getEnvBool("INCLUDE_GARDENING", true),
		IncludePrinting3D:     getEnvBool("INCLUDE_PRINTING3D", true),
		IncludeCamping:        getEnvBool("INCLUDE_CAMPING", true),
		IncludeJoke:           getEnvBool("INCLUDE_JOKE", true),
		IncludeFoodTakes:      getEnvBool("INCLUDE_FOOD_TAKES", true),
		IncludeEvents:         getEnvBool("INCLUDE_EVENTS", true),
		MaxPeople:             getEnvInt("MAX_PEOPLE", 2),
		CacheDir:              getEnvOrDefault("CACHE_DIR", ".cache"),
		ContentRotationWeeks:  getEnvInt("CONTENT_ROTATION_WEEKS", 6),
		TestDateSeed:          getEnvInt("TEST_DATE_SEED", 0),
		AIGeneratedContent:    getEnvBool("AI_GENERATED_CONTENT", false),
	}

	// RSS feed URLs - support multiple feeds (comma-separated)
	feedURLs := getEnvOrDefault("RSS_FEED_URLS", "https://www.onthisday.com/rss/today-in-history.xml,https://unbelievablefactsblog.com/rss,http://feeds.feedburner.com/FutilityCloset,https://www.kickassfacts.com/feed/,https://www.mentalfloss.com/feed")
	cfg.RSSFeedURLs = parseCommaSeparated(feedURLs)

	// Holiday feed URL
	cfg.HolidayFeedURL = getEnvOrDefault("HOLIDAY_FEED_URL", "https://api.checkiday.com/rss?tz=America/New_York")

	// Default event selection prompt
	cfg.EventSelectionPrompt = getEnvOrDefault("EVENT_SELECTION_PROMPT",
		`You are analyzing historical events that happened on this day. Your task is to select the most interesting, unusual, and surprising events that would make people say "wow, I didn't know that!"

PRIMARY SELECTION CRITERIA (in order of importance):
1. **UNUSUAL & UNEXPECTED**: Prioritize bizarre, quirky, counterintuitive, or surprising events that defy expectations
2. **CULTURAL CURIOSITIES**: Events involving strange customs, unusual inventions, bizarre laws, peculiar traditions, or odd historical figures
3. **SURPRISING FIRSTS**: Unexpected "first time" events, especially if they seem ahead of their time or oddly specific
4. **HISTORICAL ODDITIES**: Events that seem absurd in hindsight, accidental discoveries, strange coincidences, or ironic outcomes
5. **LESSER-KNOWN GEMS**: Fascinating stories that most people haven't heard of - avoid obvious major events everyone knows

STRONGLY AVOID:
- Typical wars, battles, and military conquests (unless truly bizarre or unexpected)
- Common political elections/inaugurations (unless historically weird)
- Tragic disasters, terrorist attacks, mass casualties
- Recent events (last 10 years) unless extraordinarily unusual
- Obvious "textbook" historical events everyone knows

DIVERSITY REQUIREMENTS:
- **Time Period**: Prefer older events (pre-1950) as they tend to be more unusual
- **Geographic**: Include events from around the world - avoid US-centric selection
- **Category**: Mix science oddities, cultural quirks, arts/entertainment, bizarre inventions, strange laws, peculiar traditions
- **Tone**: Events should be fascinating, entertaining, or mind-blowing rather than somber

When selecting %d events, prioritize the MOST UNUSUAL and UNEXPECTED stories. Think "fascinating historical trivia" rather than "important historical milestones."

Format your response as JSON with the following structure:
{
  "events": [
    {
      "year": "YYYY",
      "title": "Brief event title",
      "description": "Engaging 2-3 sentence description with context and significance",
      "category": "Category of event (e.g., Science, Arts, Politics, Military, Sports, Culture, Technology, Social)"
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

// parseCommaSeparated parses a comma-separated string into a slice
func parseCommaSeparated(s string) []string {
	var result []string
	parts := strings.Split(s, ",")
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
