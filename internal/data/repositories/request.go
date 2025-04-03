package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/jet/table"
	"beholder-api/internal/utils"
	"database/sql"
	"fmt"
)

type RequestRepository struct {
	db *sql.DB
}

func NewRequestRepository(db *sql.DB) *RequestRepository {
	return &RequestRepository{db: db}
}

func (rr *RequestRepository) Create(request model.Requests) utils.Action[*models.Request] {
	err := table.Requests.INSERT().MODEL(&request).
		RETURNING(table.Requests.AllColumns).Query(rr.db, &request)
	if err != nil {
		code := 400
		fmt.Println(err.Error())
		return utils.NewLeft[utils.Failure, *models.Request](utils.NewUnknownFailure("Error creating request", &code))
	}
	return utils.NewRight[utils.Failure](models.RequestFromDataModel(request))
}

func (rr *RequestRepository) Get(pagination dtos.PaginationDto) utils.Action[*[]*models.Request] {
	dest := []models.FullRequestDataModel{}
	err := table.Requests.SELECT(
		table.Requests.AllColumns,
		table.Environments.AllColumns,
		table.Sessions.AllColumns,
	).
		FROM(
			table.Requests.
				INNER_JOIN(
					table.Environments,
					table.Requests.EnvironmentID.EQ(table.Environments.ID),
				).
				LEFT_JOIN(
					table.Sessions,
					table.Requests.SessionID.EQ(table.Sessions.ID),
				),
		).
		ORDER_BY(table.Requests.CreatedAt.DESC()).
		LIMIT(pagination.Take.Int64).
		OFFSET(pagination.Skip.Int64).
		Query(rr.db, &dest)

	if err != nil {
		fmt.Print(err.Error())
		code := 400
		return utils.NewLeft[utils.Failure, *[]*models.Request](utils.NewUnknownFailure("Error getting requests", &code))
	}

	return utils.NewRight[utils.Failure](models.RequestsFromFullDataModels(dest))
}
