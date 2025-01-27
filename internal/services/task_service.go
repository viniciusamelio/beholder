package services

import "beholder-api/internal/application/models"

type TaskService interface {
	Execute([]*models.Call) error
}
