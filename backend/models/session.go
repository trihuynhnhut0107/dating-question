package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct{
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name *string `json:"name"` 
	Questions []Question `gorm:"many2many:session_questions;" json:"questions"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
