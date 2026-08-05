package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dpeterka/history-slackbot/internal/birthday"
	"github.com/dpeterka/history-slackbot/internal/funfacts"
	"github.com/dpeterka/history-slackbot/internal/llm"
	"github.com/dpeterka/history-slackbot/internal/wikipedia"
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
	Type     string       `json:"type"`
	Text     *TextObject  `json:"text,omitempty"`
	Elements []TextObject `json:"elements,omitempty"`
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

// PostComplete posts complete message with all content types
func (p *Poster) PostComplete(events []llm.SelectedEvent, majorHoliday string, notablePeople []wikipedia.Person, funFact *funfacts.FunFact) error {
	if len(events) == 0 && majorHoliday == "" && len(notablePeople) == 0 && funFact == nil {
		return fmt.Errorf("no content to post")
	}

	// Special case: if funFact is "events" type but we have no events, we can't post
	if funFact != nil && funFact.Type == "events" && len(events) == 0 {
		log.Printf("Skipping post: content type is 'events' but no events available")
		return nil
	}

	message := p.formatCompleteMessage(events, majorHoliday, notablePeople, funFact)

	reqBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Log the payload size and webhook URL (partial for security)
	webhookPreview := p.webhookURL
	if len(webhookPreview) > 50 {
		webhookPreview = webhookPreview[:50] + "..."
	}
	log.Printf("Posting to Slack webhook: %s, payload size: %d bytes", webhookPreview, len(reqBody))

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

	// Log Slack response for debugging
	responseText := string(body)
	log.Printf("Slack API response: status=%d, body='%s'", resp.StatusCode, responseText)

	if resp.StatusCode != http.StatusOK {
		log.Printf("ERROR: Slack webhook rejected the message with status %d", resp.StatusCode)
		return fmt.Errorf("Slack API request failed with status %d: %s", resp.StatusCode, responseText)
	}

	// Check if Slack returned an error in the body
	if responseText != "ok" {
		log.Printf("WARNING: Slack returned unexpected response (not 'ok'): '%s'", responseText)
		log.Printf("This may indicate the message was not posted successfully")
	} else {
		log.Printf("SUCCESS: Slack webhook accepted the message (returned 'ok')")
	}

	return nil
}

// PostMajorHoliday posts a simple major holiday announcement
func (p *Poster) PostMajorHoliday(holiday string) error {
	now := time.Now()
	dateStr := now.Format("Monday, January 2")

	blocks := []Block{
		{
			Type: "header",
			Text: &TextObject{
				Type: "plain_text",
				Text: fmt.Sprintf("From the desk of the Grant - %s", dateStr),
			},
		},
		{
			Type: "divider",
		},
		{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*Today is %s* 🎉", holiday),
			},
		},
	}

	message := SlackMessage{
		Blocks: blocks,
	}

	reqBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	webhookPreview := p.webhookURL
	if len(webhookPreview) > 50 {
		webhookPreview = webhookPreview[:50] + "..."
	}
	log.Printf("Posting major holiday to Slack webhook: %s, payload size: %d bytes", webhookPreview, len(reqBody))

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

	responseText := string(body)
	log.Printf("Slack API response for major holiday: status=%d, body='%s'", resp.StatusCode, responseText)

	if resp.StatusCode != http.StatusOK {
		log.Printf("ERROR: Slack webhook rejected the major holiday message with status %d", resp.StatusCode)
		return fmt.Errorf("Slack API request failed with status %d: %s", resp.StatusCode, responseText)
	}

	if responseText != "ok" {
		log.Printf("WARNING: Slack returned unexpected response (not 'ok'): '%s'", responseText)
	} else {
		log.Printf("SUCCESS: Slack webhook accepted the major holiday message (returned 'ok')")
	}

	return nil
}

