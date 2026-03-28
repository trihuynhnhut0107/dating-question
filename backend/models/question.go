package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Questions []Question `json:"questions"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Question struct{
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Content string `json:"content"`
	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}