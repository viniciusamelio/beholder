package dtos

import (
	"beholder-api/internal/application/models"

	"github.com/gosimple/slug"
)

type UpdateEnvironmentDto struct {
	ID          int      `json:"id" query:"id" validate:"required"`
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	Description *string  `json:"description" validate:"max=100"`
	BaseUrl     string   `json:"base_url" validate:"required,min=3,max=100,url"`
	Tags        []string `json:"tags" validate:"max=10"`
}

func (u UpdateEnvironmentDto) ToModel() models.Environment {
	for i, v := range u.Tags {
		u.Tags[i] = slug.Make(v)
	}
	environment := models.Environment{
		UID:         u.ID,
		Description: *u.Description,
		Name:        slug.Make(u.Name),
		BaseUrl:     u.BaseUrl,
		Tags:        u.Tags,
	}

	return environment
}
