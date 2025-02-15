package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/jet/table"
	"beholder-api/internal/utils"
	"database/sql"
	"fmt"
	"strconv"

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
	dest := model.Sessions{}
	err := table.Sessions.INSERT(
		table.Sessions.ID,
		table.Sessions.EnvironmentID,
		table.Sessions.UserID,
		table.Sessions.Tags,
	).VALUES(
		session.ID,
		session.EnvironmentID,
		session.UserID,
		session.Tags,
	).RETURNING(
		table.Sessions.AllColumns,
	).Query(sr.db, &dest)

	if err != nil {
		fmt.Print(err.Error())
		code := 404
		return utils.NewLeft[utils.Failure, *models.Session](utils.NewUnknownFailure("environment not found", &code))
	}

	return models.SessionFromDataModel(dest)
}

func (sr *SessionRepository) Get(pagination dtos.PaginationDto) utils.Either[utils.Failure, *[]*model.Sessions] {
	dest := []*model.Sessions{}

	err := table.Sessions.SELECT(
		table.Sessions.AllColumns,
	).ORDER_BY(table.Sessions.EnvironmentID.DESC()).
		OFFSET(pagination.Skip.Int64).
		LIMIT(pagination.Take.Int64).
		Query(sr.db, &dest)

	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, *[]*model.Sessions](utils.NewUnknownFailure("failed to get sessions", &code))
	}

	return utils.NewRight[utils.Failure](&dest)
}

func (sr *SessionRepository) Delete(id int) utils.Either[utils.Failure, bool] {
	_, err := table.Sessions.DELETE().WHERE(
		table.Sessions.ID.EQ(
			sqlite.String(strconv.Itoa(id))),
	).Exec(sr.db)

	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, bool](utils.NewUnknownFailure("failed to delete session", &code))
	}
	return utils.NewRight[utils.Failure](true)

}

func (sr *SessionRepository) GetCalls(id int) utils.Either[utils.Failure, *dtos.GetRequestsFromSessionResponseDto] {
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
				sqlite.String(strconv.Itoa(id))),
		).Query(sr.db, &dest)
	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, *dtos.GetRequestsFromSessionResponseDto](utils.NewUnknownFailure("failed to get calls", &code))
	}

	return utils.NewRight[utils.Failure](&dest)
}
