package dtos

import "beholder-api/internal/jet/model"

type GetCallsFromSessionResponseDto struct {
	BaseURL string
	Calls   *[]*model.Calls
}
