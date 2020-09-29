package entity

import "time"

// Expense type
type Expense struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Total      float64   `json:"total"`
	Attachment string    `json:"attachment"`
	CreatedAt  time.Time `json:"createdAt"`
}
