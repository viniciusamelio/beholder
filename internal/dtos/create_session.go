package dtos

import "beholder-api/internal/application/models"

type CreateSessionDto struct {
	EnvUID int    `json:"env_uid" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

func (dto CreateSessionDto) ToModel() models.Session {
	return models.Session{
		EnvUID: dto.EnvUID,
		UserID: dto.UserID,
	}
}
