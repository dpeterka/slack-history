package foodtakes

import (
	"time"

	"github.com/dpeterka/history-slackbot/internal/rotation"
)

// Take represents a deliberately controversial (but harmless) food opinion,
// served daily to bait Slack thread arguments.
type Take struct {
	Text     string
	Category string
}

// GetRandomTake returns today's unhinged food take
func GetRandomTake() Take {
	return GetRandomTakeWithSeed(0)
}

// GetRandomTakeWithSeed returns a food take with optional seed override.
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomTakeWithSeed(seed int) Take {
	takes := getAllTakes()

	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	return takes[rotation.PickIndex(len(takes), seed)]
}

func getAllTakes() []Take {
	return []Take{
		// Pizza & Italian
		{Text: "Pineapple on pizza is fine. The real crime is people who order plain cheese at a place with forty toppings. Cowards.", Category: "Pizza"},
		{Text: "Deep dish pizza is a casserole. A delicious casserole, but the moment you need a fork, it stopped being pizza.", Category: "Pizza"},
		{Text: "Cold pizza for breakfast is objectively better than hot pizza for dinner. The flavors have had time to think about what they've done.", Category: "Pizza"},
		{Text: "The crust is the best part of the pizza and people who leave 'pizza bones' on their plate should have to explain themselves to a judge.", Category: "Pizza"},
		{Text: "Spaghetti should be cut with a knife if that's what brings you peace. The Italians can't hear you. You're safe.", Category: "Pasta"},
		{Text: "Mac and cheese from the blue box is its own food, unrelated to real mac and cheese, and both are perfect. Comparing them is a category error.", Category: "Pasta"},

		// Condiments
		{Text: "Ketchup belongs on eggs. If your first instinct is disgust, ask yourself: have you tried it, or were you just raised to fear joy?", Category: "Condiments"},
		{Text: "Mayonnaise is the most versatile condiment on Earth and it gets treated like a war criminal. Justice for mayo.", Category: "Condiments"},
		{Text: "Ranch on pizza is a personality flaw, but honey on pizza is enlightenment. I don't make the rules.", Category: "Condiments"},
		{Text: "Mustard is the only condiment with range: yellow for hot dogs, dijon for dressing, whole grain for showing off. Ketchup has one note and it's 'sweet.'", Category: "Condiments"},
		{Text: "If a sandwich is dry, that's a you problem. There are eleven condiments in your fridge. This was preventable.", Category: "Condiments"},

		// Breakfast
		{Text: "Cereal is soup. Cold soup with milk broth. Sit with that.", Category: "Breakfast"},
		{Text: "Pancakes are just a syrup delivery system, and thick fluffy pancakes are inefficient at their one job. Thin pancakes supremacy.", Category: "Breakfast"},
		{Text: "Brunch is breakfast that costs $24 and requires a reservation. You're paying a cover charge for eggs.", Category: "Breakfast"},
		{Text: "Scrambled eggs should be soft and slightly wet. If your eggs bounce, you've made a rubber product, not breakfast.", Category: "Breakfast"},
		{Text: "Toast is the most underrated food on Earth. Bread, improved by fire, in ninety seconds. Civilization peaked early.", Category: "Breakfast"},

		// Sandwiches & Structure
		{Text: "A hot dog is a sandwich. A taco is a sandwich. A Pop-Tart is a ravioli. Once you accept chaos, the categories stop hurting.", Category: "Food Theory"},
		{Text: "Cutting a sandwich diagonally makes it taste better. This is not psychology, it's physics, and I will not be taking questions.", Category: "Food Theory"},
		{Text: "The bottom bun of a burger does 90% of the structural work for 10% of the credit. It's the bassist of the sandwich world.", Category: "Food Theory"},
		{Text: "Club sandwiches have a third slice of bread for no structural reason. It's bread doing a cameo. Nobody asked for it and nobody stops it.", Category: "Food Theory"},
		{Text: "Wraps are just sandwiches that gave up on themselves.", Category: "Food Theory"},

		// Snacks & Sides
		{Text: "Fries belong in the burger. Not beside it. In it. The people demand integration.", Category: "Snacks"},
		{Text: "A slightly stale chip has better structural integrity and mouthfeel than a fresh one. The 20-minute-open bag is the peak window.", Category: "Snacks"},
		{Text: "Trail mix M&Ms exist to make you eat raisins under false pretenses. It's entrapment.", Category: "Snacks"},
		{Text: "Popcorn is a breakfast food. It's corn. It's fiber. It's basically a deconstructed cereal. Live your truth.", Category: "Snacks"},
		{Text: "The last bite of a banana is a different, worse food than the first bite, and everyone knows it but no one says it.", Category: "Snacks"},

		// Drinks
		{Text: "Iced coffee in winter is not a contradiction, it's commitment. Hot coffee people simply lack conviction.", Category: "Drinks"},
		{Text: "Milk is a deeply weird drink for an adult to order at a restaurant and yet water with 'lemon' is somehow fine. Explain the rules.", Category: "Drinks"},
		{Text: "Pulp-free orange juice is just orange-flavored water for people afraid of texture.", Category: "Drinks"},
		{Text: "La Croix doesn't taste like fruit. It tastes like someone whispered the name of a fruit in another room. That's the appeal.", Category: "Drinks"},

		// Dessert
		{Text: "Cake is usually mediocre and we all pretend otherwise because of the occasion. The frosting is carrying that entire industry.", Category: "Dessert"},
		{Text: "Warm cookie > any cake, at any wedding, ever. Serve cookies at weddings. Break the cycle.", Category: "Dessert"},
		{Text: "Fruit on dessert is a garnish, not a dessert. A fruit plate at the end of a meal is a resignation letter.", Category: "Dessert"},
		{Text: "The edge brownie is the best brownie, and center-piece people are the reason we can't have nice things.", Category: "Dessert"},
		{Text: "Ice cream in a cup with a cone on the side is the correct order. All of the cone experience, none of the drip-based time pressure.", Category: "Dessert"},

		// Kitchen Behavior
		{Text: "Recipes that say '30 minutes' mean 30 minutes for the person who invented the recipe. For you it's an hour and two dishes you didn't plan on.", Category: "Cooking"},
		{Text: "The air fryer is just a small loud oven and it still deserves every bit of the hype. Sometimes marketing tells the truth by accident.", Category: "Cooking"},
		{Text: "Garlic amounts in recipes are a starting bid. 'Two cloves' means six. This is settled law.", Category: "Cooking"},
		{Text: "Leftovers taste better than the original meal roughly 60% of the time, and for curry it's 100%, and this is the strongest argument for meal prep ever made.", Category: "Cooking"},
		{Text: "Washing the dishes as you cook is the highest form of self-respect. Cooking with a full sink is living in the past and the future's mess at once.", Category: "Cooking"},
	}
}
