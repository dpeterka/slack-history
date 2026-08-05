package emo

import (
	"github.com/dpeterka/history-slackbot/internal/rotation"
	"time"
)

// Comment represents an emo/philosophical comment
type Comment struct {
	Text     string
	Category string
}

// GetRandomComment returns a random emo comment using current date as seed
func GetRandomComment() Comment {
	return GetRandomCommentWithSeed(0)
}

// GetRandomCommentWithSeed returns a random emo comment with optional seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomCommentWithSeed(seed int) Comment {
	comments := getAllComments()

	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	return comments[rotation.PickIndex(len(comments), seed)]
}

// GetRandomCommentByCategory returns a random comment from a specific category
func GetRandomCommentByCategory(category string) Comment {
	var filtered []Comment
	allComments := getAllComments()

	for _, comment := range allComments {
		if comment.Category == category {
			filtered = append(filtered, comment)
		}
	}

	if len(filtered) == 0 {
		return GetRandomComment()
	}

	// Use current date for seed
	now := time.Now()
	seed := now.Year()*10000 + int(now.Month())*100 + now.Day()

	return filtered[rotation.PickIndex(len(filtered), seed)]
}

// getAllComments returns all comments (used internally)
func getAllComments() []Comment {
	return []Comment{
		// Work
		{Text: "Another day, another existential crisis at the office. At least the coffee is consistent.", Category: "Work"},
		{Text: "They say 'do what you love.' Nobody mentioned the part where you learn to resent it.", Category: "Work"},
		{Text: "Meetings: where minutes are taken and hours are lost, much like our will to live.", Category: "Work"},
		{Text: "My work-life balance is perfect. I'm equally dissatisfied with both.", Category: "Work"},
		{Text: "Some people find purpose in their careers. I found a really good nap spot in the supply closet.", Category: "Work"},
		{Text: "The only thing harder than Monday morning is pretending everything is fine during Monday morning.", Category: "Work"},
		{Text: "Corporate synergy: a fancy term for collective suffering with better lighting.", Category: "Work"},
		{Text: "I'm not saying work is meaningless, but have you ever seen a cat worry about quarterly reports?", Category: "Work"},
		{Text: "Climbing the corporate ladder is exhausting when you're not even sure you want to be on the roof.", Category: "Work"},
		{Text: "Every email is a tiny scream into the void, and the void replies with 'per my last email.'", Category: "Work"},

		// Life
		{Text: "Life is a series of increasingly concerning surprises, punctuated by brief moments of snack time.", Category: "Life"},
		{Text: "We're all just walking contradictions wrapped in skin suits, pretending we have it together.", Category: "Life"},
		{Text: "The universe is indifferent to your struggles, but at least the stars look pretty while you suffer.", Category: "Life"},
		{Text: "Growing up is realizing that nobody actually knows what they're doing, they're just better at hiding it.", Category: "Life"},
		{Text: "Life tip: Lower your expectations. Then lower them again. Now you're getting somewhere.", Category: "Life"},
		{Text: "We're all protagonists in stories nobody's reading, including ourselves.", Category: "Life"},
		{Text: "Existential dread pairs well with coffee. Also with tea. And lunch. And basically everything.", Category: "Life"},
		{Text: "The human experience: 10% living, 90% overthinking what you said three years ago.", Category: "Life"},
		{Text: "Reality is just a collective hallucination we've all agreed to participate in.", Category: "Life"},
		{Text: "Some seek meaning in life. I'm just trying to remember where I put my keys.", Category: "Life"},
		{Text: "We're cosmic accidents on a floating rock, but somehow we still care about matching socks.", Category: "Life"},
		{Text: "Life is like a box of chocolates: confusing, occasionally disappointing, and someone already ate the good ones.", Category: "Life"},
		{Text: "Some days you're the pigeon. Other days you're the statue. Most days you're both.", Category: "Life"},
		{Text: "Happiness is temporary. Awkwardness is forever. Plan accordingly.", Category: "Life"},
		{Text: "We spend our whole lives pretending we're not all slowly falling apart.", Category: "Life"},
		{Text: "The only constant is change, and that change is usually for the weird.", Category: "Life"},
		{Text: "Adulthood: where you're simultaneously too old and too young for everything.", Category: "Life"},
		{Text: "Some people have their life together. I have a collection of expired coupons and regrets.", Category: "Life"},
		{Text: "Life's a journey. Unfortunately, I forgot to check the map and now I'm lost in a Target parking lot.", Category: "Life"},
		{Text: "We're all just extras in everyone else's story, hoping someone notices our cameo.", Category: "Life"},

		// Work (2026 refresh)
		{Text: "My calendar says 'focus time.' My soul says 'staring out the window time.' We compromise on neither.", Category: "Work"},
		{Text: "This meeting could have been an email. The email could have been nothing. Everything could have been nothing.", Category: "Work"},
		{Text: "They gave me a 'stretch assignment.' I am now stretched. Nothing else has changed.", Category: "Work"},
		{Text: "The printer knows when you're in a hurry. It feeds on urgency. We all serve something.", Category: "Work"},
		{Text: "I put 'quick sync' on your calendar. We both know nothing about it will be quick, and nothing will sync.", Category: "Work"},
		{Text: "Performance review season: where you describe your suffering in the past tense and call it growth.", Category: "Work"},
		{Text: "The office plant is dying and nobody waters it. It's the most honest coworker I have.", Category: "Work"},
		{Text: "'Let's take this offline' — the corporate way of saying this problem will now live forever, unsolved, in the dark.", Category: "Work"},
		{Text: "I've been 'circling back' so long I've achieved a stable orbit.", Category: "Work"},
		{Text: "My out-of-office reply is the only version of me with boundaries.", Category: "Work"},
		{Text: "Every all-hands starts with 'exciting news' the way every horror movie starts with a house at a great price.", Category: "Work"},
		{Text: "The break room coffee isn't good, but neither is anything else, so at least it's thematically consistent.", Category: "Work"},

		// Life (2026 refresh)
		{Text: "I bought a planner to organize my life. It's now a journal of things I didn't do.", Category: "Life"},
		{Text: "Sleep is just death being shy, and my insomnia means even death doesn't want to commit.", Category: "Life"},
		{Text: "My phone screen time report is the closest thing I have to an honest biography.", Category: "Life"},
		{Text: "Self-care is important. That's why I've scheduled my breakdown for a long weekend.", Category: "Life"},
		{Text: "The candle I lit for ambiance is the most alive thing in this apartment, and I resent its enthusiasm.", Category: "Life"},
		{Text: "Every houseplant I own is a tiny hospice I run with great optimism.", Category: "Life"},
		{Text: "I finally found inner peace. It was behind the couch, covered in dust, next to a remote from a TV I no longer own.", Category: "Life"},
		{Text: "New year, new me. Same debts, same knees, same inexplicable dread at 3 AM. But new.", Category: "Life"},
		{Text: "My step counter congratulated me today. It takes so little to impress something with no standards.", Category: "Life"},
		{Text: "They say time heals all wounds. Time also causes most of them. Time is playing both sides.", Category: "Life"},
		{Text: "I made a five-year plan once. That was six years ago. No further questions.", Category: "Life"},
		{Text: "The void stares back, but lately it just seems tired too, and honestly that's the most connection I've felt all week.", Category: "Life"},
		{Text: "Somewhere out there is the best version of me, and I hope he's not paying rent either.", Category: "Life"},
	}
}
