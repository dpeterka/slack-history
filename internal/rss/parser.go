package rss

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Feed represents an RSS feed
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the RSS channel
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item represents an RSS item
type Item struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Categories  []string `xml:"category"`
	GUID        string   `xml:"guid"`
}

// HistoricalEvent represents a parsed historical event
type HistoricalEvent struct {
	Year        string
	Title       string
	Description string
	Category    string
	Link        string
	RawItem     Item
}

// Holiday represents a fun holiday
type Holiday struct {
	Title       string
	Description string
	Link        string
}

// Parser handles RSS feed parsing
type Parser struct {
	client *http.Client
}

// NewParser creates a new RSS parser
func NewParser() *Parser {
	return &Parser{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// FetchAndParse fetches and parses an RSS feed from the given URL
func (p *Parser) FetchAndParse(url string) ([]HistoricalEvent, error) {
	// Create request with browser headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add Chrome browser headers to avoid 403 errors
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,application/rss+xml,text/html;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")

	// Fetch the RSS feed
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Handle gzip-compressed response
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	// Read the response body
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the XML
	var feed Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	// Convert items to historical events
	events := make([]HistoricalEvent, 0, len(feed.Channel.Items))
	for _, item := range feed.Channel.Items {
		event := p.parseItem(item)
		events = append(events, event)
	}

	return events, nil
}

// parseItem parses an RSS item into a HistoricalEvent
func (p *Parser) parseItem(item Item) HistoricalEvent {
	event := HistoricalEvent{
		Title:       item.Title,
		Description: cleanHTML(item.Description),
		Link:        item.Link,
		RawItem:     item,
	}

	// Extract year from title if present (e.g., "1969: Apollo 11...")
	if parts := strings.SplitN(item.Title, ":", 2); len(parts) == 2 {
		event.Year = strings.TrimSpace(parts[0])
		event.Title = strings.TrimSpace(parts[1])
	}

	// Use first category if available
	if len(item.Categories) > 0 {
		event.Category = item.Categories[0]
	}

	return event
}

// cleanHTML strips basic HTML tags from a string
// For more robust HTML cleaning, consider using golang.org/x/net/html
func cleanHTML(s string) string {
	// Remove common HTML tags
	s = strings.ReplaceAll(s, "<br>", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	s = strings.ReplaceAll(s, "<p>", "\n")
	s = strings.ReplaceAll(s, "</p>", "\n")

	// Remove all remaining tags
	for {
		start := strings.Index(s, "<")
		if start == -1 {
			break
		}
		end := strings.Index(s[start:], ">")
		if end == -1 {
			break
		}
		s = s[:start] + s[start+end+1:]
	}

	// Clean up whitespace
	s = strings.TrimSpace(s)

	// Remove excessive newlines
	for strings.Contains(s, "\n\n\n") {
		s = strings.ReplaceAll(s, "\n\n\n", "\n\n")
	}

	return s
}

// FetchMultipleFeeds fetches and parses multiple RSS feeds
func (p *Parser) FetchMultipleFeeds(urls []string) ([]HistoricalEvent, error) {
	var allEvents []HistoricalEvent

	for _, url := range urls {
		events, err := p.FetchAndParse(url)
		if err != nil {
			// Log error but continue with other feeds
			fmt.Printf("Warning: failed to fetch feed %s: %v\n", url, err)
			continue
		}
		allEvents = append(allEvents, events...)
	}

	if len(allEvents) == 0 {
		return nil, fmt.Errorf("no events fetched from any feed")
	}

	return allEvents, nil
}

// FetchHolidays fetches holidays from a holiday RSS feed
func (p *Parser) FetchHolidays(url string) ([]Holiday, error) {
	// Create request with browser headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add Chrome browser headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,application/rss+xml,text/html;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")

	// Fetch the RSS feed
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Handle gzip-compressed response
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	// Read the response body
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the XML
	var feed Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	// Convert items to holidays
	holidays := make([]Holiday, 0, len(feed.Channel.Items))
	for _, item := range feed.Channel.Items {
		holiday := Holiday{
			Title:       item.Title,
			Description: cleanHTML(item.Description),
			Link:        item.Link,
		}
		holidays = append(holidays, holiday)
	}

	return holidays, nil
}
