package rotation

import (
	"testing"
	"time"
)

// seedFor converts a time.Time to the YYYYMMDD seed format used by the packs.
func seedFor(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}

func TestPickIndexDeterministic(t *testing.T) {
	if PickIndex(50, 20260804) != PickIndex(50, 20260804) {
		t.Error("same seed should always yield the same index")
	}
}

func TestPickIndexNoRepeatsWithinCycle(t *testing.T) {
	n := 46 // size of the smaller packs
	start := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)

	seen := map[int][]time.Time{}
	for i := 0; i < 365; i++ {
		d := start.AddDate(0, 0, i)
		idx := PickIndex(n, seedFor(d))
		if idx < 0 || idx >= n {
			t.Fatalf("index %d out of range [0,%d)", idx, n)
		}
		seen[idx] = append(seen[idx], d)
	}

	// A full year over a 46-item pack must surface every index.
	if len(seen) != n {
		t.Errorf("expected all %d indices used over a year, got %d", n, len(seen))
	}
}

func TestPickIndexExhaustsPackBeforeRepeating(t *testing.T) {
	n := 46
	// Walk from the epoch so days align with cycle boundaries exactly.
	start := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	inCycle := map[int]bool{}
	for i := 0; i < n; i++ {
		d := start.AddDate(0, 0, i)
		idx := PickIndex(n, seedFor(d))
		if inCycle[idx] {
			t.Fatalf("index %d repeated within one %d-day cycle", idx, n)
		}
		inCycle[idx] = true
	}
	if len(inCycle) != n {
		t.Errorf("expected %d distinct indices in one cycle, got %d", n, len(inCycle))
	}
}

func TestPickIndexSmallPacks(t *testing.T) {
	if PickIndex(0, 20260804) != 0 || PickIndex(1, 20260804) != 0 {
		t.Error("packs of size 0 or 1 should return index 0")
	}
}
