package question

import "github.com/google/uuid"

type QuestionDTO struct {
    ID         uuid.UUID `json:"id"`
    Content    string    `json:"content"`
    CategoryID uuid.UUID `json:"category_id"`
}

