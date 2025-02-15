package dtos

import (
	"beholder-api/internal/jet/model"
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

func (dto *CreateRequestDto) ToModel() model.Requests {
	ID := int32(utils.GenSnowflakeID())
	EnvironmentID := int32(dto.EnvironmentID)
	var SessionID int32
	if dto.SessionID != nil {
		SessionID = int32(*dto.SessionID)
	}
	now := time.Now()
	return model.Requests{
		ID:            &ID,
		Name:          dto.Name,
		EnvironmentID: &EnvironmentID,
		Path:          dto.Path,
		SessionID:     &SessionID,
		Method:        dto.Method,
		Headers:       dto.Headers,
		Body:          dto.Body,
		CalledAt:      *dto.CalledAt,
		CreatedAt:     &now,
	}
}
