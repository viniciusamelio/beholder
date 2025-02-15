package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/jet/table"
	"beholder-api/internal/utils"
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/sqlite"
)

type SessionRepository struct {
	db        *sql.DB
	tableName string
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		tableName: "session",
		db:        db,
	}
}

func (sr *SessionRepository) Create(session model.Sessions) utils.Either[utils.Failure, *models.Session] {
	err := table.Sessions.INSERT().MODEL(session).RETURNING(
		table.Sessions.AllColumns,
	).Query(sr.db, &session)

	if err != nil {
		fmt.Print(err.Error())
		code := 404
		return utils.NewLeft[utils.Failure, *models.Session](utils.NewUnknownFailure("environment not found", &code))
	}

	return utils.NewRight[utils.Failure](models.SessionFromDataModel(session))
}

func (sr *SessionRepository) Get(pagination dtos.PaginationDto) utils.Either[utils.Failure, *[]*models.Session] {
	var dest []models.FullSessionDataModel

	stmt := table.Sessions.SELECT(
		table.Sessions.AllColumns,
		table.Environments.AllColumns,
	).FROM(
		table.Sessions.INNER_JOIN(
			table.Environments,
			table.Sessions.EnvironmentID.EQ(table.Environments.ID),
		),
	).ORDER_BY(table.Sessions.CreatedAt.DESC()).
		OFFSET(pagination.Skip.Int64).
		LIMIT(pagination.Take.Int64)

	err := stmt.Query(sr.db, &dest)

	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, *[]*models.Session](utils.NewUnknownFailure("failed to get sessions", &code))
	}

	return utils.NewRight[utils.Failure](models.SessionsFromFullDataModel(dest))
}

func (sr *SessionRepository) Delete(id int) utils.Either[utils.Failure, bool] {
	_, err := table.Sessions.DELETE().WHERE(
		table.Sessions.ID.EQ(
			sqlite.Int32(int32(id))),
	).Exec(sr.db)

	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, bool](utils.NewUnknownFailure("failed to delete session", &code))
	}
	return utils.NewRight[utils.Failure](true)

}

func (sr *SessionRepository) GetRequests(id int) utils.Either[utils.Failure, *dtos.GetRequestsFromSessionResponseDto] {
	dest := dtos.GetRequestsFromSessionResponseDto{}
	err := table.Requests.SELECT(
		table.Requests.AllColumns.As("requests"),
		table.Environments.BaseURL.AS("base_url"),
	).
		FROM(table.Requests.Table.INNER_JOIN(
			table.Sessions.Table,
			table.Requests.SessionID.EQ(table.Sessions.ID),
		)).
		WHERE(
			table.Requests.SessionID.EQ(
				sqlite.Int32(int32(id))),
		).Query(sr.db, &dest)
	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, *dtos.GetRequestsFromSessionResponseDto](utils.NewUnknownFailure("failed to get calls", &code))
	}

	return utils.NewRight[utils.Failure](&dest)
}
