package hottub

import (
	"math/rand"
	"time"
)

// Tip represents a hot tub care tip
type Tip struct {
	Text     string
	Category string
}

// GetRandomTip returns a random hot tub tip using current date as seed
func GetRandomTip() Tip {
	return GetRandomTipWithSeed(0)
}

// GetRandomTipWithSeed returns a random hot tub tip with optional seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomTipWithSeed(seed int) Tip {
	tips := []Tip{
		// Water Chemistry
		{Text: "Test your hot tub water at least twice a week—your pH doesn't care about your busy schedule.", Category: "Chemistry"},
		{Text: "Shocking your hot tub after heavy use isn't optional, it's a public service to your future guests.", Category: "Chemistry"},
		{Text: "Cloudy water means chemistry's off or your filter's crying for help. Probably both.", Category: "Chemistry"},
		{Text: "Alkalinity between 80-120 ppm keeps pH stable. Think of it as the hot tub's emotional support system.", Category: "Chemistry"},
		{Text: "Calcium hardness matters: too low and your equipment corrodes, too high and you're soaking in scale soup.", Category: "Chemistry"},
		{Text: "Bromine is gentler than chlorine at hot temperatures, but either way, don't drink the hot tub water.", Category: "Chemistry"},
		{Text: "Keep sanitizer levels consistent—bacteria don't take days off, and neither should your maintenance routine.", Category: "Chemistry"},
		{Text: "Your test strips aren't lying. If all the readings look wrong, they're probably all right. Fix it.", Category: "Chemistry"},
		{Text: "Low sanitizer? You're basically marinating in a human tea. Don't be that person.", Category: "Chemistry"},
		{Text: "pH below 7.2? Your hot tub is now an acid bath corroding everything. pH above 7.8? Scale city. Sweet spot: 7.4-7.6.", Category: "Chemistry"},

		// Filter Maintenance
		{Text: "Rinse your filter weekly, deep clean monthly, replace yearly. This is not a suggestion, it's survival.", Category: "Filters"},
		{Text: "A clogged filter makes your pump work harder than a personal trainer on commission. Show it mercy.", Category: "Filters"},
		{Text: "Spray your filter with a hose, not a pressure washer—you're cleaning it, not exorcising demons.", Category: "Filters"},
		{Text: "Overnight filter soaks in filter cleaner are like a spa day for your spa. Don't skip them.", Category: "Filters"},
		{Text: "Keep a spare filter cartridge. When one's soaking, the other's working. Tag team filtration.", Category: "Filters"},
		{Text: "If your water looks like a swamp and your filter's clean, congrats—it's probably your water chemistry.", Category: "Filters"},
		{Text: "A $50 filter cartridge lasts a year. Skipping replacement costs you way more in chemicals and headaches.", Category: "Filters"},
		{Text: "Check your filter basket weekly for leaves, hair ties, and the dreams you've been avoiding.", Category: "Filters"},

		// Water Changes
		{Text: "Drain and refill your hot tub every 3-4 months. Fresh starts aren't just for January.", Category: "Water Changes"},
		{Text: "Your hot tub holds dissolved solids like grudges—eventually you need to let them go and start fresh.", Category: "Water Changes"},
		{Text: "Before refilling, wipe down the shell and rinse the pipes. Out of sight, out of mind leads to biofilm buildup.", Category: "Water Changes"},
		{Text: "Use a submersible pump for draining—waiting for gravity is like watching paint dry, but wetter.", Category: "Water Changes"},
		{Text: "Refill through the filter housing to avoid airlocks. Your pump will thank you by not screaming.", Category: "Water Changes"},
		{Text: "After refilling, balance chemistry within 24 hours. Procrastination breeds algae and regret.", Category: "Water Changes"},
		{Text: "Hot tub cover off while draining? Enjoy the free leaf collection and extra cleaning time.", Category: "Water Changes"},

		// Cover Care
		{Text: "Your hot tub cover isn't a diving board, trampoline, or snow storage unit. Treat it with respect.", Category: "Covers"},
		{Text: "Wipe down your cover monthly with vinyl cleaner. UV rays are harsh, and so is neglect.", Category: "Covers"},
		{Text: "A waterlogged cover weighs as much as your regrets. Replace it before your back does the math.", Category: "Covers"},
		{Text: "Cover locks keep kids, pets, and drunk friends out. Use them, or explain to your insurance why you didn't.", Category: "Covers"},
		{Text: "Flip your cover weekly to prevent one side from becoming the 'warped, saggy' side.", Category: "Covers"},
		{Text: "Condition the vinyl twice a year. Think of it as moisturizer for your hot tub's hat.", Category: "Covers"},
		{Text: "A cracked cover leaks heat faster than gossip at a book club. Fix it or buy a new one.", Category: "Covers"},
		{Text: "Never drag your cover—lift and fold. This isn't a sleeping bag, it's a several-hundred-dollar investment.", Category: "Covers"},

		// Jets and Pumps
		{Text: "Low jet pressure? Check your filter first, then your water level, then question your life choices.", Category: "Equipment"},
		{Text: "Air in the lines? Loosen the union at the pump, let it burp out the air, tighten it back up. Spa baby needs burping.", Category: "Equipment"},
		{Text: "Pump making weird noises? It's not trying to communicate, it's crying for help. Listen to it.", Category: "Equipment"},
		{Text: "Winterize your hot tub properly or risk frozen pipes and a very expensive spring surprise.", Category: "Equipment"},
		{Text: "Clean your jets twice a year by running a line flush product. Biofilm buildup is real and it's gross.", Category: "Equipment"},
		{Text: "Your circulation pump runs 24/7. It's the unsung hero. Keep it happy with clean filters and balanced water.", Category: "Equipment"},
		{Text: "Jet nozzles rotate for a reason. If they don't, they're clogged with calcium or broken dreams. Soak in vinegar.", Category: "Equipment"},
		{Text: "Pump leaking? Don't ignore it. Water and electricity are frenemies at best. Call a professional.", Category: "Equipment"},

		// General Wisdom
		{Text: "Shower before soaking. Your hot tub doesn't want to meet your day's dirt, oils, and existential baggage.", Category: "General"},
		{Text: "Hot tub not a laundry machine: bathing suits with fabric softener = foam party. Not the fun kind.", Category: "General"},
		{Text: "Keep water temp at 104°F max. Higher temps don't make you more relaxed, just more cooked.", Category: "General"},
		{Text: "A hot tub logbook sounds nerdy but saves you from chemistry chaos. Track it or regret it.", Category: "General"},
		{Text: "Leaves, bugs, and rain bring organics that feed algae. Cover up when not in use, it's that simple.", Category: "General"},
		{Text: "Hard water problems? A pre-filter on your garden hose during fills prevents scale buildup before it starts.", Category: "General"},
		{Text: "Foam on the surface? Too many dissolved solids or body oils. Antifoam is a bandaid—drain and refill is the cure.", Category: "General"},
		{Text: "Hot tubs are like relationships: neglect the small stuff and eventually everything falls apart loudly.", Category: "General"},
		{Text: "Read your owner's manual at least once. Yes, it's boring. So is replacing parts you broke by guessing.", Category: "General"},
		{Text: "Invest in good chemicals, not the cheap stuff. Your hot tub's health isn't the place to cut corners.", Category: "General"},
	}

	// Use provided seed or current date for seed
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}
	r := rand.New(rand.NewSource(int64(seed)))

	return tips[r.Intn(len(tips))]
}

