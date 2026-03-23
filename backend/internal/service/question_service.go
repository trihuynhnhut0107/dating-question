package service

import (
	"dating-question-backend/internal/models"

	"gorm.io/gorm"
)

type CreateQuestionInput struct {
	Text       string           `json:"text"        binding:"required"`
	CategoryID uint             `json:"category_id" binding:"required"`
	Difficulty models.Difficulty `json:"difficulty"`
	IsActive   *bool            `json:"is_active"`
}

type UpdateQuestionInput struct {
	Text       *string           `json:"text"`
	CategoryID *uint             `json:"category_id"`
	Difficulty *models.Difficulty `json:"difficulty"`
	IsActive   *bool             `json:"is_active"`
}

type ListQuestionsFilter struct {
	CategoryID *uint
	Difficulty *models.Difficulty
	IsActive   *bool
}

type QuestionService struct {
	db *gorm.DB
}

func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{db: db}
}

func (s *QuestionService) GetAll(filter ListQuestionsFilter) ([]models.Question, error) {
	var questions []models.Question
	q := s.db.Preload("Category")
	if filter.CategoryID != nil {
		q = q.Where("category_id = ?", *filter.CategoryID)
	}
	if filter.Difficulty != nil {
		q = q.Where("difficulty = ?", *filter.Difficulty)
	}
	if filter.IsActive != nil {
		q = q.Where("is_active = ?", *filter.IsActive)
	}
	err := q.Find(&questions).Error
	return questions, err
}

func (s *QuestionService) GetByID(id uint) (*models.Question, error) {
	var question models.Question
	err := s.db.Preload("Category").First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (s *QuestionService) Create(input CreateQuestionInput) (*models.Question, error) {
	difficulty := input.Difficulty
	if difficulty == "" {
		difficulty = models.DifficultyEasy
	}
	isActive := true
	if input.IsActive != nil {
		isActive = *input.IsActive
	}
	question := models.Question{
		Text:       input.Text,
		CategoryID: input.CategoryID,
		Difficulty: difficulty,
		IsActive:   isActive,
	}
	err := s.db.Create(&question).Error
	return &question, err
}

func (s *QuestionService) Update(id uint, input UpdateQuestionInput) (*models.Question, error) {
	var question models.Question
	if err := s.db.First(&question, id).Error; err != nil {
		return nil, err
	}
	if input.Text != nil {
		question.Text = *input.Text
	}
	if input.CategoryID != nil {
		question.CategoryID = *input.CategoryID
	}
	if input.Difficulty != nil {
		question.Difficulty = *input.Difficulty
	}
	if input.IsActive != nil {
		question.IsActive = *input.IsActive
	}
	err := s.db.Save(&question).Error
	return &question, err
}

func (s *QuestionService) Delete(id uint) error {
	return s.db.Delete(&models.Question{}, id).Error
}

func (s *QuestionService) GetRandom(count int, categoryID *uint, difficulty *models.Difficulty) ([]models.Question, error) {
	var questions []models.Question
	q := s.db.Preload("Category").Where("is_active = ?", true)
	if categoryID != nil {
		q = q.Where("category_id = ?", *categoryID)
	}
	if difficulty != nil {
		q = q.Where("difficulty = ?", *difficulty)
	}
	err := q.Order("RANDOM()").Limit(count).Find(&questions).Error
	return questions, err
}
