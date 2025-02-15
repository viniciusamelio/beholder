package models

import (
	"beholder-api/internal/jet/model"
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

func RequestFromDataModel(r model.Requests) Request {
	ID := r.ID
	EnvironmentID := *r.EnvironmentID
	var SessionID int
	if r.SessionID != nil {
		SessionID = int(*r.SessionID)
	}
	return Request{
		ID:            int(*ID),
		EnvironmentID: int(EnvironmentID),
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

func RequestFromDataModelSlice(r []model.Requests) []Request {
	requests := []Request{}

	for i := 0; i < len(r); i++ {
		requests = append(requests, RequestFromDataModel(r[i]))
	}

	return requests

}

func (r *Request) ToModel() model.Requests {
	now := time.Now()
	EnvironmentID := int32(r.EnvironmentID)
	var SessionID int32
	if r.SessionID != nil {
		SessionID = int32(*r.SessionID)
	}
	ID := int32(r.ID)
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
