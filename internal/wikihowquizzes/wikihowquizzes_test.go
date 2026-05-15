package wikihowquizzes

import (
	"testing"
)

func TestGetRandomQuiz(t *testing.T) {
	quiz := GetRandomQuiz()

	if quiz.Title == "" {
		t.Error("Expected non-empty quiz title")
	}

	if quiz.URL == "" {
		t.Error("Expected non-empty quiz URL")
	}
}

func TestGetRandomQuizConsistency(t *testing.T) {
	quiz1 := GetRandomQuiz()
	quiz2 := GetRandomQuiz()

	if quiz1.Title != quiz2.Title {
		t.Error("Expected same quiz for the same day")
	}

	if quiz1.URL != quiz2.URL {
		t.Error("Expected same URL for the same day")
	}
}

func TestGetRandomQuizWithSeed(t *testing.T) {
	seed := 20251225
	quiz := GetRandomQuizWithSeed(seed)

	if quiz.Title == "" {
		t.Error("Expected non-empty quiz title")
	}

	if quiz.URL == "" {
		t.Error("Expected non-empty quiz URL")
	}

	quiz2 := GetRandomQuizWithSeed(seed)
	if quiz.Title != quiz2.Title {
		t.Error("Expected same quiz for same seed")
	}
}

func TestAllQuizzesHaveContent(t *testing.T) {
	quizzes := getAllQuizzes()

	for i, quiz := range quizzes {
		if quiz.Title == "" {
			t.Errorf("Quiz %d has empty title", i)
		}

		if quiz.URL == "" {
			t.Errorf("Quiz %d has empty URL", i)
		}
	}
}

func TestMinimumNumberOfQuizzes(t *testing.T) {
	quizzes := getAllQuizzes()

	if len(quizzes) < 30 {
		t.Errorf("Expected at least 30 quizzes, got %d", len(quizzes))
	}
}
