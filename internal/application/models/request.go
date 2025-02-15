package models

import (
	"beholder-api/internal/jet/model"
	"beholder-api/internal/utils"
	"strconv"
	"time"

	"github.com/gosimple/slug"
)

type Request struct {
	ID            int        `json:"id"`
	EnvironmentID int        `json:"environment_id"`
	SessionID     *int       `json:"session_id,omitzero"`
	UserID        string     `json:"user_id,omitempty"`
	Method        string     `json:"method"`
	Name          string     `json:"name"`
	Path          string     `json:"path"`
	Headers       *string    `json:"headers,omitempty"`
	Body          *string    `json:"body,omitempty"`
	CalledAt      time.Time  `json:"called_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

func NewRequestFromModel(r model.Requests) Request {
	ID, _ := strconv.Atoi(*r.ID)
	EnvironmentID, _ := strconv.Atoi(*r.EnvironmentID)
	SessionID, _ := strconv.Atoi(*r.SessionID)
	return Request{
		ID:            ID,
		EnvironmentID: EnvironmentID,
		SessionID:     &SessionID,
		UserID:        r.UserID,
		Method:        r.Method,
		Name:          r.Name,
		Path:          r.Path,
		Headers:       r.Headers,
		Body:          r.Body,
		CalledAt:      r.CalledAt,
		CreatedAt:     *r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}

}

func NewRequestFromModelSlice(r []model.Requests) []Request {
	requests := []Request{}

	for i := 0; i < len(r); i++ {
		requests = append(requests, NewRequestFromModel(r[i]))
	}

	return requests

}

func (r *Request) ToModel() model.Requests {
	now := time.Now()
	ID := strconv.Itoa(utils.GenSnowflakeID())
	EnvironmentID := strconv.Itoa(r.EnvironmentID)
	var SessionID string
	if r.SessionID != nil {
		SessionID = strconv.Itoa(*r.SessionID)
	}
	return model.Requests{
		ID:            &ID,
		EnvironmentID: &EnvironmentID,
		SessionID:     &SessionID,
		UserID:        r.UserID,
		Method:        r.Method,
		Name:          slug.Make(r.Name),
		Path:          r.Path,
		Headers:       r.Headers,
		Body:          r.Body,
		CalledAt:      r.CalledAt,
		CreatedAt:     &now,
		UpdatedAt:     r.UpdatedAt,
	}
}
