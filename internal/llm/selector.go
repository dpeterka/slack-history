package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dpeterka/history-slackbot/internal/rss"
)

// Selector uses an LLM to select interesting events
type Selector struct {
	apiKey      string
	model       string
	client      *http.Client
	maxEvents   int
	promptTemplate string
}

// SelectedEvent represents an event selected by the LLM
type SelectedEvent struct {
	Year        string `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

// SelectionResponse represents the LLM's response
type SelectionResponse struct {
	Events []SelectedEvent `json:"events"`
}

// NewSelector creates a new event selector
func NewSelector(apiKey, model string, maxEvents int, promptTemplate string) *Selector {
	return &Selector{
		apiKey:      apiKey,
		model:       model,
		maxEvents:   maxEvents,
		promptTemplate: promptTemplate,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// SelectEvents uses Claude API to select the most interesting events
func (s *Selector) SelectEvents(events []rss.HistoricalEvent) ([]SelectedEvent, error) {
	if len(events) == 0 {
		return nil, fmt.Errorf("no events to select from")
	}

	// Format events for the LLM prompt
	eventsText := s.formatEventsForPrompt(events)

	// Create the prompt
	prompt := fmt.Sprintf(s.promptTemplate, s.maxEvents)
	prompt += "\n\nHere are today's historical events:\n\n" + eventsText

	// Call Claude API
	response, err := s.callClaudeAPI(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to call Claude API: %w", err)
	}

	// Parse the response
	selected, err := s.parseSelection(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse selection: %w", err)
	}

	return selected, nil
}

// formatEventsForPrompt formats events into a readable text format
func (s *Selector) formatEventsForPrompt(events []rss.HistoricalEvent) string {
	var buf bytes.Buffer

	for i, event := range events {
		buf.WriteString(fmt.Sprintf("%d. ", i+1))
		if event.Year != "" {
			buf.WriteString(fmt.Sprintf("[%s] ", event.Year))
		}
		buf.WriteString(event.Title)
		if event.Description != "" {
			buf.WriteString("\n   ")
			buf.WriteString(event.Description)
		}
		if event.Category != "" {
			buf.WriteString(fmt.Sprintf("\n   Category: %s", event.Category))
		}
		buf.WriteString("\n\n")
	}

	return buf.String()
}

// ClaudeRequest represents the request structure for Claude API
type ClaudeRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

// Message represents a message in the Claude API
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ClaudeResponse represents the response from Claude API
type ClaudeResponse struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"`
	Role    string          `json:"role"`
	Content []ContentBlock  `json:"content"`
	Model   string          `json:"model"`
	Usage   UsageInfo       `json:"usage"`
}

// ContentBlock represents a content block in Claude's response
type ContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// UsageInfo represents token usage information
type UsageInfo struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// callClaudeAPI makes a request to the Claude API
func (s *Selector) callClaudeAPI(prompt string) (string, error) {
	request := ClaudeRequest{
		Model:     s.model,
		MaxTokens: 2048,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var claudeResp ClaudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(claudeResp.Content) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	return claudeResp.Content[0].Text, nil
}

// parseSelection parses the LLM's response into selected events
func (s *Selector) parseSelection(response string) ([]SelectedEvent, error) {
	// The LLM should return JSON, but it might be wrapped in markdown code blocks
	response = extractJSON(response)

	var selection SelectionResponse
	if err := json.Unmarshal([]byte(response), &selection); err != nil {
		return nil, fmt.Errorf("failed to unmarshal selection: %w (response: %s)", err, response)
	}

	if len(selection.Events) == 0 {
		return nil, fmt.Errorf("no events in selection")
	}

	return selection.Events, nil
}

// extractJSON extracts JSON from markdown code blocks if present
func extractJSON(s string) string {
	// Check if response is wrapped in ```json ... ```
	if bytes.Contains([]byte(s), []byte("```json")) {
		start := bytes.Index([]byte(s), []byte("```json"))
		if start != -1 {
			start += len("```json")
			end := bytes.Index([]byte(s[start:]), []byte("```"))
			if end != -1 {
				return s[start : start+end]
			}
		}
	}

	// Check if response is wrapped in ``` ... ```
	if bytes.Contains([]byte(s), []byte("```")) {
		start := bytes.Index([]byte(s), []byte("```"))
		if start != -1 {
			start += len("```")
			end := bytes.Index([]byte(s[start:]), []byte("```"))
			if end != -1 {
				return s[start : start+end]
			}
		}
	}

	return s
}
