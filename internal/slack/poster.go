package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dpeterka/history-slackbot/internal/llm"
)

// Poster handles posting messages to Slack
type Poster struct {
	webhookURL string
	client     *http.Client
}

// NewPoster creates a new Slack poster
func NewPoster(webhookURL string) *Poster {
	return &Poster{
		webhookURL: webhookURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SlackMessage represents a Slack message
type SlackMessage struct {
	Text        string       `json:"text,omitempty"`
	Blocks      []Block      `json:"blocks,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Block represents a Slack block
type Block struct {
	Type     string        `json:"type"`
	Text     *TextObject   `json:"text,omitempty"`
	Elements []TextObject  `json:"elements,omitempty"`
}

// TextObject represents a text object in Slack
type TextObject struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Attachment represents a Slack attachment
type Attachment struct {
	Color  string  `json:"color,omitempty"`
	Blocks []Block `json:"blocks,omitempty"`
}

// PostEvents posts selected events to Slack
func (p *Poster) PostEvents(events []llm.SelectedEvent) error {
	return p.PostEventsWithHolidays(events, nil)
}

// PostEventsWithHolidays posts selected events and holidays to Slack
func (p *Poster) PostEventsWithHolidays(events []llm.SelectedEvent, holidays []string) error {
	if len(events) == 0 && len(holidays) == 0 {
		return fmt.Errorf("no events or holidays to post")
	}

	message := p.formatMessageWithHolidays(events, holidays)

	reqBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	req, err := http.NewRequest("POST", p.webhookURL, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// formatMessage formats events into a Slack message with blocks
func (p *Poster) formatMessage(events []llm.SelectedEvent) SlackMessage {
	return p.formatMessageWithHolidays(events, nil)
}

// formatMessageWithHolidays formats events and holidays into a Slack message with blocks
func (p *Poster) formatMessageWithHolidays(events []llm.SelectedEvent, holidays []string) SlackMessage {
	now := time.Now()
	dateStr := now.Format("Monday, January 2")

	// Create header block
	blocks := []Block{
		{
			Type: "header",
			Text: &TextObject{
				Type: "plain_text",
				Text: fmt.Sprintf("📅 On This Day in History - %s", dateStr),
			},
		},
		{
			Type: "divider",
		},
	}

	// Add holidays section if present
	if len(holidays) > 0 {
		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: "*🎉 Today's Fun Holidays*",
			},
		})

		// Add each holiday
		holidayText := ""
		for i, holiday := range holidays {
			if i > 0 {
				holidayText += "\n"
			}
			holidayText += fmt.Sprintf("• %s", holiday)
		}

		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: holidayText,
			},
		})

		blocks = append(blocks, Block{
			Type: "divider",
		})
	}

	// Add each event as a section
	for i, event := range events {
		// Event header with year and category
		header := fmt.Sprintf("*%s* • %s", event.Year, event.Category)

		// Event title
		titleText := fmt.Sprintf("*%s*", event.Title)

		// Full event block
		eventText := fmt.Sprintf("%s\n\n%s\n\n%s", header, titleText, event.Description)

		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: eventText,
			},
		})

		// Add divider between events (but not after the last one)
		if i < len(events)-1 {
			blocks = append(blocks, Block{
				Type: "divider",
			})
		}
	}

	// Add footer
	blocks = append(blocks, Block{
		Type: "context",
		Elements: []TextObject{
			{
				Type: "mrkdwn",
				Text: "_Curated by AI from today's historical events_",
			},
		},
	})

	return SlackMessage{
		Blocks: blocks,
	}
}

// PostSimpleMessage posts a simple text message to Slack
func (p *Poster) PostSimpleMessage(text string) error {
	message := SlackMessage{
		Text: text,
	}

	reqBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	req, err := http.NewRequest("POST", p.webhookURL, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// FormatEventsAsText formats events as plain text (for testing or simple posts)
func FormatEventsAsText(events []llm.SelectedEvent) string {
	var buf strings.Builder

	now := time.Now()
	dateStr := now.Format("Monday, January 2, 2006")

	buf.WriteString(fmt.Sprintf("📅 On This Day in History - %s\n\n", dateStr))

	for i, event := range events {
		buf.WriteString(fmt.Sprintf("%d. %s - %s\n", i+1, event.Year, event.Title))
		buf.WriteString(fmt.Sprintf("   Category: %s\n", event.Category))
		buf.WriteString(fmt.Sprintf("   %s\n", event.Description))
		if i < len(events)-1 {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
