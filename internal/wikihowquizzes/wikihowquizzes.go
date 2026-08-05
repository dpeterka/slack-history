package wikihowquizzes

import (
	"github.com/dpeterka/history-slackbot/internal/rotation"
	"time"
)

// Quiz represents a WikiHow quiz
type Quiz struct {
	Title string
	URL   string
}

// GetRandomQuiz returns a random WikiHow quiz using current date as seed
func GetRandomQuiz() Quiz {
	return GetRandomQuizWithSeed(0)
}

// GetRandomQuizWithSeed returns a random WikiHow quiz with optional seed override.
// If seed is 0, uses current date. Otherwise uses the provided seed for testing.
func GetRandomQuizWithSeed(seed int) Quiz {
	quizzes := getAllQuizzes()

	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}

	return quizzes[rotation.PickIndex(len(quizzes), seed)]
}

func getAllQuizzes() []Quiz {
	return []Quiz{
		{Title: "Am I Cool Quiz", URL: "https://www.wikihow.com/Am-I-Cool-Quiz"},
		{Title: "Am I Funny Quiz", URL: "https://www.wikihow.com/Am-I-Funny-Quiz"},
		{Title: "Am I Toxic Quiz", URL: "https://www.wikihow.com/Am-I-Toxic-Quiz"},
		{Title: "Am I In Love Quiz", URL: "https://www.wikihow.com/Am-I-In-Love-Quiz"},
		{Title: "Am I Weird Quiz", URL: "https://www.wikihow.com/Am-I-Weird-Quiz"},
		{Title: "Am I the Problem Quiz", URL: "https://www.wikihow.com/Am-I-the-Problem-Quiz"},
		{Title: "Am I the Jealous Type Quiz", URL: "https://www.wikihow.com/Am-I-the-Jealous-Type-Quiz"},
		{Title: "Am I a People Pleaser Quiz", URL: "https://www.wikihow.com/Am-I-a-People-Pleaser-Quiz"},
		{Title: "Am I a Good Kisser Quiz", URL: "https://www.wikihow.com/Am-I-a-Good-Kisser-Quiz"},
		{Title: "Am I Ohio Quiz", URL: "https://www.wikihow.com/Am-I-Ohio-Quiz"},
		{Title: "Am I More Golden Retriever or Black Cat Quiz", URL: "https://www.wikihow.com/Am-I-More-Golden-Retriever-or-Black-Cat-Quiz"},
		{Title: "Animal Personality Quiz", URL: "https://www.wikihow.com/Animal-Personality-Quiz"},
		{Title: "Apology Language Quiz", URL: "https://www.wikihow.com/Apology-Language-Quiz"},
		{Title: "Attachment Style Quiz", URL: "https://www.wikihow.com/Attachment-Style-Quiz"},
		{Title: "Aura Points Quiz", URL: "https://www.wikihow.com/Aura-Points-Quiz"},
		{Title: "Avatar the Last Airbender Quiz", URL: "https://www.wikihow.com/Avatar-the-Last-Airbender-Quiz"},
		{Title: "Best Friend Quiz", URL: "https://www.wikihow.com/Best-Friend-Quiz"},
		{Title: "Biggest Flaw Quiz", URL: "https://www.wikihow.com/Biggest-Flaw-Quiz"},
		{Title: "Bird Quiz", URL: "https://www.wikihow.com/Bird-Quiz"},
		{Title: "Career Quiz", URL: "https://www.wikihow.com/Career-Quiz"},
		{Title: "Conrad or Jeremiah Quiz", URL: "https://www.wikihow.com/Conrad-or-Jeremiah-Quiz"},
		{Title: "Countries of the World Quiz", URL: "https://www.wikihow.com/Countries-of-the-World-Quiz"},
		{Title: "Dark Triad Quiz", URL: "https://www.wikihow.com/Dark-Triad-Quiz"},
		{Title: "Disney Quiz", URL: "https://www.wikihow.com/Disney-Quiz"},
		{Title: "DnD Race Quiz", URL: "https://www.wikihow.com/DnD-Race-Quiz"},
		{Title: "Do I Have Main Character Energy Quiz", URL: "https://www.wikihow.com/Do-I-Have-Main-Character-Energy-Quiz"},
		{Title: "Do I Have a Crush Quiz", URL: "https://www.wikihow.com/Do-I-Have-a-Crush-Quiz"},
		{Title: "Does He Like Me Quiz", URL: "https://www.wikihow.com/Does-He-Like-Me-Quiz"},
		{Title: "Does She Like Me Quiz", URL: "https://www.wikihow.com/Does-She-Like-Me-Quiz"},
		{Title: "Does My Crush Like Me Quiz", URL: "https://www.wikihow.com/Does-My-Crush-Like-Me-Quiz"},
		{Title: "Element Quiz", URL: "https://www.wikihow.com/Element-Quiz"},
		{Title: "Face Shape Quiz", URL: "https://www.wikihow.com/Face-Shape-Quiz"},
		{Title: "Finish the Lyrics Quiz", URL: "https://www.wikihow.com/Finish-the-Lyrics-Quiz"},
		{Title: "Four Tendencies Quiz", URL: "https://www.wikihow.com/Four-Tendencies-Quiz"},
		{Title: "Friendship Quiz", URL: "https://www.wikihow.com/Friendship-Quiz"},
		{Title: "Future Husband Name Quiz", URL: "https://www.wikihow.com/Future-Husband-Name-Quiz"},
		{Title: "Game of Thrones House Quiz", URL: "https://www.wikihow.com/Game-of-Thrones-House-Quiz"},
		{Title: "Gen Z Slang Quiz", URL: "https://www.wikihow.com/Gen-Z-Slang-Quiz"},
		{Title: "General Knowledge Quiz", URL: "https://www.wikihow.com/General-Knowledge-Quiz"},
		{Title: "Hogwarts House Quiz", URL: "https://www.wikihow.com/Hogwarts-House-Quiz"},
		{Title: "Hottest Feature Quiz", URL: "https://www.wikihow.com/Hottest-Feature-Quiz"},
		{Title: "Mean Girls Quiz", URL: "https://www.wikihow.com/Mean-Girls-Quiz"},
		{Title: "Music Quiz", URL: "https://www.wikihow.com/Music-Quiz"},
		{Title: "Patronus Quiz", URL: "https://www.wikihow.com/Patronus-Quiz"},
		{Title: "Percy Jackson Cabin Quiz", URL: "https://www.wikihow.com/Percy-Jackson-Cabin-Quiz"},
		{Title: "Personality Quiz", URL: "https://www.wikihow.com/Personality-Quiz"},
		{Title: "Rizz Quiz", URL: "https://www.wikihow.com/Rizz-Quiz"},
		{Title: "Skincare Quiz", URL: "https://www.wikihow.com/Skincare-Quiz"},
		{Title: "Social Battery Quiz", URL: "https://www.wikihow.com/Social-Battery-Quiz"},
		{Title: "Sorting Hat Quiz", URL: "https://www.wikihow.com/Sorting-Hat-Quiz"},
		{Title: "Soulmate Quiz", URL: "https://www.wikihow.com/Soulmate-Quiz"},
		{Title: "Soulmate Name Quiz", URL: "https://www.wikihow.com/Soulmate-Name-Quiz"},
		{Title: "Style Quiz", URL: "https://www.wikihow.com/Style-Quiz"},
		{Title: "Taylor Swift Quiz", URL: "https://www.wikihow.com/Taylor-Swift-Quiz"},
		{Title: "Trauma Response Quiz", URL: "https://www.wikihow.com/Trauma-Response-Quiz"},
		{Title: "Trivia Quiz", URL: "https://www.wikihow.com/Trivia-Quiz"},
		{Title: "Twilight Quiz", URL: "https://www.wikihow.com/Twilight-Quiz"},
		{Title: "What Book Should I Read Quiz", URL: "https://www.wikihow.com/What-Book-Should-I-Read-Quiz"},
		{Title: "What Bug Are You Quiz", URL: "https://www.wikihow.com/What-Bug-Are-You-Quiz"},
		{Title: "What Color Should I Dye My Hair Quiz", URL: "https://www.wikihow.com/What-Color-Should-I-Dye-My-Hair-Quiz"},
		{Title: "What Decade Do I Belong In Quiz", URL: "https://www.wikihow.com/What-Decade-Do-I-Belong-In-Quiz"},
	}
}
