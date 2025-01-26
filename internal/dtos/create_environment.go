package dtos

import (
	"beholder-api/internal/application/models"

	"github.com/gosimple/slug"
)

type CreateEnvironmentDto struct {
	Name    string   `json:"name" validate:"required,min=3,max=100"`
	BaseUrl string   `json:"base_url" validate:"required,min=3,max=100,url"`
	Tags    []string `json:"tags" validate:"max=10"`
}

func (dto CreateEnvironmentDto) ToModel() models.Environment {
	for i, v := range dto.Tags {
		dto.Tags[i] = slug.Make(v)
	}

	environment := models.Environment{
		Name:    slug.Make(dto.Name),
		BaseUrl: dto.BaseUrl,
		Tags:    dto.Tags,
	}
	return environment
}
