package emo

import (
	"math/rand"
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
	comments := []Comment{
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

		// Relationships
		{Text: "Love is just two people agreeing to be weird together until one of them gets tired.", Category: "Relationships"},
		{Text: "Relationships are like group projects where you're paired with someone as dysfunctional as you are.", Category: "Relationships"},
		{Text: "They say communication is key. Nobody mentioned the door is locked from the inside.", Category: "Relationships"},
		{Text: "Every relationship is just two damaged people deciding their traumas are compatible.", Category: "Relationships"},
		{Text: "Being vulnerable is important. So is having a good therapist on speed dial.", Category: "Relationships"},
		{Text: "Love languages: acts of service, quality time, and pretending you didn't hear that comment.", Category: "Relationships"},
		{Text: "Relationships are 50% compromise and 50% wondering if you should have gotten a cat instead.", Category: "Relationships"},
		{Text: "We're all just looking for someone who will tolerate our 3am anxiety spirals.", Category: "Relationships"},
		{Text: "True love is finding someone whose emotional baggage fits in your trunk.", Category: "Relationships"},
		{Text: "Intimacy is sharing your deepest fears and then immediately regretting it.", Category: "Relationships"},
		{Text: "A healthy relationship is two people who occasionally like each other at the same time.", Category: "Relationships"},
		{Text: "They say opposites attract. Mostly they just confuse each other loudly.", Category: "Relationships"},

		// Mixed/General Angst
		{Text: "Some days you're the pigeon. Other days you're the statue. Most days you're both.", Category: "Life"},
		{Text: "Happiness is temporary. Awkwardness is forever. Plan accordingly.", Category: "Life"},
		{Text: "We spend our whole lives pretending we're not all slowly falling apart.", Category: "Life"},
		{Text: "The only constant is change, and that change is usually for the weird.", Category: "Life"},
		{Text: "Adulthood: where you're simultaneously too old and too young for everything.", Category: "Life"},
		{Text: "Some people have their life together. I have a collection of expired coupons and regrets.", Category: "Life"},
		{Text: "Life's a journey. Unfortunately, I forgot to check the map and now I'm lost in a Target parking lot.", Category: "Life"},
		{Text: "We're all just extras in everyone else's story, hoping someone notices our cameo.", Category: "Life"},
	}

	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}
	r := rand.New(rand.NewSource(int64(seed)))

	return comments[r.Intn(len(comments))]
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
	r := rand.New(rand.NewSource(int64(seed)))

	return filtered[r.Intn(len(filtered))]
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

		// Relationships
		{Text: "Love is just two people agreeing to be weird together until one of them gets tired.", Category: "Relationships"},
		{Text: "Relationships are like group projects where you're paired with someone as dysfunctional as you are.", Category: "Relationships"},
		{Text: "They say communication is key. Nobody mentioned the door is locked from the inside.", Category: "Relationships"},
		{Text: "Every relationship is just two damaged people deciding their traumas are compatible.", Category: "Relationships"},
		{Text: "Being vulnerable is important. So is having a good therapist on speed dial.", Category: "Relationships"},
		{Text: "Love languages: acts of service, quality time, and pretending you didn't hear that comment.", Category: "Relationships"},
		{Text: "Relationships are 50% compromise and 50% wondering if you should have gotten a cat instead.", Category: "Relationships"},
		{Text: "We're all just looking for someone who will tolerate our 3am anxiety spirals.", Category: "Relationships"},
		{Text: "True love is finding someone whose emotional baggage fits in your trunk.", Category: "Relationships"},
		{Text: "Intimacy is sharing your deepest fears and then immediately regretting it.", Category: "Relationships"},
		{Text: "A healthy relationship is two people who occasionally like each other at the same time.", Category: "Relationships"},
		{Text: "They say opposites attract. Mostly they just confuse each other loudly.", Category: "Relationships"},
	}
}
