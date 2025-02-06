package dtos

import (
	"beholder-api/internal/jet/model"
	"strings"

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
	environment := model.Environments{
		Name:        slug.Make(dto.Name),
		BaseURL:     &dto.BaseUrl,
		Description: &dto.Description,
		Tags:        &tags,
	}
	return environment
}
