package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dpeterka/history-slackbot/internal/funfacts"
	"github.com/dpeterka/history-slackbot/internal/llm"
	"github.com/dpeterka/history-slackbot/internal/rss"
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

// PostEvents posts selected events to Slack
func (p *Poster) PostEvents(events []llm.SelectedEvent) error {
	return p.PostEventsWithHolidays(events, nil)
}

// PostEventsWithHolidays posts selected events and holidays to Slack
func (p *Poster) PostEventsWithHolidays(events []llm.SelectedEvent, holidays []rss.Holiday) error {
	return p.PostEventsWithHolidaysAndMajor(events, holidays, "")
}

// PostEventsWithHolidaysAndMajor posts selected events, holidays, and major holiday to Slack
func (p *Poster) PostEventsWithHolidaysAndMajor(events []llm.SelectedEvent, holidays []rss.Holiday, majorHoliday string) error {
	// Create a funFact for holidays if provided
	var funFact *funfacts.FunFact
	if len(holidays) > 0 {
		funFact = &funfacts.FunFact{
			Type:     "holidays",
			Holidays: holidays,
		}
	}
	return p.PostComplete(events, majorHoliday, nil, funFact)
}

// PostComplete posts complete message with all content types
func (p *Poster) PostComplete(events []llm.SelectedEvent, majorHoliday string, notablePeople []wikipedia.Person, funFact *funfacts.FunFact) error {
	if len(events) == 0 && majorHoliday == "" && len(notablePeople) == 0 && funFact == nil {
		return fmt.Errorf("no content to post")
	}

	message := p.formatCompleteMessage(events, majorHoliday, notablePeople, funFact)

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
func (p *Poster) formatMessageWithHolidays(events []llm.SelectedEvent, holidays []rss.Holiday) SlackMessage {
	return p.formatMessageWithHolidaysAndMajor(events, holidays, "")
}

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
	if funFact != nil && funFact.Type != "events" && funFact.Type != "people" {
		// Add title for the fact type (skip for events and people as they have their own section titles)
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

	// Add major holiday section if present
	if majorHoliday != "" {
		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*Today is %s*", majorHoliday),
			},
		})

		blocks = append(blocks, Block{
			Type: "divider",
		})
	}

	// Quote and holidays are now handled via funFact above

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

// formatMessageWithHolidaysAndMajor formats events, holidays, and major holiday into a Slack message with blocks
func (p *Poster) formatMessageWithHolidaysAndMajor(events []llm.SelectedEvent, holidays []rss.Holiday, majorHoliday string) SlackMessage {
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

	// Add major holiday section if present
	if majorHoliday != "" {
		blocks = append(blocks, Block{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*Today is %s*", majorHoliday),
			},
		})

		blocks = append(blocks, Block{
			Type: "divider",
		})
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

		// Add each holiday with link
		holidayText := ""
		for i, holiday := range holidays {
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

		blocks = append(blocks, Block{
			Type: "divider",
		})
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
