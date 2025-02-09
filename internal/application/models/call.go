package models

import "time"

type Call struct {
	ID            int     `json:"id"`
	SessionID     *int    `json:"session_id"`
	EnvironmentID *int    `json:"environment_id"`
	Name          string  `json:"name"`
	Method        string  `json:"method"`
	Path          *string `json:"path"`
	Body          *string `json:"body"`
	UserID        *string `json:"user_id"`
	QueryParams   *string `json:"query_params"`

	CalledAt  time.Time  `json:"called_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
