package dtos

import "beholder-api/internal/application/models"

type GetRequestsFromSessionResponseDto struct {
	BaseURL  string             `json:"base_url"`
	Requests *[]*models.Request `json:"requests"`
}
