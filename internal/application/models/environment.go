package models

import (
	"beholder-api/internal/jet/model"
	"strings"
	"time"
)

type Environment struct {
	ID          int        `json:"id"`
	Description *string    `json:"description,omitempty"`
	Name        string     `json:"name"`
	Tags        []string   `json:"tags,omitempty"`
	BaseURL     string     `json:"base_url"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`

	Sessions *[]*Session `json:"sessions,omitempty"`
}

func EnvironmentFromDataModel(model model.Environments) *Environment {
	Tags := []string{}
	if model.Tags != nil {
		Tags = strings.Split(*model.Tags, ", ")
	}

	var ID int

	if model.ID != nil {
		ID = int(*model.ID)
	}

	return &Environment{
		ID:          ID,
		Description: model.Description,
		Name:        model.Name,
		Tags:        Tags,
		BaseURL:     *model.BaseURL,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

}

func EnvironmentsFromDataModel(models []model.Environments) *[]*Environment {
	environments := []*Environment{}
	for _, model := range models {
		environments = append(environments, EnvironmentFromDataModel(model))
	}
	return &environments
}