// formatMessage formats events into a Slack message with blocks
// formatCompleteMessage formats all content types into a Slack message
func (p *Poster) formatCompleteMessage(events []llm.SelectedEvent, majorHoliday string, notablePeople []wikipedia.Person, funFact *funfacts.FunFact) SlackMessage {
	now := time.Now()
	dateStr := now.Format("Monday, January 2")

	// Create header block
	blocks := []Block{
		{
			Type: "header",
			Text: &TextObject{
				Type: "plain_text",
				Text: fmt.Sprintf("From the desk of the Grant - %s", dateStr),
			},
		},
		{
			Type: "divider",
		},
	}

	// Add fun fact section if present (emo, blobby, wikihow, quote, holidays, people)
	if funFact != nil && funFact.Type != "events" && funFact.Type != "people" && funFact.Type != "joke" {
		// Add title for the fact type (skip for events and people as they have their own section titles;
		// jokes render bare without any title for non-sequitur impact)
		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*%s*", funFact.GetDisplayTitle()),
			},
		})
	}

	// Handle different fact types
	if funFact != nil {
		switch funFact.Type {
		case "events":
			// Events are rendered in their own section below; emitting the
			// funFact here would produce an empty section block, which Slack
			// rejects with invalid_blocks
		case "quote":
			// Format quote with author
			quoteText := fmt.Sprintf("_%s_\n— %s", funFact.Text, funFact.Author)
			blocks = append(blocks, Block{
				Type: "section",
				Text: &TextObject{
					Type: "mrkdwn",
					Text: quoteText,
				},
			})
		case "holidays":
			// Format holidays as bullet list with links
			holidayText := ""
			for i, holiday := range funFact.Holidays {
				if i > 0 {
					holidayText += "\n"
				}
				// Add link if available
				if holiday.Link != "" {
					holidayText += fmt.Sprintf("• <%s|%s>", holiday.Link, holiday.Title)
				} else {
					holidayText += fmt.Sprintf("• %s", holiday.Title)
				}
			}
			blocks = append(blocks, Block{
				Type: "section",
				Text: &TextObject{
					Type: "mrkdwn",
					Text: holidayText,
				},
			})
		case "people":
			// Format notable people (births and deaths)
			// Separate births and deaths
			var births, deaths []wikipedia.Person
			for _, person := range funFact.NotablePeople {
				if person.Type == "birth" {
					births = append(births, person)
				} else if person.Type == "death" {
					deaths = append(deaths, person)
				}
			}

			if len(births) > 0 {
				blocks = append(blocks, Block{
					Type: "section",
					Text: &TextObject{
						Type: "mrkdwn",
						Text: "*🎂 Born on This Day*",
					},
				})

				for _, person := range births {
					// Format name with Wikipedia link if available
					var nameText string
					if person.WikiURL != "" {
						// Extract just the person's name (before comma if present)
						displayName := person.Name
						commaIndex := strings.Index(person.Name, ",")
						if commaIndex > 0 {
							displayName = person.Name[:commaIndex]
						}
						// Slack markdown link format: <URL|Link Text>
						nameText = fmt.Sprintf("<%s|*%s*> (%d)", person.WikiURL, displayName, person.Year)
					} else {
						nameText = fmt.Sprintf("*%s* (%d)", person.Name, person.Year)
					}

					personText := nameText
					if person.Description != "" {
						personText += fmt.Sprintf("\n%s", person.Description)
					}
					blocks = append(blocks, Block{
						Type: "section",
						Text: &TextObject{
							Type: "mrkdwn",
							Text: personText,
						},
					})
				}
			}

			if len(deaths) > 0 {
				blocks = append(blocks, Block{
					Type: "section",
					Text: &TextObject{
						Type: "mrkdwn",
						Text: "*🕊️ Died on This Day*",
					},
				})

				for _, person := range deaths {
					// Format name with Wikipedia link if available
					var nameText string
					if person.WikiURL != "" {
						// Extract just the person's name (before comma if present)
						displayName := person.Name
						commaIndex := strings.Index(person.Name, ",")
						if commaIndex > 0 {
							displayName = person.Name[:commaIndex]
						}
						// Slack markdown link format: <URL|Link Text>
						nameText = fmt.Sprintf("<%s|*%s*> (%d)", person.WikiURL, displayName, person.Year)
					} else {
						nameText = fmt.Sprintf("*%s* (%d)", person.Name, person.Year)
					}

					personText := nameText
					if person.Description != "" {
						personText += fmt.Sprintf("\n%s", person.Description)
					}
					blocks = append(blocks, Block{
						Type: "section",
						Text: &TextObject{
							Type: "mrkdwn",
							Text: personText,
						},
					})
				}
			}
		default:
			// Standard text formatting for emo, blobby, wikihow
			if funFact.Text == "" {
				// A section block with empty text is rejected by Slack
				break
			}
			factText := funFact.Text
			if funFact.ShouldDisplayAsItalic() {
				factText = fmt.Sprintf("_%s_", factText)
			}

			// Add URL link for WikiHow articles
			if funFact.URL != "" {
				factText = fmt.Sprintf("<%s|%s>", funFact.URL, factText)
			}

			blocks = append(blocks, Block{
				Type: "section",
				Text: &TextObject{
					Type: "mrkdwn",
					Text: factText,
				},
			})
		}

		// Add divider after fun fact content (but not for events, as they have their own formatting)
		if funFact.Type != "events" {
			blocks = append(blocks, Block{
				Type: "divider",
			})
		}
	}

	// Quote and holidays are now handled via funFact above
	// Major holidays are posted as a separate message

	// Add notable people section if present
	if len(notablePeople) > 0 {
		// Separate births and deaths
		var births, deaths []wikipedia.Person
		for _, person := range notablePeople {
			if person.Type == "birth" {
				births = append(births, person)
			} else if person.Type == "death" {
				deaths = append(deaths, person)
			}
		}

		if len(births) > 0 {
			blocks = append(blocks, Block{
				Type: "section",
				Text: &TextObject{
					Type: "mrkdwn",
					Text: "*🎂 Born on This Day*",
				},
			})

			for _, person := range births {
				// Format name with Wikipedia link if available
				var nameText string
				if person.WikiURL != "" {
					// Slack markdown link format: <URL|Link Text>
					nameText = fmt.Sprintf("<%s|*%s*> (%d)", person.WikiURL, person.Name, person.Year)
				} else {
					nameText = fmt.Sprintf("*%s* (%d)", person.Name, person.Year)
				}

				personText := nameText
				if person.Description != "" {
					personText += fmt.Sprintf("\n%s", person.Description)
				}
				blocks = append(blocks, Block{
					Type: "section",
					Text: &TextObject{
						Type: "mrkdwn",
						Text: personText,
					},
				})
			}

			blocks = append(blocks, Block{
				Type: "divider",
			})
		}

		if len(deaths) > 0 {
			blocks = append(blocks, Block{
				Type: "section",
				Text: &TextObject{
					Type: "mrkdwn",
					Text: "*🕊️ Died on This Day*",
				},
			})

			for _, person := range deaths {
				// Format name with Wikipedia link if available
				var nameText string
				if person.WikiURL != "" {
					// Slack markdown link format: <URL|Link Text>
					nameText = fmt.Sprintf("<%s|*%s*> (%d)", person.WikiURL, person.Name, person.Year)
				} else {
					nameText = fmt.Sprintf("*%s* (%d)", person.Name, person.Year)
				}

				personText := nameText
				if person.Description != "" {
					personText += fmt.Sprintf("\n%s", person.Description)
				}
				blocks = append(blocks, Block{
					Type: "section",
					Text: &TextObject{
						Type: "mrkdwn",
						Text: personText,
					},
				})
			}

			blocks = append(blocks, Block{
				Type: "divider",
			})
		}
	}

	// Add each event as a section
	for i, event := range events {
		// Event header with year only
		header := fmt.Sprintf("*%s*", event.Year)

		// Event title - add Wikipedia link if available
		var titleText string
		if event.WikiURL != "" {
			titleText = fmt.Sprintf("<%s|*%s*>", event.WikiURL, event.Title)
		} else {
			titleText = fmt.Sprintf("*%s*", event.Title)
		}

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

// PostBirthday posts a special birthday message for the bot
func (p *Poster) PostBirthday(birthdayMsg *birthday.BirthdayMessage) error {
	now := time.Now()
	dateStr := now.Format("Monday, January 2")

	blocks := []Block{
		{
			Type: "header",
			Text: &TextObject{
				Type: "plain_text",
				Text: fmt.Sprintf("From the desk of the Grant - %s", dateStr),
			},
		},
		{
			Type: "divider",
		},
		{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: birthdayMsg.Message,
			},
		},
	}

	// Add Giphy GIF if available
	if birthdayMsg.GiphyURL != "" {
		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("🎁 *Birthday GIF:* <%s|Click here for celebration!>", birthdayMsg.GiphyURL),
			},
		})
	}

	// Add YouTube video
	if birthdayMsg.YouTubeURL != "" {
		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("🎵 *Birthday Jam:* <%s|Watch the celebration video!>", birthdayMsg.YouTubeURL),
			},
		})
	}

	// Add a final celebratory note
	blocks = append(blocks, Block{
		Type: "divider",
	})
	blocks = append(blocks, Block{
		Type: "section",
		Text: &TextObject{
			Type: "mrkdwn",
			Text: "_Thank you for being part of my journey. Here's to many more years of historical facts, philosophical musings, and the occasional Mr Blobby reference!_ 🎊",
		},
	})

	message := SlackMessage{
		Blocks: blocks,
	}

	reqBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal birthday message: %w", err)
	}

	webhookPreview := p.webhookURL
	if len(webhookPreview) > 50 {
		webhookPreview = webhookPreview[:50] + "..."
	}
	log.Printf("Posting birthday message to Slack webhook: %s, payload size: %d bytes", webhookPreview, len(reqBody))

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

	responseText := string(body)
	log.Printf("Slack API response for birthday: status=%d, body='%s'", resp.StatusCode, responseText)

	if resp.StatusCode != http.StatusOK {
		log.Printf("ERROR: Slack webhook rejected the birthday message with status %d", resp.StatusCode)
		return fmt.Errorf("Slack API request failed with status %d: %s", resp.StatusCode, responseText)
	}

	if responseText != "ok" {
		log.Printf("WARNING: Slack returned unexpected response (not 'ok'): '%s'", responseText)
	} else {
		log.Printf("SUCCESS: Slack webhook accepted the birthday message (returned 'ok')")
	}

	return nil
}
