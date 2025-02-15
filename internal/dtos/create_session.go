package dtos

import (
	"beholder-api/internal/jet/model"
	"beholder-api/internal/utils"
	"strings"
	"time"
)

type CreateSessionDto struct {
	EnvironmentID int      `json:"environment_id" validate:"required"`
	UserID        string   `json:"user_id" validate:"required"`
	Tags          []string `json:"tags" validate:"max=10"`
}

func (dto CreateSessionDto) ToModel() model.Sessions {
	envId := int32(dto.EnvironmentID)
	tags := strings.Join(dto.Tags, ", ")
	id := int32(utils.GenSnowflakeID())
	now := time.Now()
	return model.Sessions{
		ID:            &id,
		EnvironmentID: &envId,
		UserID:        &dto.UserID,
		Tags:          &tags,
		CreatedAt:     &now,
	}
}
