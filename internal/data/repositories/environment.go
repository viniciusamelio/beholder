package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/gen/som/where"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/jet/table"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"context"
	"database/sql"
	"time"
)

type EnvironmentRepository struct {
	ds        services.SomDatasource
	db        *sql.DB
	tableName string
}

func NewEnvironmentRepository(ds *services.SomDatasource, db *sql.DB) *EnvironmentRepository {
	return &EnvironmentRepository{ds: *ds, tableName: "environment", db: db}
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

func (er *EnvironmentRepository) GetDetailed(id int, pagination dtos.PaginationDto) utils.Either[utils.Failure, *models.Environment] {
	context := context.Background()
	foundItem, err := er.ds.Environment().
		Query().
		Filter(where.Environment.UID.Equal(id)).
		First(context)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	sessions, err := er.ds.Session().
		Query().
		Filter(where.Session.EnvUID.Equal(id)).
		Limit(int(pagination.Take.Int64)).
		Offset(int(pagination.Skip.Int64)).
		Order().
		All(context)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	foundItem.Sessions = &sessions
	return utils.NewRight[utils.Failure](foundItem)
}

func (er *EnvironmentRepository) Create(env model.Environments) utils.Either[utils.Failure, *model.Environments] {
	envModel := model.Environments{}

	err := table.Environments.INSERT(
		table.Environments.ID,
		table.Environments.Name,
		table.Environments.Description,
		table.Environments.Tags,
		table.Environments.BaseURL,
	).VALUES(utils.GenSnowflakeID(), env.Name, "Something cool", env.Tags, env.BaseURL).
		RETURNING(table.Environments.AllColumns).
		Query(er.db, &envModel)

	if err != nil {
		status := 400
		return utils.NewLeft[utils.Failure, *model.Environments](utils.NewUnknownFailure(err.Error(), &status))
	}
	return utils.NewRight[utils.Failure](&envModel)

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
