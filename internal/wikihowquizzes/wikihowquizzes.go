package wikihowquizzes

import (
	"math/rand"
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
	r := rand.New(rand.NewSource(int64(seed)))

	return quizzes[r.Intn(len(quizzes))]
}

func getAllQuizzes() []Quiz {
	return []Quiz{
		{Title: "Am I Pretty Quiz", URL: "https://www.wikihow.com/Am-I-Pretty-Quiz"},
		{Title: "Am I Hot Quiz", URL: "https://www.wikihow.com/Am-I-Hot-Quiz"},
		{Title: "Am I Ugly Quiz", URL: "https://www.wikihow.com/Am-I-Ugly-Quiz"},
		{Title: "Am I Depressed Quiz", URL: "https://www.wikihow.com/Am-I-Depressed-Quiz"},
		{Title: "Am I in Love Quiz", URL: "https://www.wikihow.com/Am-I-in-Love-Quiz"},
		{Title: "Am I Boring Quiz", URL: "https://www.wikihow.com/Am-I-Boring-Quiz"},
		{Title: "Am I Lazy Quiz", URL: "https://www.wikihow.com/Am-I-Lazy-Quiz"},
		{Title: "Am I an Introvert or Extrovert Quiz", URL: "https://www.wikihow.com/Am-I-an-Introvert-or-Extrovert-Quiz"},
		{Title: "Am I Smart Quiz", URL: "https://www.wikihow.com/Am-I-Smart-Quiz"},
		{Title: "Am I Toxic Quiz", URL: "https://www.wikihow.com/Am-I-Toxic-Quiz"},
		{Title: "Am I Annoying Quiz", URL: "https://www.wikihow.com/Am-I-Annoying-Quiz"},
		{Title: "Am I an Empath Quiz", URL: "https://www.wikihow.com/Am-I-an-Empath-Quiz"},
		{Title: "Am I a Picky Eater Quiz", URL: "https://www.wikihow.com/Am-I-a-Picky-Eater-Quiz"},
		{Title: "Am I Selfish Quiz", URL: "https://www.wikihow.com/Am-I-Selfish-Quiz"},
		{Title: "Am I Funny Quiz", URL: "https://www.wikihow.com/Am-I-Funny-Quiz"},
		{Title: "Am I a Good Friend Quiz", URL: "https://www.wikihow.com/Am-I-a-Good-Friend-Quiz"},
		{Title: "Am I Bisexual Quiz", URL: "https://www.wikihow.com/Am-I-Bisexual-Quiz"},
		{Title: "Does He Like Me Quiz", URL: "https://www.wikihow.com/Does-He-Like-Me-Quiz"},
		{Title: "Does She Like Me Quiz", URL: "https://www.wikihow.com/Does-She-Like-Me-Quiz"},
		{Title: "Is My Crush into Me Quiz", URL: "https://www.wikihow.com/Is-My-Crush-into-Me-Quiz"},
		{Title: "Will I Ever Get a Boyfriend Quiz", URL: "https://www.wikihow.com/Will-I-Ever-Get-a-Boyfriend-Quiz"},
		{Title: "Will I Ever Get a Girlfriend Quiz", URL: "https://www.wikihow.com/Will-I-Ever-Get-a-Girlfriend-Quiz"},
		{Title: "Should I Break Up with Him Quiz", URL: "https://www.wikihow.com/Should-I-Break-Up-with-Him-Quiz"},
		{Title: "Should I Break Up with Her Quiz", URL: "https://www.wikihow.com/Should-I-Break-Up-with-Her-Quiz"},
		{Title: "What Should I Do with My Life Quiz", URL: "https://www.wikihow.com/What-Should-I-Do-with-My-Life-Quiz"},
		{Title: "What Kind of Friend Are You Quiz", URL: "https://www.wikihow.com/What-Kind-of-Friend-Are-You-Quiz"},
		{Title: "What Is My Aesthetic Quiz", URL: "https://www.wikihow.com/What-Is-My-Aesthetic-Quiz"},
		{Title: "What Should I Eat Quiz", URL: "https://www.wikihow.com/What-Should-I-Eat-Quiz"},
		{Title: "What Is My Spirit Animal Quiz", URL: "https://www.wikihow.com/What-Is-My-Spirit-Animal-Quiz"},
		{Title: "What Is My Personality Type Quiz", URL: "https://www.wikihow.com/What-Is-My-Personality-Type-Quiz"},
		{Title: "What Career Is Right for Me Quiz", URL: "https://www.wikihow.com/What-Career-Is-Right-for-Me-Quiz"},
		{Title: "What Color Should I Dye My Hair Quiz", URL: "https://www.wikihow.com/What-Color-Should-I-Dye-My-Hair-Quiz"},
		{Title: "Which Hogwarts House Am I In Quiz", URL: "https://www.wikihow.com/Which-Hogwarts-House-Am-I-In-Quiz"},
		{Title: "Which Disney Princess Am I Quiz", URL: "https://www.wikihow.com/Which-Disney-Princess-Am-I-Quiz"},
		{Title: "How Mean Am I Quiz", URL: "https://www.wikihow.com/How-Mean-Am-I-Quiz"},
		{Title: "How Old Do I Act Quiz", URL: "https://www.wikihow.com/How-Old-Do-I-Act-Quiz"},
		{Title: "How Smart Am I Quiz", URL: "https://www.wikihow.com/How-Smart-Am-I-Quiz"},
		{Title: "Are You Ready for a Baby Quiz", URL: "https://www.wikihow.com/Are-You-Ready-for-a-Baby-Quiz"},
		{Title: "Should I Get a Dog Quiz", URL: "https://www.wikihow.com/Should-I-Get-a-Dog-Quiz"},
		{Title: "Should I Get a Cat Quiz", URL: "https://www.wikihow.com/Should-I-Get-a-Cat-Quiz"},
		{Title: "Should I Cut My Hair Quiz", URL: "https://www.wikihow.com/Should-I-Cut-My-Hair-Quiz"},
		{Title: "Am I Pretty or Ugly Quiz", URL: "https://www.wikihow.com/Am-I-Pretty-or-Ugly-Quiz"},
	}
}
