package llm

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Generator writes a brand-new content item in the style of an
// existing static pack, so daily content never repeats.
type Generator struct {
	apiKey string
	model  string
	client *http.Client
}

// NewGenerator creates a content generator
func NewGenerator(apiKey, model string) *Generator {
	return &Generator{
		apiKey: apiKey,
		model:  model,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// GenerateFreshItem produces one new item matching the pack's style.
// packDescription says what the pack is; styleExamples calibrate tone and length.
// Returns the item text, or an error (callers should fall back to the static pack).
func (g *Generator) GenerateFreshItem(packDescription string, styleExamples []string) (string, error) {
	var examples strings.Builder
	for i, ex := range styleExamples {
		fmt.Fprintf(&examples, "%d. %s\n", i+1, ex)
	}

	prompt := fmt.Sprintf(`You write daily content for a lighthearted workplace Slack bot.

Today's content type: %s

Here are examples of past items, to calibrate tone, length, and style:
%s
Write ONE brand-new item of this type. Requirements:
- Match the tone, voice, and approximate length of the examples
- It must be genuinely new — not a rewording of any example
- Keep it workplace-friendly (irreverent is fine, offensive is not)
- Respond with ONLY the item text: no numbering, no quotes around it, no preamble, no explanation`, packDescription, examples.String())

	// Reuse the same API plumbing as the event selector.
	s := &Selector{apiKey: g.apiKey, model: g.model, client: g.client}
	response, err := s.callClaudeAPI(prompt)
	if err != nil {
		return "", fmt.Errorf("generation failed: %w", err)
	}

	item := strings.TrimSpace(response)
	item = strings.Trim(item, `"`)
	if item == "" {
		return "", fmt.Errorf("generation returned empty content")
	}
	return item, nil
}
