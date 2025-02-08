package dtos

import (
	"beholder-api/internal/jet/model"
	"strconv"
	"time"
)

type CreateCallDto struct {
	SessionID     *int      `json:"session_id"`
	EnvironmentID *int      `json:"environment_id"`
	Name          string    `json:"name" validate:"required"`
	Method        string    `json:"method" validate:"required"`
	Headers       *string   `json:"headers"`
	Path          *string   `json:"path"`
	Body          *string   `json:"body"`
	QueryParams   *string   `json:"query_params"`
	UserID        *string   `json:"user_id"`
	CalledAt      time.Time `json:"called_at" validate:"required"`
}

func (c *CreateCallDto) ToModel() model.Calls {
	sessionID := strconv.Itoa(*c.SessionID)
	environmentID := strconv.Itoa(*c.EnvironmentID)
	return model.Calls{
		SessionID:     &sessionID,
		EnvironmentID: &environmentID,
		Name:          c.Name,
		Method:        c.Method,
		Path:          c.Path,
		Body:          c.Body,
		QueryParams:   c.QueryParams,
		UserID:        c.UserID,
		CalledAt:      c.CalledAt,
	}
}
