package rss

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/andybalholm/brotli"
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

// sanitizeXML fixes common XML issues in RSS feeds
func sanitizeXML(body []byte) []byte {
	// First, check for and remove BOM (Byte Order Mark)
	if len(body) >= 3 && body[0] == 0xEF && body[1] == 0xBB && body[2] == 0xBF {
		body = body[3:]
	}

	// Remove null bytes which are illegal in XML
	body = bytes.ReplaceAll(body, []byte{0}, []byte{})

	// Remove other control characters that are illegal in XML (except tab, newline, carriage return)
	cleaned := make([]byte, 0, len(body))
	for _, b := range body {
		// Keep printable characters, tab (9), newline (10), carriage return (13)
		if b >= 32 || b == 9 || b == 10 || b == 13 {
			cleaned = append(cleaned, b)
		}
	}
	body = cleaned

	// Remove invalid UTF-8 sequences
	if !utf8.Valid(body) {
		v := make([]byte, 0, len(body))
		for i := 0; i < len(body); {
			r, size := utf8.DecodeRune(body[i:])
			if r == utf8.RuneError && size == 1 {
				i++
				continue
			}
			v = append(v, body[i:i+size]...)
			i += size
		}
		body = v
	}

	// Look for <?xml declaration and ensure it's at the very start
	// Remove any whitespace before the XML declaration
	xmlDeclStart := bytes.Index(body, []byte("<?xml"))
	if xmlDeclStart > 0 {
		body = body[xmlDeclStart:]
	}

	// Fix common HTML entity issues that aren't properly encoded in XML
	// Convert problematic HTML entities to their actual characters
	entityReplacements := map[string]string{
		"&rsquo;":  "'",
		"&lsquo;":  "'",
		"&rdquo;":  "\"",
		"&ldquo;":  "\"",
		"&hellip;": "...",
		"&ndash;":  "-",
		"&mdash;":  "—",
		"&nbsp;":   " ",
		"&apos;":   "'",
		"&raquo;":  ">>",
		"&laquo;":  "<<",
	}

	for old, new := range entityReplacements {
		body = bytes.ReplaceAll(body, []byte(old), []byte(new))
	}

	// Fix standalone ampersands that aren't part of entities
	// This is a simple approach: replace & followed by space or certain punctuation
	body = bytes.ReplaceAll(body, []byte("& "), []byte("&amp; "))
	body = bytes.ReplaceAll(body, []byte("&\n"), []byte("&amp;\n"))
	body = bytes.ReplaceAll(body, []byte("&\""), []byte("&amp;\""))
	body = bytes.ReplaceAll(body, []byte("&'"), []byte("&amp;'"))

	return body
}

// FetchAndParse fetches and parses an RSS feed from the given URL
func (p *Parser) FetchAndParse(url string) ([]HistoricalEvent, error) {
	// Create request with browser headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add Chrome browser headers to avoid 403 errors
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,application/rss+xml,text/html;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Referer", "https://www.google.com/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// Fetch the RSS feed
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Handle compressed response (gzip or brotli)
	var reader io.Reader = resp.Body
	contentEncoding := resp.Header.Get("Content-Encoding")
	switch contentEncoding {
	case "gzip":
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	case "br":
		reader = brotli.NewReader(resp.Body)
	}

	// Read the response body
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Sanitize XML to fix common encoding and entity issues
	body = sanitizeXML(body)

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
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,application/rss+xml,text/html;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Referer", "https://www.google.com/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// Fetch the RSS feed
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Handle compressed response (gzip or brotli)
	var reader io.Reader = resp.Body
	contentEncoding := resp.Header.Get("Content-Encoding")
	switch contentEncoding {
	case "gzip":
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	case "br":
		reader = brotli.NewReader(resp.Body)
	}

	// Read the response body
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Sanitize XML to fix common encoding and entity issues
	body = sanitizeXML(body)

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
