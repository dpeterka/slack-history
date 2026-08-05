package slack

import (
	"encoding/json"
	"testing"

	"github.com/dpeterka/history-slackbot/internal/funfacts"
	"github.com/dpeterka/history-slackbot/internal/llm"
)

// assertNoEmptyTextBlocks marshals the message the way PostComplete does and
// fails if any block has a text object with empty text — Slack rejects those
// payloads with 400 invalid_blocks.
func assertNoEmptyTextBlocks(t *testing.T, msg SlackMessage) {
	t.Helper()
	raw, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("failed to marshal message: %v", err)
	}
	var decoded struct {
		Blocks []struct {
			Type string `json:"type"`
			Text *struct {
				Text string `json:"text"`
			} `json:"text"`
		} `json:"blocks"`
	}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("failed to unmarshal message: %v", err)
	}
	if len(decoded.Blocks) == 0 {
		t.Fatal("expected at least one block")
	}
	for i, b := range decoded.Blocks {
		if b.Type == "divider" {
			continue
		}
		if b.Text == nil || b.Text.Text == "" {
			t.Errorf("block %d (%s) has empty text — Slack rejects this as invalid_blocks", i, b.Type)
		}
	}
}

// Regression test for the 2026-08-05 failure: an "events" day funFact has no
// Text, and rendering it through the default case produced an empty section.
func TestFormatCompleteMessageEventsDay(t *testing.T) {
	p := NewPoster("https://example.invalid/webhook")
	events := []llm.SelectedEvent{
		{
			Year:        "1976",
			Title:       "Stevie Wonder Signs Record-Breaking $13 Million Contract with Motown",
			Description: "The deal was the largest in recording history at the time.",
			// No WikiURL — lookup failed on the real run too
		},
	}
	funFact := &funfacts.FunFact{Type: "events", ShowEvents: true}

	assertNoEmptyTextBlocks(t, p.formatCompleteMessage(events, "", nil, funFact))
}

func TestFormatCompleteMessageTextPacks(t *testing.T) {
	p := NewPoster("https://example.invalid/webhook")
	for _, typ := range []string{"emo", "blobby", "camping", "joke", "foodtakes", "hottub", "gardening", "printing3d"} {
		t.Run(typ, func(t *testing.T) {
			funFact := &funfacts.FunFact{Type: typ, Text: "Some daily content.", Category: "General"}
			assertNoEmptyTextBlocks(t, p.formatCompleteMessage(nil, "", nil, funFact))
		})
	}
}

func TestFormatCompleteMessageEmptyTextDoesNotPanic(t *testing.T) {
	p := NewPoster("https://example.invalid/webhook")
	funFact := &funfacts.FunFact{Type: "emo", Text: ""}
	assertNoEmptyTextBlocks(t, p.formatCompleteMessage(nil, "", nil, funFact))
}
