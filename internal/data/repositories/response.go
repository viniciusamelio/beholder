package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"context"
	"time"
)

type ResponseRepository struct {
	ds        services.SomDatasource
	tableName string
}

func NewResponseRepository(ds services.SomDatasource) ResponseRepository {
	return ResponseRepository{ds: ds, tableName: "response"}
}

func (er *ResponseRepository) Create(response models.Response) utils.Either[utils.Failure, *models.Response] {
	now := time.Now()
	response.CreatedAt = &now
	response.UID = utils.GenSnowflakeID()
	err := er.ds.Response().Create(context.Background(), &response)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Response](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](&response)
}
