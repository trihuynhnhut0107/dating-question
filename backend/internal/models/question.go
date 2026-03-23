package models

import "gorm.io/gorm"

type Difficulty string

const (
DifficultyEasy   Difficulty = "easy"
DifficultyMedium Difficulty = "medium"
DifficultyHard   Difficulty = "hard"
)

type Question struct {
	gorm.Model
	Text             string     `gorm:"type:text;not null"                        json:"text"`
	CategoryID       uint       `gorm:"not null;index"                            json:"category_id"`
	Category         Category   `gorm:"foreignKey:CategoryID"                     json:"category,omitempty"`
	Difficulty       Difficulty `gorm:"type:varchar(10);default:'easy';not null"  json:"difficulty"`
	IsActive         bool       `gorm:"default:true;not null"                     json:"is_active"`
}
