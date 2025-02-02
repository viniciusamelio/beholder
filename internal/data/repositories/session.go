package repositories

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/dtos"
	"beholder-api/internal/gen/som/where"
	"beholder-api/internal/gen/som/with"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"context"
	"time"
)

type SessionRepository struct {
	ds        services.SomDatasource
	tableName string
}

func NewSessionRepository(ds *services.SomDatasource) *SessionRepository {
	return &SessionRepository{
		ds:        *ds,
		tableName: "session",
	}
}

func (sr *SessionRepository) Create(session models.Session) utils.Either[utils.Failure, *models.Session] {
	session.UID = utils.GenSnowflakeID()
	now := time.Now()
	session.CreatedAt = &now

	env, err := sr.ds.Environment().Query().Filter(
		where.Environment.UID.Equal(session.EnvUID),
	).First(
		context.Background(),
	)

	if err != nil {
		code := 404
		return utils.NewLeft[utils.Failure, *models.Session](utils.NewUnknownFailure("environment not found", &code))
	}
	session.Env = env
	err = sr.ds.Session().Create(context.Background(), &session)
	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, *models.Session](utils.NewUnknownFailure("failed to create session", &code))
	}
	return utils.NewRight[utils.Failure](&session)
}

func (sr *SessionRepository) GetByID(id int) utils.Either[utils.Failure, *models.Session] {
	found, err := sr.ds.Session().Query().Filter(
		where.Session.UID.Equal(
			id,
		),
	).First(
		context.Background(),
	)
	if err != nil {
		code := 404
		return utils.NewLeft[utils.Failure, *models.Session](utils.NewUnknownFailure("failed to get session", &code))
	}

	return utils.NewRight[utils.Failure](found)
}

func (sr *SessionRepository) Get(pagination dtos.PaginationDto) utils.Either[utils.Failure, *[]*models.Session] {
	sessions, err := sr.ds.Session().Query().
		Limit(int(*pagination.Take.Ptr())).
		Offset(int(pagination.Skip.Int64)).
		Fetch(
			with.Session.Env(),
		).
		All(context.Background())
	if err != nil {
		code := 400
		return utils.NewLeft[utils.Failure, *[]*models.Session](utils.NewUnknownFailure("failed to get sessions", &code))
	}

	return utils.NewRight[utils.Failure](&sessions)
}

func (sr *SessionRepository) Delete(id int) utils.Either[utils.Failure, bool] {
	foundSessionOrFailure := sr.GetByID(id)
	var result utils.Either[utils.Failure, bool]

	foundSessionOrFailure.Fold(
		func(f utils.Failure) {
			result = utils.NewLeft[utils.Failure, bool](f)
		},
		func(s *models.Session) {
			err := sr.ds.Session().Delete(context.Background(), s)
			if err != nil {
				code := 400
				result = utils.NewLeft[utils.Failure, bool](utils.NewUnknownFailure("failed to delete session", &code))
			}
			result = utils.NewRight[utils.Failure](true)
		},
	)
	return result

}

func (sr *SessionRepository) GetCalls(id int) utils.Either[utils.Failure, []*models.Call] {
	foundSessionOrFailure := sr.GetByID(id)
	var result utils.Either[utils.Failure, []*models.Call]
	foundSessionOrFailure.Fold(
		func(f utils.Failure) {
			result = utils.NewLeft[utils.Failure, []*models.Call](f)
		},
		func(s *models.Session) {
			calls, err := sr.ds.Call().Query().Filter(
				where.Call.SessionUID.Equal(&id),
			).Fetch(
				with.Call.Session(),
				with.Call.Session().Env(),
			).All(
				context.Background(),
			)
			if err != nil {
				code := 400
				result = utils.NewLeft[utils.Failure, []*models.Call](utils.NewUnknownFailure("failed to get calls", &code))
			}
			result = utils.NewRight[utils.Failure](calls)
		},
	)
	return result
}
