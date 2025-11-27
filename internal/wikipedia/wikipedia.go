package wikipedia

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Person represents a notable person from Wikipedia's "On This Day"
type Person struct {
	Name        string
	Year        int
	Description string
	Type        string // "birth" or "death"
	WikiURL     string
}

// WikiResponse represents the Wikipedia API response
type WikiResponse struct {
	Births []WikiPerson `json:"births"`
	Deaths []WikiPerson `json:"deaths"`
}

// WikiPerson represents a person from the API
type WikiPerson struct {
	Text  string       `json:"text"`
	Year  int          `json:"year"`
	Pages []WikiPage   `json:"pages"`
}

// WikiPage represents Wikipedia page details
type WikiPage struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ContentURLs ContentURLs `json:"content_urls"`
}

// ContentURLs contains links to the Wikipedia article
type ContentURLs struct {
	Desktop DesktopURL `json:"desktop"`
}

// DesktopURL contains the desktop page URL
type DesktopURL struct {
	Page string `json:"page"`
}

// FetchBirthsAndDeaths fetches notable births and deaths for a specific date
func FetchBirthsAndDeaths(month, day int, limit int) ([]Person, error) {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/all/%02d/%02d", month, day)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Wikipedia requires a User-Agent header
	req.Header.Set("User-Agent", "HistorySlackbot/1.0 (https://github.com/dpeterka/history-slackbot)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from Wikipedia API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Wikipedia API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var wikiResp WikiResponse
	if err := json.Unmarshal(body, &wikiResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	var people []Person

	// Split limit evenly between births and deaths
	birthLimit := limit / 2
	deathLimit := limit - birthLimit // Handle odd numbers by giving extra to deaths

	// Process births
	for i, birth := range wikiResp.Births {
		if i >= birthLimit {
			break
		}
		person := convertToPerson(birth, "birth")
		people = append(people, person)
	}

	// Process deaths
	for i, death := range wikiResp.Deaths {
		if i >= deathLimit {
			break
		}
		person := convertToPerson(death, "death")
		people = append(people, person)
	}

	return people, nil
}

// FetchBirths fetches only notable births for a specific date
func FetchBirths(month, day int, limit int) ([]Person, error) {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/births/%02d/%02d", month, day)
	return fetchFromEndpoint(url, "birth", limit)
}

// FetchDeaths fetches only notable deaths for a specific date
func FetchDeaths(month, day int, limit int) ([]Person, error) {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/deaths/%02d/%02d", month, day)
	return fetchFromEndpoint(url, "death", limit)
}

func fetchFromEndpoint(url string, personType string, limit int) ([]Person, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Wikipedia requires a User-Agent header
	req.Header.Set("User-Agent", "HistorySlackbot/1.0 (https://github.com/dpeterka/history-slackbot)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from Wikipedia API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Wikipedia API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// The births/deaths endpoints return data wrapped in an object
	var result map[string][]WikiPerson
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Get the array from either "births" or "deaths" key
	var wikiPeople []WikiPerson
	if personType == "birth" {
		wikiPeople = result["births"]
	} else {
		wikiPeople = result["deaths"]
	}

	var people []Person
	for i, wp := range wikiPeople {
		if i >= limit {
			break
		}
		person := convertToPerson(wp, personType)
		people = append(people, person)
	}

	return people, nil
}

func convertToPerson(wp WikiPerson, personType string) Person {
	person := Person{
		Name: wp.Text,
		Year: wp.Year,
		Type: personType,
	}

	// Extract description and URL from first page if available
	if len(wp.Pages) > 0 {
		page := wp.Pages[0]
		person.Description = page.Description
		person.WikiURL = page.ContentURLs.Desktop.Page
	}

	return person
}

// FetchTodaysBirthsAndDeaths is a convenience function for today's date
func FetchTodaysBirthsAndDeaths(limit int) ([]Person, error) {
	now := time.Now()
	return FetchBirthsAndDeaths(int(now.Month()), now.Day(), limit)
}

// SearchWikipediaURL searches Wikipedia for a term and returns the URL of the first result
func SearchWikipediaURL(searchTerm string) (string, error) {
	// Wikipedia opensearch API with URL encoding
	encodedTerm := url.QueryEscape(searchTerm)
	apiURL := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&search=%s&limit=1&format=json",
		encodedTerm)

	// Create request with User-Agent header (required by Wikipedia)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "HistorySlackBot/1.0 (https://github.com/dpeterka/history-slackbot)")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch from Wikipedia: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Wikipedia opensearch returns: [search term, [titles], [descriptions], [urls]]
	var results []interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Check if we have results
	if len(results) < 4 {
		return "", fmt.Errorf("unexpected response format")
	}

	// Extract URLs array (4th element)
	urls, ok := results[3].([]interface{})
	if !ok || len(urls) == 0 {
		return "", fmt.Errorf("no results found")
	}

	// Get first URL
	resultURL, ok := urls[0].(string)
	if !ok {
		return "", fmt.Errorf("invalid URL format")
	}

	return resultURL, nil
}
