package dtos

import (
	"beholder-api/internal/application/models"
	"time"
)

type CreateCallDto struct {
	SessionUID  *int      `json:"session_uid"`
	EnvUID      *int      `json:"env_uid"`
	Name        string    `json:"name" validate:"required"`
	Method      string    `json:"method" validate:"required"`
	Headers     *string   `json:"headers"`
	Path        *string   `json:"path"`
	Body        *string   `json:"body"`
	QueryParams *string   `json:"query_params"`
	UserID      *string   `json:"user_id"`
	CalledAt    time.Time `json:"called_at" validate:"required"`
}

func (c *CreateCallDto) ToModel() models.Call {
	return models.Call{
		SessionUID:  c.SessionUID,
		EnvUID:      c.EnvUID,
		Name:        c.Name,
		Method:      c.Method,
		Path:        c.Path,
		Body:        c.Body,
		QueryParams: c.QueryParams,
		UserID:      c.UserID,
		CalledAt:    c.CalledAt,
	}
}
