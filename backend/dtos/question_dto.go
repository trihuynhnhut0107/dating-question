package dtos

import "github.com/google/uuid"

type QuestionDTO struct {
    ID         uuid.UUID `json:"id"`
    Content    string    `json:"content"`
    CategoryID uuid.UUID `json:"category_id"`
}
type CategoryDTO struct {
    ID    uuid.UUID `json:"id"`
    Name  string    `json:"name"`
    Icon  string    `json:"icon"`
}
type PaginatedResponse[T any] struct {
    Data       []T    `json:"data"`
    NextCursor string `json:"next_cursor"`
    Limit      int    `json:"limit"`
}