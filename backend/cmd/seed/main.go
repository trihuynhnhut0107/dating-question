package main

import (
	"dating-question/config"
	"dating-question/models"
	"fmt"
	"log"
)

func main() {
    config.LoadConfig()
    config.ConnectDatabase()

    log.Println("Seeding categories...")
    categories := []models.Category{
        {Name: "Deep Conversations"},
        {Name: "Funny & Light"},
        {Name: "Future & Goals"},
        {Name: "Past & Memories"},
        {Name: "Spicy & Bold"},
    }

    for i := range categories {
        config.DB.Where(models.Category{Name: categories[i].Name}).FirstOrCreate(&categories[i])
    }

    log.Println("Seeding 1000 questions...")
    templates := []string{
        "What is your favorite %s memory?",
        "If you could change one thing about your %s, what would it be?",
        "How do you feel about %s in a relationship?",
        "What is your biggest fear regarding %s?",
        "Tell me a story about %s.",
    }
    topics := []string{"childhood", "travel", "work", "family", "dating"}

    var questions []models.Question
    for i := 0; i < 1000; i++ {
        cat := categories[i%len(categories)]
        content := fmt.Sprintf(templates[i%len(templates)], topics[i%len(topics)])
        questions = append(questions, models.Question{
            Content:    fmt.Sprintf("%d. %s", i+1, content),
            CategoryID: cat.ID,
        })
    }

    // Batch insert for performance
    if err := config.DB.CreateInBatches(questions, 100).Error; err != nil {
        log.Fatalf("Failed to seed questions: %v", err)
    }

    log.Println("Seeding completed successfully!")
}
