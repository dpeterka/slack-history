package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dpeterka/history-slackbot/internal/cache"
	"github.com/dpeterka/history-slackbot/internal/config"
	"github.com/dpeterka/history-slackbot/internal/funfacts"
	"github.com/dpeterka/history-slackbot/internal/holidays"
	"github.com/dpeterka/history-slackbot/internal/llm"
	"github.com/dpeterka/history-slackbot/internal/rss"
	"github.com/dpeterka/history-slackbot/internal/scheduler"
	"github.com/dpeterka/history-slackbot/internal/slack"
	"github.com/dpeterka/history-slackbot/internal/wikipedia"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting History Slackbot...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Configuration loaded successfully")
	log.Printf("Model: %s", cfg.ClaudeModel)
	log.Printf("Max events: %d", cfg.MaxEvents)
	log.Printf("Schedule: %s", cfg.ScheduleCron)
	log.Printf("Run once: %v", cfg.RunOnce)

	// Create the job that fetches and posts events
	job := createJob(cfg)

	// Create scheduler
	sched := scheduler.NewScheduler(job, 0, cfg.RunOnce)

	// Setup signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start scheduler in a goroutine
	errChan := make(chan error, 1)
	go func() {
		// Convert "minute hour * * *" to cron format with seconds "0 minute hour * * *"
		cronExpr := "0 " + cfg.ScheduleCron
		errChan <- sched.StartWithCron(ctx, cronExpr)
	}()

	// Wait for shutdown signal or error
	select {
	case sig := <-sigChan:
		log.Printf("Received signal: %v", sig)
		cancel()
	case err := <-errChan:
		if err != nil && err != context.Canceled {
			log.Printf("Scheduler error: %v", err)
		}
	}

	log.Println("History Slackbot stopped")
}

// filterFunHolidays filters out serious/political holidays and keeps only fun ones
func filterFunHolidays(holidays []rss.Holiday) []rss.Holiday {
	// Keywords that indicate serious/political/religious holidays to skip
	seriousKeywords := []string{
		"International", "World", "National Awareness", "Day for",
		"Memorial", "Remembrance", "Victims", "Prevention",
		"Human Rights", "Peace", "Conflict", "War", "Violence",
		"Exploitation", "Poverty", "Hunger", "Disease",
		"Awareness Week", "Awareness Month", "Solidarity",
		"Against", "United Nations", "Commemoration",
	}

	var funHolidays []rss.Holiday
	for _, holiday := range holidays {
		isFun := true
		title := holiday.Title

		// Check if title contains any serious keywords
		for _, keyword := range seriousKeywords {
			if contains(title, keyword) {
				isFun = false
				break
			}
		}

		if isFun {
			funHolidays = append(funHolidays, holiday)
		}
	}

	return funHolidays
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		strings.Contains(strings.ToLower(s), strings.ToLower(substr)))
}

