package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/gen/som/where"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"context"
	"time"
)

type EnvironmentRepository struct {
	ds        services.SomDatasource
	tableName string
}

func NewEnvironmentRepository(ds *services.SomDatasource) *EnvironmentRepository {
	return &EnvironmentRepository{ds: *ds, tableName: "environment"}
}

func (er *EnvironmentRepository) Get(pagination dtos.PaginationDto) utils.Either[utils.Failure, *[]*models.Environment] {
	context := context.Background()
	foundItems, err := er.ds.Environment().
		Query().
		Limit(int(pagination.Take.Int64)).
		Offset(int(pagination.Skip.Int64)).
		Order().
		All(context)
	if err != nil {
		return utils.NewLeft[utils.Failure, *[]*models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](&foundItems)
}

func (er *EnvironmentRepository) Create(env models.Environment) utils.Either[utils.Failure, *models.Environment] {
	now := time.Now()
	env.CreatedAt = &now
	env.UID = utils.GenSnowflakeID()
	err := er.ds.Environment().Create(context.Background(), &env)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](&env)
}

func (er *EnvironmentRepository) Update(ID int, env models.Environment) utils.Either[utils.Failure, *models.Environment] {
	context := context.Background()
	foundItem, err := er.ds.Environment().Query().Filter(
		where.Environment.UID.Equal(
			ID,
		),
	).First(context)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	now := time.Now()
	foundItem.BaseUrl = env.BaseUrl
	foundItem.Name = env.Name
	foundItem.UpdatedAt = &now
	foundItem.Tags = env.Tags
	err = er.ds.Environment().Update(context, foundItem)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](foundItem)
}
