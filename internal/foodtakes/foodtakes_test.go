package foodtakes

import "testing"

func TestGetRandomTakeWithSeedDeterministic(t *testing.T) {
	a := GetRandomTakeWithSeed(20260804)
	b := GetRandomTakeWithSeed(20260804)
	if a.Text != b.Text {
		t.Error("same seed should return the same take")
	}
	if a.Text == "" || a.Category == "" {
		t.Error("take should have text and category")
	}
}

func TestGetRandomTakeVariesByDay(t *testing.T) {
	seen := map[string]bool{}
	for day := 1; day <= 10; day++ {
		take := GetRandomTakeWithSeed(20260800 + day)
		seen[take.Text] = true
	}
	if len(seen) < 2 {
		t.Error("expected different takes across different days")
	}
}

func TestAllTakesValid(t *testing.T) {
	takes := getAllTakes()
	if len(takes) < 40 {
		t.Errorf("expected at least 40 takes, got %d", len(takes))
	}
	seen := map[string]bool{}
	for _, take := range takes {
		if take.Text == "" || take.Category == "" {
			t.Errorf("take with empty text or category: %+v", take)
		}
		if seen[take.Text] {
			t.Errorf("duplicate take: %s", take.Text)
		}
		seen[take.Text] = true
	}
}
