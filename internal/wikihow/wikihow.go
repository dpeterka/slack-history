package wikihow

import (
	"math/rand"
	"time"
)

// Article represents a WikiHow article
type Article struct {
	Title string
	URL   string
}

// GetRandomArticle returns a random WikiHow article using current date as seed
func GetRandomArticle() Article {
	return GetRandomArticleWithSeed(0)
}

// GetRandomArticleWithSeed returns a random WikiHow article with optional seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomArticleWithSeed(seed int) Article {
	articles := getAllArticles()

	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}
	r := rand.New(rand.NewSource(int64(seed)))

	return articles[r.Intn(len(articles))]
}

func getAllArticles() []Article {
	return []Article{
		{
			Title: "How to Pretend to Be Sick",
			URL:   "https://www.wikihow.com/Pretend-to-Be-Sick",
		},
		{
			Title: "How to Disappear Completely",
			URL:   "https://www.wikihow.com/Disappear-Completely",
		},
		{
			Title: "How to Pretend to Be Happy",
			URL:   "https://www.wikihow.com/Pretend-to-Be-Happy",
		},
		{
			Title: "How to Fake Your Own Death",
			URL:   "https://www.wikihow.com/Fake-Your-Own-Death",
		},
		{
			Title: "How to Live in Your Car",
			URL:   "https://www.wikihow.com/Live-in-Your-Car",
		},
		{
			Title: "How to Be Mysterious",
			URL:   "https://www.wikihow.com/Be-Mysterious",
		},
		{
			Title: "How to Survive a Plane Crash",
			URL:   "https://www.wikihow.com/Survive-a-Plane-Crash",
		},
		{
			Title: "How to Fake Confidence",
			URL:   "https://www.wikihow.com/Fake-Confidence",
		},
		{
			Title: "How to Lie Convincingly",
			URL:   "https://www.wikihow.com/Lie-Convincingly",
		},
		{
			Title: "How to Survive a Bear Attack",
			URL:   "https://www.wikihow.com/Survive-a-Bear-Attack",
		},
		{
			Title: "How to Survive If You Get Stranded on an Island",
			URL:   "https://www.wikihow.com/Survive-on-a-Deserted-Island",
		},
		{
			Title: "How to Make Fake Vomit",
			URL:   "https://www.wikihow.com/Make-Fake-Vomit",
		},
		{
			Title: "How to Cope With Having No Friends",
			URL:   "https://www.wikihow.com/Cope-With-Having-No-Friends",
		},
		{
			Title: "How to Start a Cult",
			URL:   "https://www.wikihow.com/Start-a-Cult",
		},
		{
			Title: "How to Survive a Shark Attack",
			URL:   "https://www.wikihow.com/Survive-a-Shark-Attack",
		},
		{
			Title: "How to Be a Ninja",
			URL:   "https://www.wikihow.com/Be-a-Ninja",
		},
		{
			Title: "How to Make Holy Water",
			URL:   "https://www.wikihow.com/Make-Holy-Water",
		},
		{
			Title: "How to Join the Illuminati",
			URL:   "https://www.wikihow.com/Join-the-Illuminati",
		},
		{
			Title: "How to Survive a Tornado",
			URL:   "https://www.wikihow.com/Survive-a-Tornado",
		},
		{
			Title: "How to Survive in the Woods",
			URL:   "https://www.wikihow.com/Survive-in-the-Woods",
		},
		{
			Title: "How to Talk to Short People",
			URL:   "https://www.wikihow.com/Talk-to-Short-People",
		},
		{
			Title: "How to Ignore People",
			URL:   "https://www.wikihow.com/Ignore-People",
		},
		{
			Title: "How to Tell Someone at Work that They Smell Bad",
			URL:   "https://www.wikihow.com/Tell-Someone-at-Work-that-They-Smell-Bad",
		},
		{
			Title: "How to Annoy People",
			URL:   "https://www.wikihow.com/Annoy-People",
		},
		{
			Title: "How to Be Random",
			URL:   "https://www.wikihow.com/Be-Random",
		},
		{
			Title: "How to Be Weird",
			URL:   "https://www.wikihow.com/Be-Weird",
		},
		{
			Title: "How to Act Drunk",
			URL:   "https://www.wikihow.com/Act-Drunk",
		},
		{
			Title: "How to Be Emo",
			URL:   "https://www.wikihow.com/Be-Emo",
		},
		{
			Title: "How to Summon a Demon",
			URL:   "https://www.wikihow.com/Summon-a-Demon",
		},
		{
			Title: "How to Become a Vampire",
			URL:   "https://www.wikihow.com/Become-a-Vampire",
		},
		{
			Title: "How to Become a Werewolf",
			URL:   "https://www.wikihow.com/Become-a-Werewolf",
		},
		{
			Title: "How to Levitate",
			URL:   "https://www.wikihow.com/Levitate",
		},
		{
			Title: "How to Read Minds",
			URL:   "https://www.wikihow.com/Read-Minds",
		},
		{
			Title: "How to Control Your Dreams",
			URL:   "https://www.wikihow.com/Control-Your-Dreams",
		},
		{
			Title: "How to Lucid Dream",
			URL:   "https://www.wikihow.com/Lucid-Dream",
		},
		{
			Title: "How to Fake a Fever",
			URL:   "https://www.wikihow.com/Fake-a-Fever",
		},
		{
			Title: "How to Fake Cry",
			URL:   "https://www.wikihow.com/Fake-Cry",
		},
		{
			Title: "How to Cry on Command",
			URL:   "https://www.wikihow.com/Cry-on-Command",
		},
		{
			Title: "How to Fake Sleep",
			URL:   "https://www.wikihow.com/Fake-Sleep",
		},
		{
			Title: "How to Survive in the Jungle",
			URL:   "https://www.wikihow.com/Survive-in-the-Jungle",
		},
		{
			Title: "How to Escape from a Sinking Car",
			URL:   "https://www.wikihow.com/Escape-from-a-Sinking-Car",
		},
		{
			Title: "How to Survive an Avalanche",
			URL:   "https://www.wikihow.com/Survive-an-Avalanche",
		},
		{
			Title: "How to Survive a Wildfire",
			URL:   "https://www.wikihow.com/Survive-a-Wildfire",
		},
		{
			Title: "How to Avoid Talking to People",
			URL:   "https://www.wikihow.com/Avoid-Talking-to-People",
		},
		{
			Title: "How to Get Rid of Annoying People",
			URL:   "https://www.wikihow.com/Get-Rid-of-Annoying-People",
		},
		{
			Title: "How to Fake Laugh",
			URL:   "https://www.wikihow.com/Fake-Laugh",
		},
		{
			Title: "How to Look Smart",
			URL:   "https://www.wikihow.com/Look-Smart",
		},
		{
			Title: "How to Skip School",
			URL:   "https://www.wikihow.com/Skip-School",
		},
		{
			Title: "How to Get Out of School",
			URL:   "https://www.wikihow.com/Get-Out-of-School",
		},
		{
			Title: "How to Forge a Signature",
			URL:   "https://www.wikihow.com/Forge-a-Signature",
		},
		{
			Title: "How to Make Yourself Cry",
			URL:   "https://www.wikihow.com/Make-Yourself-Cry",
		},
		{
			Title: "How to Win a Fight",
			URL:   "https://www.wikihow.com/Win-a-Fight",
		},
		{
			Title: "How to Defend Yourself in an Extreme Street Fight",
			URL:   "https://www.wikihow.com/Defend-Yourself-in-an-Extreme-Street-Fight",
		},
	}
}