// GetRandomTipByCategory returns a random tip from a specific category
func GetRandomTipByCategory(category string) Tip {
	var filtered []Tip
	allTips := getAllTips()

	for _, tip := range allTips {
		if tip.Category == category {
			filtered = append(filtered, tip)
		}
	}

	if len(filtered) == 0 {
		return GetRandomTip()
	}

	// Use current date for seed
	now := time.Now()
	seed := now.Year()*10000 + int(now.Month())*100 + now.Day()
	r := rand.New(rand.NewSource(int64(seed)))

	return filtered[r.Intn(len(filtered))]
}

// getAllTips returns all tips (used internally)
func getAllTips() []Tip {
	return []Tip{
		// Water Chemistry
		{Text: "Test your hot tub water at least twice a week—your pH doesn't care about your busy schedule.", Category: "Chemistry"},
		{Text: "Shocking your hot tub after heavy use isn't optional, it's a public service to your future guests.", Category: "Chemistry"},
		{Text: "Cloudy water means chemistry's off or your filter's crying for help. Probably both.", Category: "Chemistry"},
		{Text: "Alkalinity between 80-120 ppm keeps pH stable. Think of it as the hot tub's emotional support system.", Category: "Chemistry"},
		{Text: "Calcium hardness matters: too low and your equipment corrodes, too high and you're soaking in scale soup.", Category: "Chemistry"},
		{Text: "Bromine is gentler than chlorine at hot temperatures, but either way, don't drink the hot tub water.", Category: "Chemistry"},
		{Text: "Keep sanitizer levels consistent—bacteria don't take days off, and neither should your maintenance routine.", Category: "Chemistry"},
		{Text: "Your test strips aren't lying. If all the readings look wrong, they're probably all right. Fix it.", Category: "Chemistry"},
		{Text: "Low sanitizer? You're basically marinating in a human tea. Don't be that person.", Category: "Chemistry"},
		{Text: "pH below 7.2? Your hot tub is now an acid bath corroding everything. pH above 7.8? Scale city. Sweet spot: 7.4-7.6.", Category: "Chemistry"},

		// Filter Maintenance
		{Text: "Rinse your filter weekly, deep clean monthly, replace yearly. This is not a suggestion, it's survival.", Category: "Filters"},
		{Text: "A clogged filter makes your pump work harder than a personal trainer on commission. Show it mercy.", Category: "Filters"},
		{Text: "Spray your filter with a hose, not a pressure washer—you're cleaning it, not exorcising demons.", Category: "Filters"},
		{Text: "Overnight filter soaks in filter cleaner are like a spa day for your spa. Don't skip them.", Category: "Filters"},
		{Text: "Keep a spare filter cartridge. When one's soaking, the other's working. Tag team filtration.", Category: "Filters"},
		{Text: "If your water looks like a swamp and your filter's clean, congrats—it's probably your water chemistry.", Category: "Filters"},
		{Text: "A $50 filter cartridge lasts a year. Skipping replacement costs you way more in chemicals and headaches.", Category: "Filters"},
		{Text: "Check your filter basket weekly for leaves, hair ties, and the dreams you've been avoiding.", Category: "Filters"},

		// Water Changes
		{Text: "Drain and refill your hot tub every 3-4 months. Fresh starts aren't just for January.", Category: "Water Changes"},
		{Text: "Your hot tub holds dissolved solids like grudges—eventually you need to let them go and start fresh.", Category: "Water Changes"},
		{Text: "Before refilling, wipe down the shell and rinse the pipes. Out of sight, out of mind leads to biofilm buildup.", Category: "Water Changes"},
		{Text: "Use a submersible pump for draining—waiting for gravity is like watching paint dry, but wetter.", Category: "Water Changes"},
		{Text: "Refill through the filter housing to avoid airlocks. Your pump will thank you by not screaming.", Category: "Water Changes"},
		{Text: "After refilling, balance chemistry within 24 hours. Procrastination breeds algae and regret.", Category: "Water Changes"},
		{Text: "Hot tub cover off while draining? Enjoy the free leaf collection and extra cleaning time.", Category: "Water Changes"},

		// Cover Care
		{Text: "Your hot tub cover isn't a diving board, trampoline, or snow storage unit. Treat it with respect.", Category: "Covers"},
		{Text: "Wipe down your cover monthly with vinyl cleaner. UV rays are harsh, and so is neglect.", Category: "Covers"},
		{Text: "A waterlogged cover weighs as much as your regrets. Replace it before your back does the math.", Category: "Covers"},
		{Text: "Cover locks keep kids, pets, and drunk friends out. Use them, or explain to your insurance why you didn't.", Category: "Covers"},
		{Text: "Flip your cover weekly to prevent one side from becoming the 'warped, saggy' side.", Category: "Covers"},
		{Text: "Condition the vinyl twice a year. Think of it as moisturizer for your hot tub's hat.", Category: "Covers"},
		{Text: "A cracked cover leaks heat faster than gossip at a book club. Fix it or buy a new one.", Category: "Covers"},
		{Text: "Never drag your cover—lift and fold. This isn't a sleeping bag, it's a several-hundred-dollar investment.", Category: "Covers"},

		// Jets and Pumps
		{Text: "Low jet pressure? Check your filter first, then your water level, then question your life choices.", Category: "Equipment"},
		{Text: "Air in the lines? Loosen the union at the pump, let it burp out the air, tighten it back up. Spa baby needs burping.", Category: "Equipment"},
		{Text: "Pump making weird noises? It's not trying to communicate, it's crying for help. Listen to it.", Category: "Equipment"},
		{Text: "Winterize your hot tub properly or risk frozen pipes and a very expensive spring surprise.", Category: "Equipment"},
		{Text: "Clean your jets twice a year by running a line flush product. Biofilm buildup is real and it's gross.", Category: "Equipment"},
		{Text: "Your circulation pump runs 24/7. It's the unsung hero. Keep it happy with clean filters and balanced water.", Category: "Equipment"},
		{Text: "Jet nozzles rotate for a reason. If they don't, they're clogged with calcium or broken dreams. Soak in vinegar.", Category: "Equipment"},
		{Text: "Pump leaking? Don't ignore it. Water and electricity are frenemies at best. Call a professional.", Category: "Equipment"},

		// General Wisdom
		{Text: "Shower before soaking. Your hot tub doesn't want to meet your day's dirt, oils, and existential baggage.", Category: "General"},
		{Text: "Hot tub not a laundry machine: bathing suits with fabric softener = foam party. Not the fun kind.", Category: "General"},
		{Text: "Keep water temp at 104°F max. Higher temps don't make you more relaxed, just more cooked.", Category: "General"},
		{Text: "A hot tub logbook sounds nerdy but saves you from chemistry chaos. Track it or regret it.", Category: "General"},
		{Text: "Leaves, bugs, and rain bring organics that feed algae. Cover up when not in use, it's that simple.", Category: "General"},
		{Text: "Hard water problems? A pre-filter on your garden hose during fills prevents scale buildup before it starts.", Category: "General"},
		{Text: "Foam on the surface? Too many dissolved solids or body oils. Antifoam is a bandaid—drain and refill is the cure.", Category: "General"},
		{Text: "Hot tubs are like relationships: neglect the small stuff and eventually everything falls apart loudly.", Category: "General"},
		{Text: "Read your owner's manual at least once. Yes, it's boring. So is replacing parts you broke by guessing.", Category: "General"},
		{Text: "Invest in good chemicals, not the cheap stuff. Your hot tub's health isn't the place to cut corners.", Category: "General"},
	}
}
