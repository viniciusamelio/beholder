package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/gen/som/where"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"context"
	"time"

	"github.com/gosimple/slug"
)

type CallRepository struct {
	ds        services.SomDatasource
	tableName string
}

func NewCallRepository(ds *services.SomDatasource) *CallRepository {
	return &CallRepository{
		ds:        *ds,
		tableName: "call",
	}
}

func (er *CallRepository) Create(call models.Call) utils.Either[utils.Failure, *models.Call] {
	now := time.Now()
	call.CreatedAt = &now
	call.Name = slug.Make(call.Name)
	call.UID = utils.GenSnowflakeID()

	session, err := er.ds.Session().Query().Filter(
		where.Session.UID.Equal(*call.SessionUID),
	).First(context.Background())

	if err != nil {
		code := 404
		return utils.NewLeft[utils.Failure, *models.Call](utils.NewUnknownFailure("session not found", &code))
	}

	call.Session = session
	err = er.ds.Call().Create(context.Background(), &call)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Call](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](&call)
}
