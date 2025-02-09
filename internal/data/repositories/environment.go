package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/jet/table"
	"beholder-api/internal/utils"
	"database/sql"
	"strconv"
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

	mappedEnvs := []*models.Environment{}

	for _, v := range dest {
		ID, _ := strconv.Atoi(*v.ID)
		env := models.Environment{
			ID:          ID,
			Name:        v.Name,
			Description: v.Description,
			Tags:        strings.Split(*v.Tags, ", "),
			BaseURL:     *v.BaseURL,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		mappedEnvs = append(mappedEnvs, &env)
	}

	return utils.NewRight[utils.Failure](&mappedEnvs)
}

func (er *EnvironmentRepository) GetDetailed(id int, pagination dtos.PaginationDto) utils.Either[utils.Failure, *models.Environment] {
	environment := models.Environment{}

	err := table.Environments.SELECT(
		table.Environments.ID,
		table.Environments.Name,
		table.Environments.Description,
		table.Environments.Tags,
		table.Environments.BaseURL,
	).FROM(
		table.Environments.INNER_JOIN(
			table.Sessions,
			table.Sessions.EnvironmentID.EQ(sqlite.String(strconv.Itoa(id))),
		),
	).WHERE(
		table.Environments.ID.EQ(sqlite.String(strconv.Itoa(id))),
	).Query(er.db, &environment)
	if err != nil {
		return utils.NewLeft[utils.Failure, *models.Environment](utils.NewUnknownFailure(err.Error(), nil))
	}
	return utils.NewRight[utils.Failure](&environment)
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
		&env.BaseURL,
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
