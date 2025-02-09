package models

import (
	"time"
)

type Environment struct {
	ID          int        `json:"id"`
	Description *string    `json:"description"`
	Name        string     `json:"name"`
	Tags        []string   `json:"tags"`
	BaseURL     string     `json:"base_url"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`

	Sessions *[]*Session `json:"sessions"`
}
