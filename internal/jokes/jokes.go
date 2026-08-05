package jokes

import (
	"github.com/dpeterka/history-slackbot/internal/rotation"
	"time"
)

// Joke represents a one-liner joke
type Joke struct {
	Text string
}

// GetRandomJoke returns a random joke using current date as seed
func GetRandomJoke() Joke {
	return GetRandomJokeWithSeed(0)
}

// GetRandomJokeWithSeed returns a random joke with optional seed override.
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomJokeWithSeed(seed int) Joke {
	jokes := getAllJokes()

	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	return jokes[rotation.PickIndex(len(jokes), seed)]
}

func getAllJokes() []Joke {
	return []Joke{
		{Text: "I told my wife she was drawing her eyebrows too high. She looked surprised."},
		{Text: "I used to be addicted to soap. I'm clean now."},
		{Text: "My therapist says I have a preoccupation with vengeance. We'll see about that."},
		{Text: "I told my wife to stop impersonating a flamingo. I had to put my foot down."},
		{Text: "I have a fear of speed bumps. I'm slowly getting over it."},
		{Text: "I'm reading a book about anti-gravity. It's impossible to put down."},
		{Text: "Why don't scientists trust atoms? Because they make up everything."},
		{Text: "I told my doctor I broke my arm in two places. He told me to stop going to those places."},
		{Text: "Parallel lines have so much in common. It's a shame they'll never meet."},
		{Text: "I told a chemistry joke once. There was no reaction."},
		{Text: "My wife told me I had to stop acting like a flamingo. So I put my foot down. Then she told me again, so I put the other one down. Now I'm just standing here."},
		{Text: "I'm on a seafood diet. I see food and I eat it."},
		{Text: "I tried to catch fog yesterday. Mist."},
		{Text: "I told my computer I needed a break. Now it won't stop sending me Kit Kat ads."},
		{Text: "I asked my electrician to fix the doorbell. He said 'I'll get back to you.' Three weeks later — no ring."},
		{Text: "I quit my job at the helium factory. I refused to be spoken to in that tone."},
		{Text: "I bought shoes from a drug dealer once. I don't know what he laced them with, but I was tripping all day."},
		{Text: "Why did the scarecrow win an award? Because he was outstanding in his field."},
		{Text: "I changed my password to 'incorrect.' Now when I forget it, my computer just tells me."},
		{Text: "I went to a bookstore and asked the saleswoman where the self-help section was. She said if she told me, it would defeat the purpose."},
		{Text: "I told my wife I was going to make a belt out of watches. She said it'd be a waist of time."},
		{Text: "What do you call a fish with no eyes? A fsh."},
		{Text: "I lost my mood ring yesterday. I don't know how I feel about it."},
		{Text: "Why don't skeletons fight each other? They don't have the guts."},
		{Text: "I used to hate facial hair. But then it grew on me."},
		{Text: "I'm terrified of elevators. I'm taking steps to avoid them."},
		{Text: "I just got fired from the calendar factory. All I did was take a day off."},
		{Text: "The rotation of the Earth really makes my day."},
		{Text: "Singing in the shower is fun until you get soap in your mouth. Then it's a soap opera."},
		{Text: "I don't trust stairs. They're always up to something."},
		{Text: "What do you call a fake noodle? An impasta."},
		{Text: "Why did the math book look so sad? It had too many problems."},
		{Text: "I went to buy camouflage trousers but I couldn't find any."},
		{Text: "My wife asked if I'd seen the dog bowl. I said I didn't know he could."},
		{Text: "I'm reading a book on the history of glue. I can't seem to put it down."},
		{Text: "Why don't oysters share their pearls? Because they're shellfish."},
		{Text: "I told my wife she was painting our ceiling wrong. She said 'I'd like to see you try.' So I did. Now we have a new ceiling and a new wife."},
		{Text: "I bought the world's worst thesaurus yesterday. Not only is it terrible, it's also terrible."},
		{Text: "I used to be a banker but I lost interest."},
		{Text: "What do you get when you cross a snowman with a vampire? Frostbite."},
		{Text: "Why did the bicycle fall over? Because it was two-tired."},
		{Text: "I tried to sue the airline for losing my luggage. I lost the case."},
		{Text: "My friend's bakery burned down last night. Now his business is toast."},
		{Text: "I'm writing a book about hurricanes and tornadoes. It's only a draft right now."},
		{Text: "Time flies like an arrow. Fruit flies like a banana."},
		{Text: "I told my wife I was leaving her because she's too obsessed with counting. She gave me a 1, a 2, a 3..."},
		{Text: "I invented a new word: plagiarism."},
		{Text: "Why do bees have sticky hair? Because they use honeycombs."},
		{Text: "I asked the librarian if the library had books about paranoia. She whispered, 'They're right behind you.'"},
		{Text: "My dog used to chase people on a bike a lot. It got so bad I had to take his bike away."},
		{Text: "I only know 25 letters of the alphabet. I don't know y."},
		{Text: "What's the best thing about Switzerland? I don't know, but the flag is a big plus."},
		{Text: "I stayed up all night wondering where the sun went. Then it dawned on me."},
		{Text: "Did you hear about the claustrophobic astronaut? He just needed a little space."},
		{Text: "Why can't you hear a pterodactyl go to the bathroom? Because the P is silent."},
		{Text: "I used to play piano by ear, but now I use my hands."},
		{Text: "What did the ocean say to the beach? Nothing, it just waved."},
		{Text: "I got a job at a bakery because I kneaded dough."},
		{Text: "Why did the golfer bring two pairs of pants? In case he got a hole in one."},
		{Text: "I'm afraid for the calendar. Its days are numbered."},
		{Text: "What do you call a belt made of watches? A waist of time. My wife's heard that one. She left anyway."},
		{Text: "How do you make a tissue dance? You put a little boogie in it."},
		{Text: "I ordered a chicken and an egg online. I'll let you know."},
		{Text: "What do you call cheese that isn't yours? Nacho cheese."},
		{Text: "Why don't eggs tell jokes? They'd crack each other up."},
		{Text: "I don't play soccer because I enjoy the sport. I'm just doing it for kicks."},
		{Text: "What's brown and sticky? A stick."},
		{Text: "Two guys walked into a bar. The third one ducked."},
		{Text: "My boss told me to have a good day. So I went home."},
		{Text: "Six out of seven dwarfs aren't Happy. Statistically, that tracks."},
		{Text: "I was going to tell a time-traveling joke, but you didn't like it."},
	}
}
