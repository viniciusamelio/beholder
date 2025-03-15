package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"beholder-api/schema"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SessionRouter(r *echo.Echo, repo repositories.SessionRepository, taskService services.TaskService) {
	g := r.Group("/session")

	g.GET("", PaginationMiddleware(func(c echo.Context) error {
		context := c.(*PaginationContext)
		repo.Get(context.Pagination).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(s *[]*models.Session) {
				Response(c, 200, s)
			},
		)
		return nil
	}))

	g.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}
		repo.GetByID(int(parsedId)).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(s *models.Session) {

				session := schema.Session{
					Id:            int32(s.ID),
					EnvironmentId: int32(s.EnvironmentID),
					UserId:        s.UserID,
					CreatedAt:     timestamppb.New(*s.CreatedAt),
					UpdatedAt:     timestamppb.New(*s.UpdatedAt),
					Tags:          s.Tags,
				}

				var requests []*schema.Request
				for _, request := range *s.Requests {
					requests = append(requests, &schema.Request{
						Id:            int32(request.ID),
						EnvironmentId: int32(request.EnvironmentID),
						SessionId:     &session.Id,
						UserId:        request.UserID,
						Method:        request.Method,
						Name:          request.Name,
						Path:          request.Path,
						Headers:       request.Headers,
						Body:          request.Body,
						CalledAt:      timestamppb.New(request.CalledAt),
						CreatedAt:     timestamppb.New(request.CreatedAt),
						UpdatedAt:     timestamppb.New(*request.UpdatedAt),
					})
				}

				ProtobufResponse(c, 200, &schema.SessionRequests{
					Session:  &session,
					Requests: requests,
				})
				Response(c, 200, s)
			},
		)
		return nil
	})

	g.POST("", func(c echo.Context) error {
		input := dtos.CreateSessionDto{}
		err := c.Bind(&input)
		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}
		repo.Create(input.ToModel()).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(s *models.Session) {
				Response(c, 201, s)
			},
		)
		return nil
	})

	g.POST("/:id/replay", func(c echo.Context) error {
		id := c.Param("id")
		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}

		repo.GetRequests(int(parsedId)).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(c *dtos.GetRequestsFromSessionResponseDto) {
				taskService.Execute(c)
			},
		)

		return nil
	})
}
