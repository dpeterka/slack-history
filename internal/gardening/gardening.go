package gardening

import (
	"math/rand"
	"time"
)

// Tip represents a gardening tip
type Tip struct {
	Text     string
	Category string
}

// GetRandomTip returns a random gardening tip using current date as seed
func GetRandomTip() Tip {
	return GetRandomTipWithSeed(0)
}

// GetRandomTipWithSeed returns a random gardening tip with optional seed override
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomTipWithSeed(seed int) Tip {
	tips := []Tip{
		// Soil & Nutrients
		{Text: "Your soil pH sweet spot is 6.0-7.0 for most vegetables. Test it, or just cross your fingers and blame the weather.", Category: "Soil"},
		{Text: "Nitrogen for leaves, phosphorus for roots, potassium for fruits. It's like a plant protein shake, but dirt-flavored.", Category: "Soil"},
		{Text: "Compost is free fertilizer if you ignore the months of rotting garbage in your backyard. Your neighbors will understand.", Category: "Soil"},
		{Text: "Over-fertilizing burns roots faster than your enthusiasm for gardening will fade. Less is more, grasshopper.", Category: "Soil"},
		{Text: "Organic soil amendments take longer but won't turn your tomatoes into chemical experiments. Patience, young botanist.", Category: "Soil"},
		{Text: "Heavy feeders like tomatoes and peppers will eat through nutrients like teenagers raid a fridge. Feed accordingly.", Category: "Soil"},
		{Text: "Perlite and vermiculite improve drainage. If your soil is basically clay, you're not gardening, you're sculpting.", Category: "Soil"},
		{Text: "Blood meal and bone meal sound metal, work great, and make your garden smell like a crime scene. Worth it.", Category: "Soil"},

		// Watering
		{Text: "Overwatering kills more plants than underwatering. If your soil is a swamp, you're not nurturing, you're drowning.", Category: "Water"},
		{Text: "Water in the morning so leaves dry before nightfall. Evening watering invites fungus like bad decisions invite regret.", Category: "Water"},
		{Text: "Inconsistent watering leads to blossom end rot on tomatoes. Your plants need routine, not your chaos.", Category: "Water"},
		{Text: "Deep, infrequent watering beats shallow, frequent splashing. Train those roots to dig deep or perish.", Category: "Water"},
		{Text: "Mulch retains moisture and regulates soil temperature. It's a blanket for your plants, except less cuddly.", Category: "Water"},
		{Text: "Drip irrigation is efficient but makes you look like you care too much. Hand watering builds character and blisters.", Category: "Water"},
		{Text: "If leaves are wilting at noon but perking up by evening, chill. They're just being dramatic. Check soil first.", Category: "Water"},

		// Light & Environment
		{Text: "Most vegetables want 6-8 hours of direct sun. If your 'garden' is under a tree, you're growing disappointment.", Category: "Light"},
		{Text: "Leggy seedlings mean not enough light. Move them closer to the window or accept your spindly failure children.", Category: "Light"},
		{Text: "South-facing windows get the most light. North-facing is for ferns and broken dreams.", Category: "Light"},
		{Text: "Temperature swings stress plants. If your grow room is a sauna by day and igloo by night, fix your HVAC.", Category: "Environment"},
		{Text: "Humidity over 60% invites mold. Under 40% invites spider mites. The Goldilocks zone is 50-55%. Good luck.", Category: "Environment"},
		{Text: "Air circulation prevents mold and strengthens stems. A fan is cheaper than therapy for powdery mildew trauma.", Category: "Environment"},
		{Text: "Cold nights below 50°F slow growth. Your plants aren't lazy, they're literally freezing. Bring them inside.", Category: "Environment"},

		// Hydroponics - General
		{Text: "Hydroponics grows faster but requires constant babysitting. It's the needy houseguest of gardening methods.", Category: "Hydroponics"},
		{Text: "pH in hydro should stay between 5.5-6.5. Test daily or enjoy your nutrient lockout surprise party.", Category: "Hydroponics"},
		{Text: "Change your nutrient solution every 2 weeks minimum. Stale water breeds algae and disappointment.", Category: "Hydroponics"},
		{Text: "Water temperature in hydro should stay under 72°F. Warmer water holds less oxygen and invites root rot.", Category: "Hydroponics"},
		{Text: "PPM measures nutrient concentration. Too low = slow growth. Too high = nutrient burn. Goldilocks this one or suffer.", Category: "Hydroponics"},
		{Text: "EC meters are more reliable than TDS/PPM meters. If you're serious about hydroponics, invest or stay mediocre.", Category: "Hydroponics"},
		{Text: "Air stones oxygenate roots in deep water culture. No oxygen = root death = crying into your failed investment.", Category: "Hydroponics"},
		{Text: "Hydroponic systems need light-proof reservoirs. Algae loves light and will turn your setup into a swamp.", Category: "Hydroponics"},

		// Hydroponics - Subtle Hints
		{Text: "When plants start flowering, switch from veg nutrients to bloom formula. Nitrogen takes a backseat to phosphorus now.", Category: "Hydroponics"},
		{Text: "Monitor your plants closely during the flowering phase—they're hungrier than usual and prone to deficiencies.", Category: "Hydroponics"},
		{Text: "If you're growing 'tomatoes' indoors year-round, remember: loose lips sink ships. And gardens.", Category: "Hydroponics"},
		{Text: "High-intensity lights dramatically increase yield but also your electric bill. Hope your 'tomatoes' are worth it.", Category: "Hydroponics"},
		{Text: "Carbon filters control odor. If your 'basil' smells like a Phish concert, your neighbors are noticing.", Category: "Hydroponics"},
		{Text: "Flushing plants with plain water 2 weeks before harvest removes excess nutrients. Your 'vegetables' will taste cleaner.", Category: "Hydroponics"},
		{Text: "Some plants benefit from 12/12 light cycles to trigger flowering. Totally normal tomato behavior, obviously.", Category: "Hydroponics"},
		{Text: "Trichome development peaks late in flowering. Use a jeweler's loupe to check for... peak vegetable ripeness.", Category: "Hydroponics"},
		{Text: "Keep your grow room's location private. You never know when someone might want to 'borrow' your rare heirloom tomatoes.", Category: "Hydroponics"},
		{Text: "If your electricity usage suddenly triples, expect questions. 'I'm really into orchids now' is a bold strategy.", Category: "Hydroponics"},
		{Text: "Training techniques like topping and LST increase yields on bushy plants. Your 'peppers' will thank you.", Category: "Hydroponics"},
		{Text: "Drying and curing your harvest properly prevents mold. Hang in dark, cool room with airflow. Yes, even for... tomatoes.", Category: "Hydroponics"},

		// Backyard Vegetables
		{Text: "Tomatoes are 90% water and 10% drama. They'll split if you water irregularly after drought. Consistent moisture or bust.", Category: "Vegetables"},
		{Text: "Companion planting basil with tomatoes deters pests and makes you feel like a genius. It's mostly placebo, but enjoy.", Category: "Vegetables"},
		{Text: "Prune tomato suckers for bigger fruit. Leave them for more tomatoes. Either way, you'll second-guess yourself.", Category: "Vegetables"},
		{Text: "Peppers like heat but not as much as you think. Above 90°F and they drop flowers like bad habits.", Category: "Vegetables"},
		{Text: "Zucchini grows so fast you'll be begging neighbors to take them. Plant one, feed a neighborhood. Plant three, make enemies.", Category: "Vegetables"},
		{Text: "Cucumbers hate inconsistent watering. Bitter cucumbers are your punishment for neglect. Hydrate or die.", Category: "Vegetables"},
		{Text: "Lettuce bolts in heat, turning bitter and useless. Like your attitude in summer traffic, but worse for salads.", Category: "Vegetables"},
		{Text: "Carrots need loose soil. If you're growing them in clay, congrats on your new crop of twisted orange nightmares.", Category: "Vegetables"},
		{Text: "Beans fix nitrogen in soil. They're the only plants that actually help instead of just taking. Plant them, they're heroes.", Category: "Vegetables"},
		{Text: "Squash vine borers will destroy your plants overnight. Check stems weekly or accept your fate as squash borer food supplier.", Category: "Vegetables"},

		// Pests & Problems
		{Text: "Aphids multiply faster than your excuses for not gardening. Spray with neem oil or accept your new insect overlords.", Category: "Pests"},
		{Text: "Spider mites love dry conditions. If you see webbing, you're already losing. Increase humidity and pray.", Category: "Pests"},
		{Text: "Caterpillars eat leaves faster than you can say 'organic gardening.' Handpick them or use BT spray. Your choice.", Category: "Pests"},
		{Text: "Powdery mildew appears when humidity is high and air circulation is low. Your plants now have plant dandruff. You failed them.", Category: "Pests"},
		{Text: "Yellow leaves usually mean nitrogen deficiency. Unless it's overwatering. Or pests. Or root rot. Good luck diagnosing.", Category: "Pests"},
		{Text: "Blossom end rot is calcium deficiency, usually from inconsistent watering. It's not the soil, it's your commitment issues.", Category: "Pests"},
		{Text: "If your whole plant is dying and you don't know why, it's probably root rot from overwatering. You loved it to death.", Category: "Pests"},

		// General Wisdom
		{Text: "Start seeds indoors 6-8 weeks before last frost. Or buy seedlings and pretend you're a gardener. No judgment.", Category: "General"},
		{Text: "Harden off seedlings before transplanting or watch them wilt dramatically like Victorian fainting ladies.", Category: "General"},
		{Text: "Crop rotation prevents soil depletion and disease. Or just dump fertilizer and hope. Modern problems, modern solutions.", Category: "General"},
		{Text: "Harvest vegetables in the morning when they're crispest. Or whenever you remember. They're still edible at noon.", Category: "General"},
		{Text: "Save seeds from heirloom varieties. Hybrids won't grow true to type. Your tomato's children will be disappointing.", Category: "General"},
		{Text: "Gardening teaches patience, humility, and acceptance of failure. It's therapy that sometimes yields tomatoes.", Category: "General"},
		{Text: "The best fertilizer is the gardener's shadow. Meaning: show up regularly or your plants will die. Metaphor solved.", Category: "General"},
		{Text: "No garden is weed-free. Make peace with this truth or spend your life in a losing battle against dandelions.", Category: "General"},
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
		// Soil & Nutrients
		{Text: "Your soil pH sweet spot is 6.0-7.0 for most vegetables. Test it, or just cross your fingers and blame the weather.", Category: "Soil"},
		{Text: "Nitrogen for leaves, phosphorus for roots, potassium for fruits. It's like a plant protein shake, but dirt-flavored.", Category: "Soil"},
		{Text: "Compost is free fertilizer if you ignore the months of rotting garbage in your backyard. Your neighbors will understand.", Category: "Soil"},
		{Text: "Over-fertilizing burns roots faster than your enthusiasm for gardening will fade. Less is more, grasshopper.", Category: "Soil"},
		{Text: "Organic soil amendments take longer but won't turn your tomatoes into chemical experiments. Patience, young botanist.", Category: "Soil"},
		{Text: "Heavy feeders like tomatoes and peppers will eat through nutrients like teenagers raid a fridge. Feed accordingly.", Category: "Soil"},
		{Text: "Perlite and vermiculite improve drainage. If your soil is basically clay, you're not gardening, you're sculpting.", Category: "Soil"},
		{Text: "Blood meal and bone meal sound metal, work great, and make your garden smell like a crime scene. Worth it.", Category: "Soil"},

		// Watering
		{Text: "Overwatering kills more plants than underwatering. If your soil is a swamp, you're not nurturing, you're drowning.", Category: "Water"},
		{Text: "Water in the morning so leaves dry before nightfall. Evening watering invites fungus like bad decisions invite regret.", Category: "Water"},
		{Text: "Inconsistent watering leads to blossom end rot on tomatoes. Your plants need routine, not your chaos.", Category: "Water"},
		{Text: "Deep, infrequent watering beats shallow, frequent splashing. Train those roots to dig deep or perish.", Category: "Water"},
		{Text: "Mulch retains moisture and regulates soil temperature. It's a blanket for your plants, except less cuddly.", Category: "Water"},
		{Text: "Drip irrigation is efficient but makes you look like you care too much. Hand watering builds character and blisters.", Category: "Water"},
		{Text: "If leaves are wilting at noon but perking up by evening, chill. They're just being dramatic. Check soil first.", Category: "Water"},

		// Light & Environment
		{Text: "Most vegetables want 6-8 hours of direct sun. If your 'garden' is under a tree, you're growing disappointment.", Category: "Light"},
		{Text: "Leggy seedlings mean not enough light. Move them closer to the window or accept your spindly failure children.", Category: "Light"},
		{Text: "South-facing windows get the most light. North-facing is for ferns and broken dreams.", Category: "Light"},
		{Text: "Temperature swings stress plants. If your grow room is a sauna by day and igloo by night, fix your HVAC.", Category: "Environment"},
		{Text: "Humidity over 60% invites mold. Under 40% invites spider mites. The Goldilocks zone is 50-55%. Good luck.", Category: "Environment"},
		{Text: "Air circulation prevents mold and strengthens stems. A fan is cheaper than therapy for powdery mildew trauma.", Category: "Environment"},
		{Text: "Cold nights below 50°F slow growth. Your plants aren't lazy, they're literally freezing. Bring them inside.", Category: "Environment"},

		// Hydroponics - General
		{Text: "Hydroponics grows faster but requires constant babysitting. It's the needy houseguest of gardening methods.", Category: "Hydroponics"},
		{Text: "pH in hydro should stay between 5.5-6.5. Test daily or enjoy your nutrient lockout surprise party.", Category: "Hydroponics"},
		{Text: "Change your nutrient solution every 2 weeks minimum. Stale water breeds algae and disappointment.", Category: "Hydroponics"},
		{Text: "Water temperature in hydro should stay under 72°F. Warmer water holds less oxygen and invites root rot.", Category: "Hydroponics"},
		{Text: "PPM measures nutrient concentration. Too low = slow growth. Too high = nutrient burn. Goldilocks this one or suffer.", Category: "Hydroponics"},
		{Text: "EC meters are more reliable than TDS/PPM meters. If you're serious about hydroponics, invest or stay mediocre.", Category: "Hydroponics"},
		{Text: "Air stones oxygenate roots in deep water culture. No oxygen = root death = crying into your failed investment.", Category: "Hydroponics"},
		{Text: "Hydroponic systems need light-proof reservoirs. Algae loves light and will turn your setup into a swamp.", Category: "Hydroponics"},

		// Hydroponics - Subtle Hints
		{Text: "When plants start flowering, switch from veg nutrients to bloom formula. Nitrogen takes a backseat to phosphorus now.", Category: "Hydroponics"},
		{Text: "Monitor your plants closely during the flowering phase—they're hungrier than usual and prone to deficiencies.", Category: "Hydroponics"},
		{Text: "If you're growing 'tomatoes' indoors year-round, remember: loose lips sink ships. And gardens.", Category: "Hydroponics"},
		{Text: "High-intensity lights dramatically increase yield but also your electric bill. Hope your 'tomatoes' are worth it.", Category: "Hydroponics"},
		{Text: "Carbon filters control odor. If your 'basil' smells like a Phish concert, your neighbors are noticing.", Category: "Hydroponics"},
		{Text: "Flushing plants with plain water 2 weeks before harvest removes excess nutrients. Your 'vegetables' will taste cleaner.", Category: "Hydroponics"},
		{Text: "Some plants benefit from 12/12 light cycles to trigger flowering. Totally normal tomato behavior, obviously.", Category: "Hydroponics"},
		{Text: "Trichome development peaks late in flowering. Use a jeweler's loupe to check for... peak vegetable ripeness.", Category: "Hydroponics"},
		{Text: "Keep your grow room's location private. You never know when someone might want to 'borrow' your rare heirloom tomatoes.", Category: "Hydroponics"},
		{Text: "If your electricity usage suddenly triples, expect questions. 'I'm really into orchids now' is a bold strategy.", Category: "Hydroponics"},
		{Text: "Training techniques like topping and LST increase yields on bushy plants. Your 'peppers' will thank you.", Category: "Hydroponics"},
		{Text: "Drying and curing your harvest properly prevents mold. Hang in dark, cool room with airflow. Yes, even for... tomatoes.", Category: "Hydroponics"},

		// Backyard Vegetables
		{Text: "Tomatoes are 90% water and 10% drama. They'll split if you water irregularly after drought. Consistent moisture or bust.", Category: "Vegetables"},
		{Text: "Companion planting basil with tomatoes deters pests and makes you feel like a genius. It's mostly placebo, but enjoy.", Category: "Vegetables"},
		{Text: "Prune tomato suckers for bigger fruit. Leave them for more tomatoes. Either way, you'll second-guess yourself.", Category: "Vegetables"},
		{Text: "Peppers like heat but not as much as you think. Above 90°F and they drop flowers like bad habits.", Category: "Vegetables"},
		{Text: "Zucchini grows so fast you'll be begging neighbors to take them. Plant one, feed a neighborhood. Plant three, make enemies.", Category: "Vegetables"},
		{Text: "Cucumbers hate inconsistent watering. Bitter cucumbers are your punishment for neglect. Hydrate or die.", Category: "Vegetables"},
		{Text: "Lettuce bolts in heat, turning bitter and useless. Like your attitude in summer traffic, but worse for salads.", Category: "Vegetables"},
		{Text: "Carrots need loose soil. If you're growing them in clay, congrats on your new crop of twisted orange nightmares.", Category: "Vegetables"},
		{Text: "Beans fix nitrogen in soil. They're the only plants that actually help instead of just taking. Plant them, they're heroes.", Category: "Vegetables"},
		{Text: "Squash vine borers will destroy your plants overnight. Check stems weekly or accept your fate as squash borer food supplier.", Category: "Vegetables"},

		// Pests & Problems
		{Text: "Aphids multiply faster than your excuses for not gardening. Spray with neem oil or accept your new insect overlords.", Category: "Pests"},
		{Text: "Spider mites love dry conditions. If you see webbing, you're already losing. Increase humidity and pray.", Category: "Pests"},
		{Text: "Caterpillars eat leaves faster than you can say 'organic gardening.' Handpick them or use BT spray. Your choice.", Category: "Pests"},
		{Text: "Powdery mildew appears when humidity is high and air circulation is low. Your plants now have plant dandruff. You failed them.", Category: "Pests"},
		{Text: "Yellow leaves usually mean nitrogen deficiency. Unless it's overwatering. Or pests. Or root rot. Good luck diagnosing.", Category: "Pests"},
		{Text: "Blossom end rot is calcium deficiency, usually from inconsistent watering. It's not the soil, it's your commitment issues.", Category: "Pests"},
		{Text: "If your whole plant is dying and you don't know why, it's probably root rot from overwatering. You loved it to death.", Category: "Pests"},

		// General Wisdom
		{Text: "Start seeds indoors 6-8 weeks before last frost. Or buy seedlings and pretend you're a gardener. No judgment.", Category: "General"},
		{Text: "Harden off seedlings before transplanting or watch them wilt dramatically like Victorian fainting ladies.", Category: "General"},
		{Text: "Crop rotation prevents soil depletion and disease. Or just dump fertilizer and hope. Modern problems, modern solutions.", Category: "General"},
		{Text: "Harvest vegetables in the morning when they're crispest. Or whenever you remember. They're still edible at noon.", Category: "General"},
		{Text: "Save seeds from heirloom varieties. Hybrids won't grow true to type. Your tomato's children will be disappointing.", Category: "General"},
		{Text: "Gardening teaches patience, humility, and acceptance of failure. It's therapy that sometimes yields tomatoes.", Category: "General"},
		{Text: "The best fertilizer is the gardener's shadow. Meaning: show up regularly or your plants will die. Metaphor solved.", Category: "General"},
		{Text: "No garden is weed-free. Make peace with this truth or spend your life in a losing battle against dandelions.", Category: "General"},
	}
}
