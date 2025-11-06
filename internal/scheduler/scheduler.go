package scheduler

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Job represents a scheduled job
type Job func(ctx context.Context) error

// Scheduler handles scheduling of jobs
type Scheduler struct {
	job      Job
	interval time.Duration
	runOnce  bool
}

// NewScheduler creates a new scheduler
func NewScheduler(job Job, interval time.Duration, runOnce bool) *Scheduler {
	return &Scheduler{
		job:      job,
		interval: interval,
		runOnce:  runOnce,
	}
}

// Start starts the scheduler
func (s *Scheduler) Start(ctx context.Context) error {
	log.Printf("Scheduler starting...")

	// If runOnce is true, execute immediately and return
	if s.runOnce {
		log.Printf("Running job once...")
		if err := s.job(ctx); err != nil {
			return fmt.Errorf("job failed: %w", err)
		}
		log.Printf("Job completed successfully")
		return nil
	}

	// Otherwise, run on a schedule
	log.Printf("Scheduling job to run every %v", s.interval)

	// Run immediately on startup
	log.Printf("Running initial job...")
	if err := s.job(ctx); err != nil {
		log.Printf("Initial job failed: %v", err)
		// Continue with scheduling even if initial run fails
	} else {
		log.Printf("Initial job completed successfully")
	}

	// Create ticker for subsequent runs
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Scheduler stopping...")
			return ctx.Err()
		case <-ticker.C:
			log.Printf("Running scheduled job...")
			if err := s.job(ctx); err != nil {
				log.Printf("Scheduled job failed: %v", err)
				// Continue running even if job fails
			} else {
				log.Printf("Scheduled job completed successfully")
			}
		}
	}
}

// StartAt starts the scheduler with an initial delay
func (s *Scheduler) StartAt(ctx context.Context, firstRun time.Time) error {
	log.Printf("Scheduler will start at %v", firstRun)

	// Calculate delay until first run
	delay := time.Until(firstRun)
	if delay < 0 {
		// If the time has already passed today, schedule for tomorrow
		delay = delay + 24*time.Hour
	}

	log.Printf("Waiting %v until first run...", delay)

	// Wait for the first run time
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		// First run time reached, start the scheduler
		return s.Start(ctx)
	}
}

// ParseCron parses a simple cron expression and returns the next run time
// Format: "minute hour * * *" (e.g., "0 9 * * *" for 9:00 AM daily)
// This is a simplified implementation. For full cron support, use github.com/robfig/cron
func ParseCron(cronExpr string) (hour, minute int, err error) {
	var dayOfMonth, month, dayOfWeek string
	n, err := fmt.Sscanf(cronExpr, "%d %d %s %s %s", &minute, &hour, &dayOfMonth, &month, &dayOfWeek)
	if err != nil || n < 2 {
		return 0, 0, fmt.Errorf("invalid cron expression: %s", cronExpr)
	}

	if hour < 0 || hour > 23 {
		return 0, 0, fmt.Errorf("invalid hour: %d", hour)
	}
	if minute < 0 || minute > 59 {
		return 0, 0, fmt.Errorf("invalid minute: %d", minute)
	}

	return hour, minute, nil
}

// NextRunTime calculates the next run time based on a cron expression
func NextRunTime(cronExpr string) (time.Time, error) {
	hour, minute, err := ParseCron(cronExpr)
	if err != nil {
		return time.Time{}, err
	}

	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

	// If the time has already passed today, schedule for tomorrow
	if now.After(next) {
		next = next.Add(24 * time.Hour)
	}

	return next, nil
}

// DailyInterval returns the duration for daily scheduling
func DailyInterval() time.Duration {
	return 24 * time.Hour
}
