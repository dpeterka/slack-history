package blobby

import (
	"github.com/dpeterka/history-slackbot/internal/rotation"
	"time"
)

// Fact represents a Mr Blobby fact
type Fact struct {
	Text     string
	Category string
}

// GetRandomFact returns a random Mr Blobby fact using current date as seed
func GetRandomFact() Fact {
	return GetRandomFactWithSeed(0)
}

// GetRandomFactWithSeed returns a random Mr Blobby fact with optional seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomFactWithSeed(seed int) Fact {
	facts := getAllFacts()

	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	return facts[rotation.PickIndex(len(facts), seed)]
}

// GetRandomFactByCategory returns a random fact from a specific category
func GetRandomFactByCategory(category string) Fact {
	var filtered []Fact
	allFacts := getAllFacts()

	for _, fact := range allFacts {
		if fact.Category == category {
			filtered = append(filtered, fact)
		}
	}

	if len(filtered) == 0 {
		return GetRandomFact()
	}

	// Use current date for seed
	now := time.Now()
	seed := now.Year()*10000 + int(now.Month())*100 + now.Day()

	return filtered[rotation.PickIndex(len(filtered), seed)]
}

// getAllFacts returns all Mr Blobby facts
func getAllFacts() []Fact {
	return []Fact{
		// Origins & Creation
		{
			Text:     "Mr Blobby was created in 1992 for Noel Edmonds' BBC show 'Noel's House Party' as a parody of children's TV characters.",
			Category: "Origins",
		},
		{
			Text:     "The character was designed to be deliberately annoying and was meant to catch celebrities off-guard by pretending to be a real children's TV character.",
			Category: "Origins",
		},
		{
			Text:     "Mr Blobby's costume weighs approximately 35 pounds (16 kg) and requires a person inside to operate, severely limiting visibility.",
			Category: "Origins",
		},
		{
			Text:     "The costume was so heavy and cumbersome that operators could only wear it for 20-30 minutes at a time before needing a break.",
			Category: "Origins",
		},

		// Pop Culture Success
		{
			Text:     "In 1993, Mr Blobby released a single called 'Mr Blobby' which reached #1 on the UK Singles Chart, beating Take That's 'Babe' for the Christmas #1 spot.",
			Category: "Music",
		},
		{
			Text:     "The 'Mr Blobby' song is considered one of the worst Christmas #1 singles in UK chart history, yet it sold over 500,000 copies.",
			Category: "Music",
		},
		{
			Text:     "Mr Blobby released a second single 'Christmas in Blobbyland' in 1995, which peaked at #36 on the UK charts.",
			Category: "Music",
		},
		{
			Text:     "Despite being a novelty character, Mr Blobby has released two albums: 'Mr Blobby' (1993) and 'Mr Blobby's Christmas' (1995).",
			Category: "Music",
		},

		// Theme Parks & Blobbyland
		{
			Text:     "Blobbyland theme park opened in Cricket St Thomas, Somerset in 1994. It closed after just 13 months due to poor attendance and numerous complaints.",
			Category: "Theme Park",
		},
		{
			Text:     "The Blobbyland park was plagued with issues: rides breaking down, staff in costume getting heatstroke, and general chaos. It's considered one of the UK's worst theme park disasters.",
			Category: "Theme Park",
		},
		{
			Text:     "After Blobbyland closed, the site fell into disrepair and became known for urban exploration. Vandalized Mr Blobby statues became internet memes in the 2010s.",
			Category: "Theme Park",
		},
		{
			Text:     "A second 'Crinkley Bottom' (Mr Blobby's fictional home) attraction opened at Pleasure Beach Blackpool in 1994 but also closed within a year.",
			Category: "Theme Park",
		},
		{
			Text:     "The abandoned Blobbyland ruins featured decaying Mr Blobby figures that were genuinely terrifying and became viral content on Reddit and YouTube.",
			Category: "Theme Park",
		},

		// Television Appearances
		{
			Text:     "Mr Blobby appeared in over 120 episodes of 'Noel's House Party' between 1992-1999, making him one of the show's most memorable recurring characters.",
			Category: "TV",
		},
		{
			Text:     "The character was so popular that he received his own BBC show 'Mr Blobby' in 1994, which ran for one season before being cancelled.",
			Category: "TV",
		},
		{
			Text:     "Mr Blobby appeared in a controversial 1994 episode where he 'pranked' celebrities by destroying expensive sets and equipment - some celebrities were genuinely upset.",
			Category: "TV",
		},
		{
			Text:     "The character made a brief return in 2009 for a 'Noel's House Party' reunion special, proving his enduring (and baffling) popularity.",
			Category: "TV",
		},

		// Cultural Impact & Legacy
		{
			Text:     "In 2014, Mr Blobby was voted the scariest British icon of all time in an online poll, beating out characters like the Daleks and Weeping Angels.",
			Category: "Legacy",
		},
		{
			Text:     "Mr Blobby has been cited as an influence on avant-garde performance art, with some calling him 'accidentally surrealist' and 'chaos incarnate.'",
			Category: "Legacy",
		},
		{
			Text:     "The phrase 'Blobby Blobby Blobby!' (his only vocabulary) has entered British slang as an expression of joyful chaos or absurdity.",
			Category: "Legacy",
		},
		{
			Text:     "Mr Blobby costumes are still worn at British sporting events, particularly cricket matches, where fans dress as him to confuse opposing teams.",
			Category: "Legacy",
		},
		{
			Text:     "Internet memes about Mr Blobby peaked in the 2010s, with 'creepy Mr Blobby' becoming a horror meme format on sites like Reddit and Tumblr.",
			Category: "Legacy",
		},

		// Controversies & Incidents
		{
			Text:     "During a 1994 appearance, Mr Blobby accidentally knocked over and broke a £12,000 ornamental vase on live television. The incident was left in the broadcast.",
			Category: "Controversies",
		},
		{
			Text:     "Multiple performers inside the Mr Blobby costume have reported injuries, heat exhaustion, and psychological distress from the limited visibility and weight.",
			Category: "Controversies",
		},
		{
			Text:     "In 1995, Mr Blobby 'attacked' then-Prime Minister John Major during a charity event, tackling him to the ground. Major was not amused.",
			Category: "Controversies",
		},
		{
			Text:     "The character has been banned from several venues after incidents of property damage, including a shopping centre in 1993 where he destroyed a Christmas display.",
			Category: "Controversies",
		},
		{
			Text:     "A 2009 study suggested that young children found Mr Blobby more frightening than comforting, with some developing temporary 'Blobby phobia.'",
			Category: "Controversies",
		},

		// Merchandise & Commercial Success
		{
			Text:     "At the height of Blobby-mania in 1994, over £35 million worth of Mr Blobby merchandise was sold in the UK, including toys, clothing, and home goods.",
			Category: "Commercial",
		},
		{
			Text:     "The Mr Blobby doll was one of the top-selling toys of Christmas 1994, despite many children finding it frightening rather than endearing.",
			Category: "Commercial",
		},
		{
			Text:     "Mr Blobby appeared in advertisements for several brands in the 1990s, though most companies later distanced themselves from the character.",
			Category: "Commercial",
		},
		{
			Text:     "A 1995 Mr Blobby video game was released for Sega Mega Drive and other platforms. It's consistently rated as one of the worst video games ever made.",
			Category: "Commercial",
		},

		// Modern References & Revivals
		{
			Text:     "Mr Blobby made a surprise cameo in the 2023 Doctor Who Christmas special, confusing international viewers who had no context for the character.",
			Category: "Modern",
		},
		{
			Text:     "The character has been referenced in shows like 'Black Mirror' and 'Inside No. 9' as a symbol of unsettling 1990s British nostalgia.",
			Category: "Modern",
		},
		{
			Text:     "In 2020, someone purchased a Mr Blobby costume at auction for £62,000, claiming it was 'an important piece of British cultural history.'",
			Category: "Modern",
		},
		{
			Text:     "A 2022 Netflix documentary about 1990s British TV dedicated an entire episode to Mr Blobby, titled 'Chaos in a Pink Suit.'",
			Category: "Modern",
		},
		{
			Text:     "Mr Blobby's Twitter account (yes, he has one) has over 50,000 followers who share memes and cursed images of the character.",
			Category: "Modern",
		},

		// Weird & Wonderful Facts
		{
			Text:     "Mr Blobby's design was partially inspired by the Michelin Man, but 'if he'd been left in the sun too long and gone slightly mad.'",
			Category: "Design",
		},
		{
			Text:     "The character has no coherent backstory - different appearances give him different origins, occupations, and even different numbers of spots.",
			Category: "Design",
		},
		{
			Text:     "Mr Blobby appears to be both a sentient being and a costume worn by different characters within the show's universe - the lore is deliberately contradictory.",
			Category: "Design",
		},
		{
			Text:     "In 2015, philosophers at Oxford University used Mr Blobby as an example of 'accidental absurdism' in a paper about modern comedy.",
			Category: "Academic",
		},
		{
			Text:     "Mr Blobby was almost sent to represent the UK at Eurovision 1994, with a campaign gaining 50,000 signatures before BBC rejected the idea.",
			Category: "Almost Happened",
		},
		{
			Text:     "The character inspired a wave of similar 'annoying mascot' characters across Europe, though none achieved the same level of infamy.",
			Category: "Influence",
		},
		{
			Text:     "In a 2019 poll, Mr Blobby was voted 'Most Chaotic Cultural Icon' by British millennials, beating out other contenders like the Crazy Frog.",
			Category: "Legacy",
		},
		{
			Text:     "The Mr Blobby costume has been exhibited in the Museum of British Broadcasting, labeled as 'a cautionary tale about 90s excess.'",
			Category: "Legacy",
		},
		{
			Text:     "Noel Edmonds has stated that Mr Blobby is his 'greatest creation and biggest regret,' often in the same sentence.",
			Category: "Creator",
		},
		{
			Text:     "During a 2018 interview, Edmonds revealed that Mr Blobby was meant to appear for just three episodes but became 'a monster we couldn't control.'",
			Category: "Creator",
		},
		{
			Text:     "Mr Blobby's 1993 single knocked Take That's 'Babe' off the UK #1 spot just in time for Christmas, then reclaimed it after briefly losing it — one of the pettiest chart battles in British history.",
			Category: "Music",
		},
		{
			Text:     "The 'Mr Blobby' single sold over 700,000 copies. That's 700,000 separate human decisions.",
			Category: "Music",
		},
		{
			Text:     "Actor Barry Killerby performed as Mr Blobby for most of the character's run, a role he reportedly approached with the seriousness of classical theatre.",
			Category: "Origins",
		},
		{
			Text:     "Mr Blobby is pink with yellow spots, has permanently startled eyes, and wears a green bow tie — a design committee decision that no one has ever taken responsibility for.",
			Category: "Design",
		},
		{
			Text:     "The Morecambe 'Crinkley Bottom' park closed after just 13 weeks, triggering a council scandal known locally as 'Blobbygate' and an official inquiry that cost taxpayers £2.5 million.",
			Category: "Theme Park",
		},
		{
			Text:     "Blobbygate remains one of the few British political scandals named after a children's TV character, a distinction historians have chosen not to celebrate.",
			Category: "Controversies",
		},
		{
			Text:     "Mr Blobby's house in Blobbyland, 'Dunblobbin,' was a real cottage with melted-looking pink architecture. Urban explorers still make pilgrimages to its remains.",
			Category: "Theme Park",
		},
		{
			Text:     "In 2023, an original Mr Blobby costume listed on eBay attracted a winning bid of £62,000 before the buyer got cold feet — proving even Blobby's collectors experience sudden clarity.",
			Category: "Modern",
		},
		{
			Text:     "Mr Blobby turned 30 in 2022 and marked the occasion with a social media campaign that several outlets described as 'a threat.'",
			Category: "Modern",
		},
		{
			Text:     "Mr Blobby once appeared on live TV wrestling Noel Edmonds to the floor, unscripted, because the performer inside couldn't see where he was going.",
			Category: "TV",
		},
		{
			Text:     "The Mr Blobby costume's eyes are fixed in an expression of eternal surprise, which performers say accurately reflects the experience of being inside it.",
			Category: "Design",
		},
		{
			Text:     "Mr Blobby merchandise included duvet covers, meaning thousands of British children voluntarily slept under his gaze.",
			Category: "Commercial",
		},
		{
			Text:     "A Mr Blobby balloon once escaped its handlers at a 1994 parade and had to be retrieved from a rooftop, in what witnesses describe as the most on-brand moment of his career.",
			Category: "Controversies",
		},
		{
			Text:     "Mr Blobby's catchphrase 'Blobby Blobby Blobby' has been analyzed by linguists as a complete, self-contained language with exactly one morpheme.",
			Category: "Academic",
		},
		{
			Text:     "BBC archives reportedly contain hours of unaired Mr Blobby footage deemed 'too chaotic' for broadcast, which given what did air is a genuinely frightening thought.",
			Category: "TV",
		},
		{
			Text:     "Mr Blobby was nearly given a spin-off cartoon series in 1995. The pilot was never released, and animators involved have declined to discuss it.",
			Category: "Almost Happened",
		},
		{
			Text:     "A Mr Blobby-themed restaurant was proposed for London's West End in 1994 at the peak of Blobby-mania. Investors withdrew when the mania ended, roughly eleven minutes later.",
			Category: "Almost Happened",
		},
		{
			Text:     "Mr Blobby has been played by multiple performers over the decades, all of whom describe the experience with the same thousand-yard stare.",
			Category: "Legacy",
		},
		{
			Text:     "In 2021, a fan-made horror short film reimagining Mr Blobby as a slasher villain went viral, requiring almost no changes to the source material.",
			Category: "Modern",
		},
		{
			Text:     "Mr Blobby's chart success meant he technically outsold Nirvana in the UK singles market in December 1993, a sentence that has ended arguments and friendships.",
			Category: "Music",
		},
		{
			Text:     "The original Blobbyland site in Somerset was finally demolished in 2024, and locals reported the bulldozers 'hesitated.'",
			Category: "Theme Park",
		},
		{
			Text:     "Mr Blobby once presented a segment on breakfast television that had to be cut short when he destroyed the weather map, delaying the forecast for the entire nation.",
			Category: "TV",
		},
		{
			Text:     "Academics have noted that Mr Blobby predates and arguably perfected the 'chaos agent' archetype later seen in internet culture — he was posting cursed content before the internet could.",
			Category: "Academic",
		},
		{
			Text:     "A wax figure of Mr Blobby was proposed for Madame Tussauds in 1994 but rejected on the grounds that he is 'already essentially a wax figure come to life.'",
			Category: "Almost Happened",
		},
		{
			Text:     "Despite everything — the injuries, the scandals, the £2.5 million inquiry — Mr Blobby remains beloved, which says more about Britain than any history book ever could.",
			Category: "Legacy",
		},
	}
}
