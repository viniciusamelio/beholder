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
	"strconv"
	"strings"
	"time"

	"github.com/go-jet/jet/v2/sqlite"
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
		table.Environments.ID.ASC(),
	).Query(er.db, &dest)

	if err != nil {
		return utils.NewLeft[utils.Failure, *[]*models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}

	mappedEnvs := []*models.Environment{}

	for _, v := range dest {
		env := models.Environment{
			Name:        v.Name,
			Description: *v.Description,
			Tags:        strings.Split(*v.Tags, ", "),
			BaseUrl:     *v.BaseURL,
		}
		mappedEnvs = append(mappedEnvs, &env)
	}

	return utils.NewRight[utils.Failure](&mappedEnvs)
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

func (er *EnvironmentRepository) Update(ID int, env models.Environment) utils.Either[utils.Failure, *model.Environments] {
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
		&env.BaseUrl,
		strings.Join(env.Tags, ", "),
		time.Now(),
	).
		WHERE(
			table.Environments.ID.EQ(sqlite.String(strconv.Itoa(ID))),
		).
		RETURNING(table.Environments.AllColumns).
		Query(er.db, &dest)

	if err != nil {
		return utils.NewLeft[utils.Failure, *model.Environments](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure](&dest)
}
