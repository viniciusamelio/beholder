package dtos

import (
	"beholder-api/internal/jet/model"
	"beholder-api/internal/utils"
	"strconv"
	"strings"
)

type CreateSessionDto struct {
	EnvUID int      `json:"env_uid" validate:"required"`
	UserID string   `json:"user_id" validate:"required"`
	Tags   []string `json:"tags" validate:"max=10"`
}

func (dto CreateSessionDto) ToModel() model.Sessions {
	envId := strconv.Itoa(dto.EnvUID)
	tags := strings.Join(dto.Tags, ", ")
	id := strconv.Itoa(utils.GenSnowflakeID())
	return model.Sessions{
		ID:            &id,
		EnvironmentID: &envId,
		UserID:        &dto.UserID,
		Tags:          &tags,
	}
}
