package dtos

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/utils"
	"time"
)

type CreateRequestDto struct {
	Name          string     `json:"name" validate:"required,min=3,max=100"`
	EnvironmentID int        `json:"environment_id" validate:"required"`
	Path          string     `json:"path" validate:"required"`
	SessionID     *int       `json:"session_id"`
	Method        string     `json:"method" validate:"required, oneof=POST PATCH DELETE GET OPTIONS PUT"`
	Headers       *string    `json:"headers"`
	Body          *string    `json:"body"`
	CalledAt      *time.Time `json:"called_at"`
}

func (dto *CreateRequestDto) ToModel() *models.Request {
	ID := utils.GenSnowflakeID()

	return &models.Request{
		ID:            ID,
		Name:          dto.Name,
		EnvironmentID: dto.EnvironmentID,
		Path:          dto.Path,
		SessionID:     dto.SessionID,
		Method:        dto.Method,
		Headers:       dto.Headers,
		Body:          dto.Body,
		CalledAt:      *dto.CalledAt,
		CreatedAt:     time.Now(),
	}
}
