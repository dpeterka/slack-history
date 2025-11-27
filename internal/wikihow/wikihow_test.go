package wikihow

import (
	"testing"
)

func TestGetRandomArticle(t *testing.T) {
	article := GetRandomArticle()

	if article.Title == "" {
		t.Error("Expected non-empty article title")
	}

	if article.URL == "" {
		t.Error("Expected non-empty article URL")
	}
}

func TestGetRandomArticleConsistency(t *testing.T) {
	// Should return same article for same day
	article1 := GetRandomArticle()
	article2 := GetRandomArticle()

	if article1.Title != article2.Title {
		t.Error("Expected same article for the same day")
	}

	if article1.URL != article2.URL {
		t.Error("Expected same URL for the same day")
	}
}

func TestGetRandomArticleWithSeed(t *testing.T) {
	// Test with specific seed
	seed := 20251225
	article := GetRandomArticleWithSeed(seed)

	if article.Title == "" {
		t.Error("Expected non-empty article title")
	}

	if article.URL == "" {
		t.Error("Expected non-empty article URL")
	}

	// Same seed should give same result
	article2 := GetRandomArticleWithSeed(seed)
	if article.Title != article2.Title {
		t.Error("Expected same article for same seed")
	}
}

func TestAllArticlesHaveContent(t *testing.T) {
	articles := getAllArticles()

	for i, article := range articles {
		if article.Title == "" {
			t.Errorf("Article %d has empty title", i)
		}

		if article.URL == "" {
			t.Errorf("Article %d has empty URL", i)
		}
	}
}

func TestMinimumNumberOfArticles(t *testing.T) {
	articles := getAllArticles()

	// Should have at least 30 articles
	if len(articles) < 30 {
		t.Errorf("Expected at least 30 articles, got %d", len(articles))
	}
}
