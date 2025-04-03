package models

import (
	"beholder-api/internal/jet/model"
	"strings"
	"time"
)

type Session struct {
	ID            int        `json:"id,omitempty"`
	EnvironmentID int        `json:"environment_id,omitempty"`
	UserID        string     `json:"user_id,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Tags          []string   `json:"tags,omitempty"`

	Environment *Environment `json:"environment,omitempty"`
	Requests    *[]*Request  `json:"requests,omitempty"`
}

type FullSessionDataModel struct {
	model.Sessions
	Environment *model.Environments
	Requests    *[]*model.Requests
}

func SessionFromDataModel(model model.Sessions) *Session {

	return &Session{
		ID:            int(*model.ID),
		EnvironmentID: int(*model.EnvironmentID),
		UserID:        *model.UserID,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		Tags:          strings.Split(*model.Tags, ", "),
	}
}

func SessionFromFullDataModel(model FullSessionDataModel) *Session {
	var requests []*Request
	if model.Requests != nil {
		for _, request := range *model.Requests {
			requests = append(requests, RequestFromDataModel(*request))
		}
	}

	return &Session{
		ID:            int(*model.ID),
		EnvironmentID: int(*model.EnvironmentID),
		UserID:        *model.UserID,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		Tags:          strings.Split(*model.Tags, ", "),
		Environment:   EnvironmentFromDataModel(*model.Environment),
		Requests:      &requests,
	}
}

func SessionsFromDataModel(models []*model.Sessions) *[]*Session {
	var sessions []*Session
	for _, model := range models {
		sessions = append(sessions, SessionFromDataModel(*model))
	}

	return &sessions
}

func SessionsFromFullDataModel(models []FullSessionDataModel) *[]*Session {
	var sessions []*Session
	for _, model := range models {
		sessions = append(sessions, SessionFromFullDataModel(model))
	}

	return &sessions
}
