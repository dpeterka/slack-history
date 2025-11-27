package wikipedia

import (
	"testing"
)

func TestFetchBirthsAndDeaths(t *testing.T) {
	// Test with November 20
	people, err := FetchBirthsAndDeaths(11, 20, 5)
	if err != nil {
		t.Fatalf("Failed to fetch births and deaths: %v", err)
	}

	if len(people) == 0 {
		t.Error("Expected at least some people")
	}

	if len(people) > 5 {
		t.Errorf("Expected max 5 people, got %d", len(people))
	}

	// Check that we have valid data
	for i, person := range people {
		if person.Name == "" {
			t.Errorf("Person %d has empty name", i)
		}
		if person.Year == 0 {
			t.Errorf("Person %d has zero year", i)
		}
		if person.Type != "birth" && person.Type != "death" {
			t.Errorf("Person %d has invalid type: %s", i, person.Type)
		}
	}
}

func TestFetchBirths(t *testing.T) {
	// Test with November 20
	people, err := FetchBirths(11, 20, 3)
	if err != nil {
		t.Fatalf("Failed to fetch births: %v", err)
	}

	if len(people) == 0 {
		t.Error("Expected at least some births")
	}

	if len(people) > 3 {
		t.Errorf("Expected max 3 people, got %d", len(people))
	}

	// All should be births
	for i, person := range people {
		if person.Type != "birth" {
			t.Errorf("Person %d should be birth, got %s", i, person.Type)
		}
		if person.Name == "" {
			t.Errorf("Person %d has empty name", i)
		}
	}
}

func TestFetchDeaths(t *testing.T) {
	// Test with November 20
	people, err := FetchDeaths(11, 20, 3)
	if err != nil {
		t.Fatalf("Failed to fetch deaths: %v", err)
	}

	if len(people) == 0 {
		t.Error("Expected at least some deaths")
	}

	if len(people) > 3 {
		t.Errorf("Expected max 3 people, got %d", len(people))
	}

	// All should be deaths
	for i, person := range people {
		if person.Type != "death" {
			t.Errorf("Person %d should be death, got %s", i, person.Type)
		}
		if person.Name == "" {
			t.Errorf("Person %d has empty name", i)
		}
	}
}

func TestFetchTodaysBirthsAndDeaths(t *testing.T) {
	people, err := FetchTodaysBirthsAndDeaths(5)
	if err != nil {
		t.Fatalf("Failed to fetch today's births and deaths: %v", err)
	}

	if len(people) == 0 {
		t.Error("Expected at least some people")
	}

	if len(people) > 5 {
		t.Errorf("Expected max 5 people, got %d", len(people))
	}
}
