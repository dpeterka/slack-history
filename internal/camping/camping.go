package camping

import (
	"github.com/dpeterka/history-slackbot/internal/rotation"
	"time"
)

// Tip represents a camping tip
type Tip struct {
	Text     string
	Category string
}

// GetRandomTip returns a random camping tip using current date as seed
func GetRandomTip() Tip {
	return GetRandomTipWithSeed(0)
}

// GetRandomTipWithSeed returns a random camping tip with optional seed override.
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomTipWithSeed(seed int) Tip {
	tips := getAllTips()

	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	return tips[rotation.PickIndex(len(tips), seed)]
}

func getAllTips() []Tip {
	return []Tip{
		// Wildlife
		{Text: "Bears can smell your toothpaste from a mile away. Brush in the morning, regret it all night.", Category: "Wildlife"},
		{Text: "That funny sound at 3 AM is a raccoon doing whatever the hell it wants. You are not in charge here.", Category: "Wildlife"},
		{Text: "If a bear approaches your campsite, do not run. You'll die tired. Stand your ground and look big, like you do at family reunions.", Category: "Wildlife"},
		{Text: "Mosquitoes don't care about your DEET-free essential oil blend. They are eating you and your beliefs.", Category: "Wildlife"},
		{Text: "Hang your food in a bear bag 10 feet up and 4 feet out from the trunk. Or feed it to a bear, your call.", Category: "Wildlife"},
		{Text: "Snakes want nothing to do with you. Step heavy, watch your hands, and stop reaching under logs like a fool.", Category: "Wildlife"},
		{Text: "Ticks check in by the dozen and never leave. Permethrin your clothes or accept Lyme as your new personality.", Category: "Wildlife"},
		{Text: "If you see a moose, the moose has already seen you, judged you, and decided you're not worth charging — yet.", Category: "Wildlife"},
		{Text: "Squirrels at popular campsites have a PhD in food theft. They will outsmart your cooler. Bring a real bear box.", Category: "Wildlife"},
		{Text: "Coyotes howling at night are not coming for you. Coyotes howling next to your tent are absolutely sizing you up.", Category: "Wildlife"},

		// Fire
		{Text: "If your fire-starting plan includes 'I have a lighter,' you have no plan. Bring a backup. Bring two.", Category: "Fire"},
		{Text: "Wet wood doesn't burn. Wet wood with gasoline burns your eyebrows. Find dry wood.", Category: "Fire"},
		{Text: "Build the fire small, feed it slow. A bonfire impresses nobody but your insurance company.", Category: "Fire"},
		{Text: "Drown your fire when you leave. 'Mostly out' is how 40,000 acres became someone's worst day.", Category: "Fire"},
		{Text: "Birch bark lights wet. Pine needles light explosive. Cardboard from your beer twelve-pack is honestly fine too.", Category: "Fire"},
		{Text: "Dryer lint smeared in petroleum jelly is the cheapest tinder on Earth and burns like spite. Pack some.", Category: "Fire"},
		{Text: "Don't burn poison ivy. The smoke gets in your lungs and your eulogy gets really specific.", Category: "Fire"},
		{Text: "If you can't make fire with one match, practice in your driveway. Not at 11 PM in a freezing drizzle with everyone watching.", Category: "Fire"},

		// Gear
		{Text: "Pack three more pairs of socks than feels reasonable. Wet feet at hour two will ruin your marriage faster than therapy can fix it.", Category: "Gear"},
		{Text: "Cotton kills. Wear it and freeze your ass off, or wear synthetics and live to complain another day.", Category: "Gear"},
		{Text: "Your $400 ultralight tent doesn't matter if you packed a 12-pound camp chair. Be honest about your priorities.", Category: "Gear"},
		{Text: "Headlamps over flashlights, always. You have two hands for a reason — neither of them is 'holding a Maglite at the latrine.'", Category: "Gear"},
		{Text: "Duct tape solves 90% of camping problems. The other 10% involve organs and a hospital.", Category: "Gear"},
		{Text: "A tarp over your tent costs $15 and saves your weekend. Skip it and discover what 'condensation' really means.", Category: "Gear"},
		{Text: "Trekking poles aren't for showoffs. They're for your knees, your balance, and the times you really need to look like a hiker.", Category: "Gear"},
		{Text: "Sleeping pad R-value matters more than your sleeping bag rating. Cold seeps from the ground, like resentment.", Category: "Gear"},
		{Text: "Test your gear in the backyard before the trip. Discovering your tent is missing a pole at 9 PM is a special kind of hell.", Category: "Gear"},

		// Weather
		{Text: "Check the forecast. Then check it again. Then assume it's wrong and pack the rain gear anyway.", Category: "Weather"},
		{Text: "Lightning over the ridge means get off the ridge. The view will be there next time. Your dental records won't help much.", Category: "Weather"},
		{Text: "Hypothermia kills people in 50-degree rain, not just blizzards. Stay dry. Stay dry. Stay damn dry.", Category: "Weather"},
		{Text: "Sun at 9,000 feet will cook you through cloud cover. Sunscreen your ears, your scalp, and the back of your neck. You're welcome.", Category: "Weather"},
		{Text: "Wind chill is real. 40 mph gusts at 40 degrees is a hard pass for any tent rated less than 3-season.", Category: "Weather"},
		{Text: "If the sky turns green in tornado country, get to a ditch. Your tent is not a storm shelter, it's a body bag.", Category: "Weather"},

		// Food
		{Text: "Freeze-dried meals taste like a wet sock that someone described 'beef stroganoff' to once. After hour 14 you won't care.", Category: "Food"},
		{Text: "Pre-cook bacon at home and freeze it. Future-you will weep with gratitude on day three.", Category: "Food"},
		{Text: "Whiskey is dense calories and dense morale. Bring more than you think you need, share less than people expect.", Category: "Food"},
		{Text: "Wash dishes 200 feet from camp. Bury food scraps deep or pack them out. Or attract every critter for half a mile, your call.", Category: "Food"},
		{Text: "Tortillas don't get crushed in your pack. Bread does. Choose your carb wisely.", Category: "Food"},
		{Text: "Hot sauce makes everything edible. Pack a bottle. Pack two. The trail does not respect blandness.", Category: "Food"},
		{Text: "Never trust a stream, no matter how sparkling it looks. Filter it, boil it, or pee blood for a week. Your pick.", Category: "Food"},

		// Tent
		{Text: "Set your tent up before sunset. Pitching in the dark turns 'fun trip' into 'why did we marry each other.'", Category: "Tent"},
		{Text: "Stake your tent even when it's calm. Wind shows up at 2 AM and your nylon palace becomes a kite.", Category: "Tent"},
		{Text: "Don't pitch your tent under widow-makers — dead branches that fall when they damn well please. Look up before you commit.", Category: "Tent"},
		{Text: "A small ground tarp under the tent extends the floor's life by years. A big ground tarp catches rain and pools it under you. Tuck it in.", Category: "Tent"},
		{Text: "Sleep with your head uphill. Your blood and your dignity both flow the wrong way otherwise.", Category: "Tent"},
		{Text: "Air out your tent every morning. The damp will haunt your gear for the rest of the trip if you don't.", Category: "Tent"},

		// People
		{Text: "Don't pitch your tent next to the latrine. You will hate yourself, your friends, and humanity by morning two.", Category: "People"},
		{Text: "Quiet hours are 10 PM to 7 AM. Your bluetooth speaker disagrees, but everyone within 200 yards does not.", Category: "People"},
		{Text: "Whoever cooks doesn't clean. Whoever didn't cook can shut up about the cleanup. This is the only contract that matters.", Category: "People"},
		{Text: "If someone in your group hasn't camped before, do not put them in charge of the map. Or the fire. Or anything sharp.", Category: "People"},
		{Text: "Pack out everything you packed in. The 'I'll just leave this orange peel, it's biodegradable' people are why we can't have nice trails.", Category: "People"},

		// General
		{Text: "Tell someone where you're going and when you'll be back. Search and rescue can't read your mind, and your phone has no bars out here.", Category: "General"},
		{Text: "Bring a paper map. Your GPS will die at the worst possible moment, and 'navigating by vibe' has killed more people than bears.", Category: "General"},
		{Text: "Cell service is a lie they tell you in the trip planning. Plan like the phone is a flashlight that texts on weekends.", Category: "General"},
		{Text: "If you forgot it at home, you didn't need it. If you forgot the medication you actually need, turn around. No martyrs.", Category: "General"},
		{Text: "Pee 200 feet from water and trail. Poop in a 6-inch hole 200 feet from everything. This is not optional, this is civilization.", Category: "General"},
	}
}
