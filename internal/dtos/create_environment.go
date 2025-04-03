package dtos

import (
	"beholder-api/internal/jet/model"
	"beholder-api/internal/utils"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type CreateEnvironmentDto struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	BaseUrl     string   `json:"base_url" validate:"required,min=3,max=100,url"`
	Description string   `json:"description" validate:"max=1000"`
	Tags        []string `json:"tags" validate:"max=10"`
}

func (dto CreateEnvironmentDto) ToModel() model.Environments {
	for i, v := range dto.Tags {
		dto.Tags[i] = slug.Make(v)
	}
	tags := strings.Join(dto.Tags, ", ")
	ID := int32(utils.GenSnowflakeID())
	now := time.Now()
	environment := model.Environments{
		ID:          &ID,
		Name:        slug.Make(dto.Name),
		BaseURL:     &dto.BaseUrl,
		Description: &dto.Description,
		Tags:        &tags,
		CreatedAt:   &now,
	}
	return environment
}
