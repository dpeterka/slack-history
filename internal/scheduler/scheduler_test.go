package scheduler

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewScheduler(t *testing.T) {
	job := func(ctx context.Context) error { return nil }
	interval := 1 * time.Hour

	scheduler := NewScheduler(job, interval, false)

	if scheduler == nil {
		t.Error("NewScheduler() returned nil")
	}
	if scheduler.interval != interval {
		t.Errorf("interval = %v, want %v", scheduler.interval, interval)
	}
	if scheduler.runOnce != false {
		t.Errorf("runOnce = %v, want false", scheduler.runOnce)
	}
}

func TestSchedulerRunOnce(t *testing.T) {
	executed := false
	job := func(ctx context.Context) error {
		executed = true
		return nil
	}

	scheduler := NewScheduler(job, 0, true)
	ctx := context.Background()

	err := scheduler.Start(ctx)
	if err != nil {
		t.Errorf("Start() returned error: %v", err)
	}
	if !executed {
		t.Error("Job was not executed")
	}
}

func TestSchedulerRunOnceWithError(t *testing.T) {
	expectedErr := errors.New("job failed")
	job := func(ctx context.Context) error {
		return expectedErr
	}

	scheduler := NewScheduler(job, 0, true)
	ctx := context.Background()

	err := scheduler.Start(ctx)
	if err == nil {
		t.Error("Start() should return error")
	}
	if err.Error() != "job failed: job failed" {
		t.Errorf("Start() error = %v, want %v", err, expectedErr)
	}
}

func TestParseCron(t *testing.T) {
	tests := []struct {
		name        string
		cronExpr    string
		expectError bool
		hour        int
		minute      int
	}{
		{
			name:        "Valid cron 9 AM",
			cronExpr:    "0 9 * * *",
			expectError: false,
			hour:        9,
			minute:      0,
		},
		{
			name:        "Valid cron 8:30 AM",
			cronExpr:    "30 8 * * *",
			expectError: false,
			hour:        8,
			minute:      30,
		},
		{
			name:        "Valid cron midnight",
			cronExpr:    "0 0 * * *",
			expectError: false,
			hour:        0,
			minute:      0,
		},
		{
			name:        "Invalid hour",
			cronExpr:    "0 25 * * *",
			expectError: true,
		},
		{
			name:        "Invalid minute",
			cronExpr:    "60 9 * * *",
			expectError: true,
		},
		{
			name:        "Invalid format",
			cronExpr:    "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hour, minute, err := ParseCron(tt.cronExpr)

			if tt.expectError {
				if err == nil {
					t.Error("ParseCron() should return error")
				}
			} else {
				if err != nil {
					t.Errorf("ParseCron() returned unexpected error: %v", err)
				}
				if hour != tt.hour {
					t.Errorf("hour = %d, want %d", hour, tt.hour)
				}
				if minute != tt.minute {
					t.Errorf("minute = %d, want %d", minute, tt.minute)
				}
			}
		})
	}
}

func TestNextRunTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		cronExpr string
		validate func(t *testing.T, nextRun time.Time)
	}{
		{
			name:     "Next run in the future today",
			cronExpr: "59 23 * * *", // 11:59 PM
			validate: func(t *testing.T, nextRun time.Time) {
				if nextRun.Before(now) {
					t.Error("Next run should be in the future")
				}
				if nextRun.Hour() != 23 || nextRun.Minute() != 59 {
					t.Errorf("Next run time = %v, want 23:59", nextRun.Format("15:04"))
				}
			},
		},
		{
			name:     "Next run tomorrow (past time today)",
			cronExpr: "0 0 * * *", // Midnight
			validate: func(t *testing.T, nextRun time.Time) {
				if nextRun.Before(now) {
					t.Error("Next run should be in the future")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextRun, err := NextRunTime(tt.cronExpr)
			if err != nil {
				t.Errorf("NextRunTime() returned error: %v", err)
			}
			tt.validate(t, nextRun)
		})
	}
}

func TestDailyInterval(t *testing.T) {
	interval := DailyInterval()
	expected := 24 * time.Hour

	if interval != expected {
		t.Errorf("DailyInterval() = %v, want %v", interval, expected)
	}
}
