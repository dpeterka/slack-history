package camping

import (
	"testing"
)

func TestGetRandomTip(t *testing.T) {
	tip := GetRandomTip()

	if tip.Text == "" {
		t.Error("Expected non-empty tip text")
	}

	if tip.Category == "" {
		t.Error("Expected non-empty tip category")
	}
}

func TestGetRandomTipConsistency(t *testing.T) {
	tip1 := GetRandomTip()
	tip2 := GetRandomTip()

	if tip1.Text != tip2.Text {
		t.Error("Expected same tip for the same day")
	}
}

func TestGetRandomTipWithSeed(t *testing.T) {
	seed := 20251225
	tip := GetRandomTipWithSeed(seed)

	if tip.Text == "" {
		t.Error("Expected non-empty tip text")
	}

	tip2 := GetRandomTipWithSeed(seed)
	if tip.Text != tip2.Text {
		t.Error("Expected same tip for same seed")
	}
}

func TestAllTipsHaveContent(t *testing.T) {
	tips := getAllTips()

	for i, tip := range tips {
		if tip.Text == "" {
			t.Errorf("Tip %d has empty text", i)
		}

		if tip.Category == "" {
			t.Errorf("Tip %d has empty category", i)
		}
	}
}

func TestMinimumNumberOfTips(t *testing.T) {
	tips := getAllTips()

	if len(tips) < 40 {
		t.Errorf("Expected at least 40 tips, got %d", len(tips))
	}
}
