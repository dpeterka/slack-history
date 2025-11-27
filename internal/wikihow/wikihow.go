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
			Title: "How to Avoid Conversation",
			URL:   "https://www.wikihow.com/Avoid-Conversation",
		},
		{
			Title: "How to Act Like You Don't Care",
			URL:   "https://www.wikihow.com/Act-Like-You-Don%27t-Care",
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
			Title: "How to Eat Glass",
			URL:   "https://www.wikihow.com/Eat-Glass",
		},
		{
			Title: "How to Hide a Body",
			URL:   "https://www.wikihow.com/Hide-a-Body",
		},
		{
			Title: "How to Pretend to Know a Language",
			URL:   "https://www.wikihow.com/Pretend-to-Know-a-Language",
		},
		{
			Title: "How to Get Out of Going to Church",
			URL:   "https://www.wikihow.com/Get-Out-of-Going-to-Church",
		},
		{
			Title: "How to Deal With Annoying People",
			URL:   "https://www.wikihow.com/Deal-With-Annoying-People",
		},
		{
			Title: "How to Escape from Quicksand",
			URL:   "https://www.wikihow.com/Escape-from-Quicksand",
		},
		{
			Title: "How to Survive a Zombie Apocalypse",
			URL:   "https://www.wikihow.com/Survive-a-Zombie-Apocalypse",
		},
		{
			Title: "How to Talk to a Girl With a Boyfriend",
			URL:   "https://www.wikihow.com/Talk-to-a-Girl-With-a-Boyfriend",
		},
		{
			Title: "How to Be Mysterious",
			URL:   "https://www.wikihow.com/Be-Mysterious",
		},
		{
			Title: "How to Look Busy at Work Without Actually Working",
			URL:   "https://www.wikihow.com/Look-Busy-at-Work-Without-Actually-Working",
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
			Title: "How to Get Away With Not Doing Your Homework",
			URL:   "https://www.wikihow.com/Get-Away-With-Not-Doing-Your-Homework",
		},
		{
			Title: "How to Survive a Bear Attack",
			URL:   "https://www.wikihow.com/Survive-a-Bear-Attack",
		},
		{
			Title: "How to Avoid Doing Something You Don't Want to Do",
			URL:   "https://www.wikihow.com/Avoid-Doing-Something-You-Don%27t-Want-to-Do",
		},
		{
			Title: "How to Know if You're in Love or Just Lonely",
			URL:   "https://www.wikihow.com/Know-if-You%27re-in-Love-or-Just-Lonely",
		},
		{
			Title: "How to Communicate With the Dead",
			URL:   "https://www.wikihow.com/Communicate-With-the-Dead",
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
			Title: "How to Tell if Your Cat Is Plotting to Kill You",
			URL:   "https://www.wikihow.com/Tell-if-Your-Cat-Is-Plotting-to-Kill-You",
		},
		{
			Title: "How to Become a Mermaid",
			URL:   "https://www.wikihow.com/Become-a-Mermaid",
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
			Title: "How to Tell if You're a Witch",
			URL:   "https://www.wikihow.com/Tell-if-You%27re-a-Witch",
		},
		{
			Title: "How to Get Revenge on Your Enemies",
			URL:   "https://www.wikihow.com/Get-Revenge-on-Your-Enemies",
		},
		{
			Title: "How to Fake Amnesia",
			URL:   "https://www.wikihow.com/Fake-Amnesia",
		},
		{
			Title: "How to Survive Prison",
			URL:   "https://www.wikihow.com/Survive-Prison",
		},
		{
			Title: "How to Survive in the Woods",
			URL:   "https://www.wikihow.com/Survive-in-the-Woods",
		},
		{
			Title: "How to Prepare for a Disaster",
			URL:   "https://www.wikihow.com/Prepare-for-a-Disaster",
		},
		{
			Title: "How to Talk to Short People",
			URL:   "https://www.wikihow.com/Talk-to-Short-People",
		},
		{
			Title: "How to Act Like a Psychopath",
			URL:   "https://www.wikihow.com/Act-Like-a-Psychopath",
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
			Title: "How to Tell Someone at Work that They Smell Bad",
			URL:   "https://www.wikihow.com/Tell-Someone-at-Work-that-They-Smell-Bad",
		},
	}
}
