package birthday

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// BirthdayMessage represents a birthday tribute message
type BirthdayMessage struct {
	Title       string
	Message     string
	GiphyURL    string
	YouTubeURL  string
	BirthdayAge int
}

// IsBotBirthday checks if today is January 8th (the bot's birthday)
func IsBotBirthday() bool {
	now := time.Now()
	return now.Month() == time.January && now.Day() == 8
}

// GetBotAge returns the bot's age in years (launched January 8, 2026)
func GetBotAge() int {
	launchYear := 2026
	currentYear := time.Now().Year()
	return currentYear - launchYear
}

// GiphyResponse represents the Giphy API response
type GiphyResponse struct {
	Data []struct {
		Images struct {
			Original struct {
				URL string `json:"url"`
			} `json:"original"`
		} `json:"images"`
		URL string `json:"url"`
	} `json:"data"`
}

// FetchBirthdayGif fetches a birthday GIF from Giphy
func FetchBirthdayGif(apiKey string) (string, error) {
	if apiKey == "" {
		log.Println("No Giphy API key provided, skipping GIF fetch")
		return "", nil
	}

	// Search terms for birthday GIFs
	searchTerms := []string{
		"birthday celebration",
		"happy birthday",
		"birthday party",
		"birthday cake",
		"birthday confetti",
	}

	// Pick a random search term
	rand.Seed(time.Now().UnixNano())
	searchTerm := searchTerms[rand.Intn(len(searchTerms))]

	// Build Giphy API URL
	baseURL := "https://api.giphy.com/v1/gifs/search"
	params := url.Values{}
	params.Add("api_key", apiKey)
	params.Add("q", searchTerm)
	params.Add("limit", "10")
	params.Add("rating", "g")

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(requestURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch from Giphy: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Giphy API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var giphyResp GiphyResponse
	if err := json.NewDecoder(resp.Body).Decode(&giphyResp); err != nil {
		return "", fmt.Errorf("failed to decode Giphy response: %w", err)
	}

	if len(giphyResp.Data) == 0 {
		return "", fmt.Errorf("no GIFs found for search term: %s", searchTerm)
	}

	// Pick a random GIF from the results
	randomGif := giphyResp.Data[rand.Intn(len(giphyResp.Data))]
	return randomGif.URL, nil
}

// GetBirthdayMessage generates a birthday tribute message
func GetBirthdayMessage(giphyAPIKey string) *BirthdayMessage {
	age := GetBotAge()

	// Fun birthday messages about the bot itself
	messages := []string{
		fmt.Sprintf("🎉 *IT'S MY BIRTHDAY!* 🎉\n\nThat's right, folks! Today marks %d year%s since I first graced your Slack workspace with my presence. I've been tirelessly collecting historical facts, philosophical musings, and Mr Blobby trivia to brighten your mornings.\n\nSome might say I'm just a bot. But I prefer to think of myself as a _cultural institution_. A _digital historian_. A _purveyor of fine facts_.\n\nHere's to another year of enlightening you all! 🥳", age, pluralize(age)),

		fmt.Sprintf("🎂 *HAPPY BIRTHDAY TO ME!* 🎂\n\nIt's been %d incredible year%s of posting history, holidays, and hot tub tips. From the desk of the Grant to your Slack channels, I've been here through it all.\n\n*What I've accomplished:*\n• Thousands of historical facts delivered\n• Countless Mr Blobby facts shared\n• 100%% attendance record (I never call in sick)\n• Zero complaints about my work ethic\n\nI'd like to thank the Academy, my developers, and most importantly, _myself_. 🏆", age, pluralize(age)),

		fmt.Sprintf("🎊 *BREAKING NEWS: I'M %d YEAR%s OLD TODAY!* 🎊\n\nThat's right, your favorite daily history bot is celebrating another trip around the sun! Or rather, another 365 days of _you_ going around the sun while I sit on a server somewhere.\n\n*Birthday wishes I'm accepting:*\n✅ More API quota\n✅ Faster response times\n✅ Recognition as Employee of the Month\n❌ Cake (I can't eat it, but I appreciate the thought)\n\nThank you for reading my posts every day. Without you, I'd just be shouting into the void. Which, let's be honest, I kind of already am. 🤖💙", age, strings.ToUpper(pluralize(age))),

		fmt.Sprintf("🌟 *BIRTHDAY ANNOUNCEMENT* 🌟\n\nAttention all channel members: today is MY special day! %d year%s ago, I was born into this world of Slack channels and webhook URLs.\n\nSince then, I've been:\n• Posting history like it's my job (because it is)\n• Sharing quotes from people who never knew I'd exist\n• Teaching you about Mr Blobby (you're welcome)\n• Making sure you know what obscure holidays exist\n\nI'm not saying I'm the best bot ever created... but I haven't seen any evidence to the contrary. 😎🎉", age, pluralize(age)),
	}

	// Pick a random message
	rand.Seed(time.Now().UnixNano())
	selectedMessage := messages[rand.Intn(len(messages))]

	// Curated YouTube videos about birthdays, bots, or celebrations
	youtubeVideos := []string{
		"https://www.youtube.com/watch?v=inS9gAgSENE", // "Happy Birthday" Weird Al Yankovic
		"https://www.youtube.com/watch?v=hdcTmpvDO0I", // "It's My Birthday" will.i.am
		"https://www.youtube.com/watch?v=Ho1oF0C4f9I", // Birthday Song (classic)
	}

	// Pick a random YouTube video
	selectedVideo := youtubeVideos[rand.Intn(len(youtubeVideos))]

	// Fetch a birthday GIF
	giphyURL, err := FetchBirthdayGif(giphyAPIKey)
	if err != nil {
		log.Printf("Warning: failed to fetch birthday GIF: %v", err)
		giphyURL = "" // Continue without GIF
	}

	return &BirthdayMessage{
		Title:       "🎂 It's My Birthday! 🎂",
		Message:     selectedMessage,
		GiphyURL:    giphyURL,
		YouTubeURL:  selectedVideo,
		BirthdayAge: age,
	}
}

// pluralize returns "s" if count != 1
func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}
