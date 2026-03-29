package common
type PaginatedResponse[T any] struct {
    Data       []T    `json:"data"`
    NextCursor string `json:"next_cursor"`
    Limit      int    `json:"limit"`
}