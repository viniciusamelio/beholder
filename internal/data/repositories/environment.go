package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/jet/table"
	"beholder-api/internal/utils"
	"database/sql"
	"strings"
	"time"

	"github.com/go-jet/jet/v2/sqlite"
)

type EnvironmentRepository struct {
	db        *sql.DB
	tableName string
}

func NewEnvironmentRepository(db *sql.DB) *EnvironmentRepository {
	return &EnvironmentRepository{tableName: "environment", db: db}
}

func (er *EnvironmentRepository) Get(pagination dtos.PaginationDto) utils.Either[utils.Failure, *[]*models.Environment] {
	dest := []model.Environments{}

	err := table.Environments.SELECT(
		table.Environments.ID,
		table.Environments.Name,
		table.Environments.Description,
		table.Environments.Tags,
		table.Environments.BaseURL,
	).FROM(
		table.Environments,
	).LIMIT(
		pagination.Take.Int64,
	).OFFSET(
		pagination.Skip.Int64,
	).ORDER_BY(
		table.Environments.ID.DESC(),
	).Query(er.db, &dest)

	if err != nil {
		return utils.NewLeft[utils.Failure, *[]*models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure](models.EnvironmentsFromDataModel(dest))
}

func (er *EnvironmentRepository) GetDetailed(id int, pagination dtos.PaginationDto) utils.Either[utils.Failure, *models.Environment] {
	environment := model.Environments{}

	err := table.Environments.SELECT(
		table.Environments.ID,
		table.Environments.Name,
		table.Environments.Description,
		table.Environments.Tags,
		table.Environments.BaseURL,
	).FROM(
		table.Environments.INNER_JOIN(
			table.Sessions,
			table.Sessions.EnvironmentID.EQ(sqlite.Int32(int32(id))),
		),
	).WHERE(
		table.Environments.ID.EQ(sqlite.Int32(int32(id))),
	).Query(er.db, &environment)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](models.EnvironmentFromDataModel(environment))
}

func (er *EnvironmentRepository) Create(env model.Environments) utils.Action[*models.Environment] {
	err := table.Environments.INSERT().MODEL(env).
		RETURNING(table.Environments.AllColumns).
		Query(er.db, &env)

	if err != nil {
		status := 400
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), &status))
	}
	return utils.NewRight[utils.Failure](models.EnvironmentFromDataModel(env))

}

func (er *EnvironmentRepository) Update(ID int, env models.Environment) utils.Either[utils.Failure, *models.Environment] {
	dest := model.Environments{}
	err := table.Environments.UPDATE(
		table.Environments.Name,
		table.Environments.Description,
		table.Environments.BaseURL,
		table.Environments.Tags,
		table.Environments.UpdatedAt,
	).SET(
		&env.Name,
		&env.Description,
		&env.BaseURL,
		strings.Join(env.Tags, ", "),
		time.Now(),
	).
		WHERE(
			table.Environments.ID.EQ(sqlite.Int32(int32(ID))),
		).
		RETURNING(table.Environments.AllColumns).
		Query(er.db, &dest)

	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure](models.EnvironmentFromDataModel(dest))
}

func (er *EnvironmentRepository) GetSessions(ID int, pagination dtos.PaginationDto) utils.Action[*[]*models.Session] {
	dest := []*model.Sessions{}
	err := table.Sessions.SELECT(
		table.Sessions.AllColumns,
	).FROM(
		table.Sessions.INNER_JOIN(
			table.Environments,
			table.Environments.ID.EQ(table.Sessions.EnvironmentID),
		),
	).WHERE(
		table.Environments.ID.EQ(sqlite.Int32(int32(ID))),
	).LIMIT(
		pagination.Take.Int64,
	).OFFSET(
		pagination.Skip.Int64,
	).ORDER_BY(
		table.Sessions.CreatedAt.DESC(),
	).Query(er.db, &dest)

	if err != nil {
		return utils.NewLeft[utils.Failure, *[]*models.Session](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure](models.SessionsFromDataModel(dest))
}

func (er *EnvironmentRepository) GetRequests(ID int, pagination dtos.PaginationDto) utils.Action[*[]*models.Request] {
	dest := []model.Requests{}
	err := table.Requests.SELECT(
		table.Requests.AllColumns,
	).FROM(
		table.Requests.INNER_JOIN(
			table.Sessions,
			table.Sessions.ID.EQ(table.Requests.SessionID),
		).INNER_JOIN(
			table.Environments,
			table.Environments.ID.EQ(table.Sessions.EnvironmentID),
		),
	).WHERE(
		table.Environments.ID.EQ(sqlite.Int32(int32(ID))),
	).LIMIT(
		pagination.Take.Int64,
	).OFFSET(
		pagination.Skip.Int64,
	).ORDER_BY(
		table.Requests.ID.DESC(),
	).Query(er.db, &dest)

	if err != nil {
		return utils.NewLeft[utils.Failure, *[]*models.Request](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure](models.RequestsFromDataModels(dest))
}