// createJob creates the main job function
func createJob(cfg *config.Config) scheduler.Job {
	return func(ctx context.Context) error {
		log.Println("=== Starting job execution ===")

		// Check if today is a weekend (skip posting on weekends)
		now := time.Now()
		weekday := now.Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			log.Printf("Skipping post - today is %s (weekend)", weekday)
			return nil
		}

		// Use configured max events
		maxEvents := cfg.MaxEvents
		log.Printf("Today is %s - posting %d event(s)", weekday, maxEvents)

		// Initialize event history cache
		eventHistory, err := cache.NewEventHistory(cfg.CacheDir)
		if err != nil {
			log.Printf("Warning: failed to initialize event cache: %v", err)
			eventHistory = nil
		} else {
			// Cleanup old events
			if err := eventHistory.Cleanup(cfg.ContentRotationWeeks * 2); err != nil {
				log.Printf("Warning: failed to cleanup cache: %v", err)
			}
		}

		// Create RSS parser
		parser := rss.NewParser()

		// Fetch events from RSS feeds
		log.Printf("Fetching events from %d feed(s)...", len(cfg.RSSFeedURLs))
		allEvents, err := parser.FetchMultipleFeeds(cfg.RSSFeedURLs)
		if err != nil {
			return err
		}
		log.Printf("Fetched %d events", len(allEvents))

		// Filter out recently posted events
		var events []rss.HistoricalEvent
		if eventHistory != nil {
			for _, event := range allEvents {
				if !eventHistory.WasRecentlyPosted(event.Title, event.Year, cfg.ContentRotationWeeks) {
					events = append(events, event)
				}
			}
			log.Printf("After filtering recent posts: %d events remain", len(events))
		} else {
			events = allEvents
		}

		// Select interesting events using LLM
		log.Println("Selecting interesting events using Claude...")
		selector := llm.NewSelector(cfg.ClaudeAPIKey, cfg.ClaudeModel, maxEvents, cfg.EventSelectionPrompt)
		selectedEvents, err := selector.SelectEvents(events)
		if err != nil {
			return err
		}
		log.Printf("Selected %d events", len(selectedEvents))

		// Fetch Wikipedia URLs for selected events
		log.Println("Fetching Wikipedia URLs for events...")
		for i := range selectedEvents {
			url, err := wikipedia.SearchWikipediaURL(selectedEvents[i].Title)
			if err != nil {
				log.Printf("Warning: failed to get Wikipedia URL for '%s': %v", selectedEvents[i].Title, err)
			} else {
				selectedEvents[i].WikiURL = url
			}
		}

		// Add selected events to cache
		if eventHistory != nil {
			for _, event := range selectedEvents {
				if err := eventHistory.AddEvent(event.Title, event.Year, event.Description); err != nil {
					log.Printf("Warning: failed to cache event: %v", err)
				}
			}
		}

		// Check for major holidays
		majorHoliday, hasMajorHoliday := holidays.GetTodaysMajorHoliday()
		if hasMajorHoliday {
			log.Printf("Today is a major holiday: %s", majorHoliday)
		}

		// Fetch holidays
		var funHolidays []rss.Holiday
		if cfg.HolidayFeedURL != "" {
			log.Println("Fetching fun holidays...")
			holidayData, err := parser.FetchHolidays(cfg.HolidayFeedURL)
			if err != nil {
				log.Printf("Warning: failed to fetch holidays: %v", err)
			} else {
				log.Printf("Fetched %d holidays", len(holidayData))
				// Filter for fun holidays (skip serious/political ones)
				filteredHolidays := filterFunHolidays(holidayData)
				log.Printf("Filtered to %d fun holidays", len(filteredHolidays))

				// Limit to MaxHolidays
				maxCount := cfg.MaxHolidays
				if maxCount > len(filteredHolidays) {
					maxCount = len(filteredHolidays)
				}
				for i := 0; i < maxCount; i++ {
					funHolidays = append(funHolidays, filteredHolidays[i])
				}
				log.Printf("Selected %d holidays to display", len(funHolidays))
			}
		}

		// Fetch notable people from Wikipedia API (for rotation)
		var notablePeople []wikipedia.Person
		if cfg.IncludePeople {
			log.Println("Fetching notable births and deaths from Wikipedia...")
			now := time.Now()
			people, err := wikipedia.FetchBirthsAndDeaths(int(now.Month()), now.Day(), cfg.MaxPeople)
			if err != nil {
				log.Printf("Warning: failed to fetch people from Wikipedia: %v", err)
			} else {
				notablePeople = people
				log.Printf("Fetched %d notable people from Wikipedia", len(notablePeople))
			}
		}

		// Select daily variety content (emo, blobby, wikihow, quote, hottub, gardening, people, events, holidays - all rotate)
		var funFact *funfacts.FunFact
		log.Println("Selecting daily variety content...")
		funFact = funfacts.GetRandomFunFactWithData(
			cfg.IncludeEmoComment,
			cfg.IncludeBlobbyFact,
			cfg.IncludeWikiHow,
			cfg.IncludeQuote,
			cfg.IncludeHotTub,
			cfg.IncludeGardening,
			cfg.IncludePeople,
			cfg.IncludeEvents,
			funHolidays,    // Pass holidays for rotation
			notablePeople,  // Pass people for rotation
			cfg.MaxPeople,
			cfg.TestDateSeed,
		)
		if funFact != nil {
			log.Printf("Selected %s", funFact.GetDisplayTitle())
			if funFact.Text != "" {
				log.Printf("Content: %s", funFact.Text)
			}
		}

		// Post to Slack - only show events if rotation selected them
		log.Println("Posting to Slack...")
		poster := slack.NewPoster(cfg.SlackWebhookURL)
		var eventsToPost []llm.SelectedEvent
		if funFact != nil && funFact.ShowEvents {
			eventsToPost = selectedEvents
			log.Printf("Posting %d events for Historical Event day", len(eventsToPost))
		} else {
			log.Println("Not posting events (rotation selected different content)")
		}
		if err := poster.PostComplete(eventsToPost, majorHoliday, nil, funFact); err != nil {
			return err
		}

		log.Println("Successfully posted to Slack!")
		log.Println("=== Job execution completed ===")

		return nil
	}
}
