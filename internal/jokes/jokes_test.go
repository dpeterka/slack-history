package jokes

import (
	"strings"
	"testing"
)

func TestGetRandomJoke(t *testing.T) {
	joke := GetRandomJoke()

	if joke.Text == "" {
		t.Error("Expected non-empty joke text")
	}
}

func TestGetRandomJokeConsistency(t *testing.T) {
	joke1 := GetRandomJoke()
	joke2 := GetRandomJoke()

	if joke1.Text != joke2.Text {
		t.Error("Expected same joke for the same day")
	}
}

func TestGetRandomJokeWithSeed(t *testing.T) {
	seed := 20251225
	joke := GetRandomJokeWithSeed(seed)

	if joke.Text == "" {
		t.Error("Expected non-empty joke text")
	}

	joke2 := GetRandomJokeWithSeed(seed)
	if joke.Text != joke2.Text {
		t.Error("Expected same joke for same seed")
	}
}

func TestAllJokesHaveContent(t *testing.T) {
	jokes := getAllJokes()

	for i, joke := range jokes {
		if joke.Text == "" {
			t.Errorf("Joke %d has empty text", i)
		}
	}
}

func TestMinimumNumberOfJokes(t *testing.T) {
	jokes := getAllJokes()

	if len(jokes) < 30 {
		t.Errorf("Expected at least 30 jokes, got %d", len(jokes))
	}
}

// Defense in depth: the rendered joke must never contain the literal phrase
// "dad joke" (case-insensitive). The product requirement is that the joke
// renders without any "dad joke" framing.
func TestNoDadJokeLabel(t *testing.T) {
	jokes := getAllJokes()

	for i, joke := range jokes {
		if strings.Contains(strings.ToLower(joke.Text), "dad joke") {
			t.Errorf("Joke %d contains forbidden substring 'dad joke': %q", i, joke.Text)
		}
	}
}
