package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dpeterka/history-slackbot/internal/config"
	"github.com/dpeterka/history-slackbot/internal/llm"
	"github.com/dpeterka/history-slackbot/internal/rss"
	"github.com/dpeterka/history-slackbot/internal/scheduler"
	"github.com/dpeterka/history-slackbot/internal/slack"
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
	var sched *scheduler.Scheduler
	if cfg.RunOnce {
		// Run once and exit
		sched = scheduler.NewScheduler(job, 0, true)
	} else {
		// Parse cron expression and calculate next run time
		nextRun, err := scheduler.NextRunTime(cfg.ScheduleCron)
		if err != nil {
			log.Fatalf("Failed to parse cron expression: %v", err)
		}

		log.Printf("Next scheduled run: %v", nextRun)

		sched = scheduler.NewScheduler(job, scheduler.DailyInterval(), false)
	}

	// Setup signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start scheduler in a goroutine
	errChan := make(chan error, 1)
	go func() {
		if cfg.RunOnce {
			errChan <- sched.Start(ctx)
		} else {
			// Start at the next scheduled time
			nextRun, _ := scheduler.NextRunTime(cfg.ScheduleCron)
			errChan <- sched.StartAt(ctx, nextRun)
		}
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

// createJob creates the main job function
func createJob(cfg *config.Config) scheduler.Job {
	return func(ctx context.Context) error {
		log.Println("=== Starting job execution ===")

		// Create RSS parser
		parser := rss.NewParser()

		// Fetch events from RSS feeds
		log.Printf("Fetching events from %d feed(s)...", len(cfg.RSSFeedURLs))
		events, err := parser.FetchMultipleFeeds(cfg.RSSFeedURLs)
		if err != nil {
			return err
		}
		log.Printf("Fetched %d events", len(events))

		// Select interesting events using LLM
		log.Println("Selecting interesting events using Claude...")
		selector := llm.NewSelector(cfg.ClaudeAPIKey, cfg.ClaudeModel, cfg.MaxEvents, cfg.EventSelectionPrompt)
		selectedEvents, err := selector.SelectEvents(events)
		if err != nil {
			return err
		}
		log.Printf("Selected %d events", len(selectedEvents))

		// Post to Slack
		log.Println("Posting to Slack...")
		poster := slack.NewPoster(cfg.SlackWebhookURL)
		if err := poster.PostEvents(selectedEvents); err != nil {
			return err
		}

		log.Println("Successfully posted to Slack!")
		log.Println("=== Job execution completed ===")

		return nil
	}
}
