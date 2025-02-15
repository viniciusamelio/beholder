package services

import "beholder-api/internal/dtos"

type TaskService interface {
	Execute(*dtos.GetRequestsFromSessionResponseDto) error
}
