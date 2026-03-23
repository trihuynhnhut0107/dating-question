package seeder

import (
	"log"

	"dating-question-backend/internal/models"

	"gorm.io/gorm"
)

type categoryInput struct {
	Name        string
	Description string
	Color       string
	Questions   []questionInput
}

type questionInput struct {
	Text       string
	Difficulty models.Difficulty
}

var seedData = []categoryInput{
	{
		Name:        "Icebreaker",
		Description: "Light and fun questions to get the conversation started",
		Color:       "#4CAF50",
		Questions: []questionInput{
			{"What's your go-to karaoke song?", models.DifficultyEasy},
			{"If you could only eat one cuisine for the rest of your life, what would it be?", models.DifficultyEasy},
			{"What's the most useless talent you have?", models.DifficultyEasy},
			{"Would you rather explore space or the deep ocean?", models.DifficultyEasy},
			{"What show are you currently binge watching?", models.DifficultyEasy},
			{"What's a hobby you've always wanted to try?", models.DifficultyEasy},
			{"If you were a kitchen appliance, which one would you be and why?", models.DifficultyEasy},
			{"What's your most controversial food opinion?", models.DifficultyEasy},
		},
	},
	{
		Name:        "Getting to Know You",
		Description: "Questions to understand someone's personality and values",
		Color:       "#2196F3",
		Questions: []questionInput{
			{"What does your ideal weekend look like?", models.DifficultyEasy},
			{"Are you more of an introvert or extrovert, and does it change over time?", models.DifficultyMedium},
			{"What's something you believed as a child that turned out to be completely wrong?", models.DifficultyEasy},
			{"What's a value you hold that most people in your life would agree with?", models.DifficultyMedium},
			{"What's the best piece of advice you've ever received?", models.DifficultyEasy},
			{"How do you usually recharge after a tough week?", models.DifficultyEasy},
			{"What's something small that brings you a lot of joy?", models.DifficultyEasy},
			{"What does success look like to you?", models.DifficultyMedium},
		},
	},
	{
		Name:        "Deep Connection",
		Description: "Meaningful questions to build emotional intimacy",
		Color:       "#9C27B0",
		Questions: []questionInput{
			{"What's a fear you've never told anyone about?", models.DifficultyHard},
			{"What moment in your life changed you the most?", models.DifficultyHard},
			{"What's something you wish people understood about you without having to explain it?", models.DifficultyHard},
			{"What does love mean to you, and has that definition changed?", models.DifficultyHard},
			{"What's a belief you hold that you're still not entirely sure about?", models.DifficultyMedium},
			{"When do you feel most like yourself?", models.DifficultyMedium},
			{"What's something you're still healing from?", models.DifficultyHard},
			{"What would you do differently if you knew nobody would judge you?", models.DifficultyMedium},
		},
	},
	{
		Name:        "Dreams & Ambitions",
		Description: "Questions about goals, dreams, and the future",
		Color:       "#FF9800",
		Questions: []questionInput{
			{"If money weren't a factor, what would you spend your days doing?", models.DifficultyEasy},
			{"Where do you see yourself in 5 years, and is that exciting or scary?", models.DifficultyMedium},
			{"What's a dream you had that you gave up on, and do you regret it?", models.DifficultyHard},
			{"What's one thing you want to accomplish before you're 80?", models.DifficultyEasy},
			{"If you could master any skill overnight, what would it be?", models.DifficultyEasy},
			{"What kind of legacy do you want to leave behind?", models.DifficultyMedium},
			{"Is there a place in the world you feel you belong but haven't lived yet?", models.DifficultyMedium},
			{"What are you most proud of that you rarely talk about?", models.DifficultyMedium},
		},
	},
	{
		Name:        "Relationship & Love",
		Description: "Questions about past experiences and views on relationships",
		Color:       "#F44336",
		Questions: []questionInput{
			{"What's your love language, and has it always been that way?", models.DifficultyEasy},
			{"What's a green flag you always look for in a partner?", models.DifficultyEasy},
			{"What did your most important relationship teach you about yourself?", models.DifficultyHard},
			{"How do you handle conflict - do you tend to avoid it or face it head on?", models.DifficultyMedium},
			{"What's a non-negotiable in a relationship for you?", models.DifficultyMedium},
			{"Do you believe people can truly change? Why or why not?", models.DifficultyMedium},
			{"What does a healthy relationship look like to you?", models.DifficultyMedium},
			{"What's something you need from a partner that you find hard to ask for?", models.DifficultyHard},
		},
	},
}

func Run(db *gorm.DB) {
	for _, cat := range seedData {
		var existing models.Category
		result := db.Where("name = ?", cat.Name).First(&existing)
		if result.Error == nil {
			log.Printf("Skipping category '%s' (already exists)", cat.Name)
			continue
		}

		category := models.Category{
			Name:        cat.Name,
			Description: cat.Description,
			Color:       cat.Color,
		}
		if err := db.Create(&category).Error; err != nil {
			log.Fatalf("Failed to seed category '%s': %v", cat.Name, err)
		}
		log.Printf("Seeded category: %s", category.Name)

		for _, q := range cat.Questions {
			question := models.Question{
				Text:       q.Text,
				CategoryID: category.ID,
				Difficulty: q.Difficulty,
				IsActive:   true,
			}
			if err := db.Create(&question).Error; err != nil {
				log.Fatalf("Failed to seed question: %v", err)
			}
		}
		log.Printf("  Seeded %d questions", len(cat.Questions))
	}
	log.Println("Seeding complete")
}
