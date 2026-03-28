package services

import (
	"dating-question/config"
	"dating-question/models"
	"time"

	"github.com/google/uuid"
)

type QuestionFilter struct {
	Cursor string
	Limit int
	CategoryID string
	Search string
	Sort string
}

func GetQuestionsByCursor(f QuestionFilter) ([]models.Question, string, error) {
    var questions []models.Question
    query := config.DB.Model(&models.Question{}).Limit(f.Limit + 1)
    // Dynamic Filter: Category
    if f.CategoryID != "" {
        if parsedID, err := uuid.Parse(f.CategoryID); err == nil {
            query = query.Where("category_id = ?", parsedID)
        }
    }
    // Dynamic Filter: Search
    if f.Search != "" {
        query = query.Where("content ILIKE ?", "%"+f.Search+"%")
    }
    // Dynamic Sorting & Cursor Logic
    isNewest := f.Sort == "newest" || f.Sort == ""
    if isNewest {
        query = query.Order("created_at DESC")
    } else {
        query = query.Order("created_at ASC")
    }
    if f.Cursor != "" {
        if isNewest {
            query = query.Where("created_at < ?", f.Cursor)
        } else {
            query = query.Where("created_at > ?", f.Cursor)
        }
    }
    err := query.Find(&questions).Error
    if err != nil {
        return nil, "", err
    }
    nextCursor := ""
    if len(questions) > f.Limit {
        nextCursor = questions[f.Limit-1].CreatedAt.Format(time.RFC3339Nano)
        questions = questions[:f.Limit]
    }
    return questions, nextCursor, nil
}