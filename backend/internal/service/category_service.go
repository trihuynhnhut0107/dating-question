package service

import (
	"dating-question-backend/internal/models"

	"gorm.io/gorm"
)

type CreateCategoryInput struct {
	Name        string `json:"name"        binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type UpdateCategoryInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Find(&categories).Error
	return categories, err
}

func (s *CategoryService) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := s.db.Preload("Questions", "is_active = ?", true).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *CategoryService) Create(input CreateCategoryInput) (*models.Category, error) {
	color := input.Color
	if color == "" {
		color = "#000000"
	}
	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
		Color:       color,
	}
	err := s.db.Create(&category).Error
	return &category, err
}

func (s *CategoryService) Update(id uint, input UpdateCategoryInput) (*models.Category, error) {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	if input.Name != nil {
		category.Name = *input.Name
	}
	if input.Description != nil {
		category.Description = *input.Description
	}
	if input.Color != nil {
		category.Color = *input.Color
	}
	err := s.db.Save(&category).Error
	return &category, err
}

func (s *CategoryService) Delete(id uint) error {
	return s.db.Delete(&models.Category{}, id).Error
}
