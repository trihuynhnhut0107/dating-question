package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UUID             string            `gorm:"type:varchar(36);uniqueIndex;not null" json:"uuid"`
	Name             string            `gorm:"type:varchar(200)"                     json:"name"`
	SessionQuestions []SessionQuestion `gorm:"foreignKey:SessionID"                  json:"session_questions,omitempty"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.UUID == "" {
		s.UUID = uuid.NewString()
	}
	return nil
}
