package models

import (
"time"

"gorm.io/gorm"
)

type SessionQuestion struct {
	gorm.Model
	SessionID   uint       `gorm:"not null;index"         json:"session_id"`
	Session     Session    `gorm:"foreignKey:SessionID"   json:"-"`
	QuestionID  uint       `gorm:"not null;index"         json:"question_id"`
	Question    Question   `gorm:"foreignKey:QuestionID"  json:"question,omitempty"`
	IsAnswered  bool       `gorm:"default:false;not null" json:"is_answered"`
	AnsweredAt  *time.Time `gorm:"default:null"           json:"answered_at"`
}
