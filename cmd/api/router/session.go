package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/jet/model"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SessionRouter(r *echo.Echo, repo repositories.SessionRepository, taskService services.TaskService) {
	g := r.Group("/session")

	g.GET("", PaginationMiddleware(func(c echo.Context) error {
		context := c.(*PaginationContext)
		repo.Get(context.Pagination).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(s *[]*model.Sessions) {
				Response(c, 200, s)
			},
		)
		return nil
	}))

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
			func(s *model.Sessions) {
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

		repo.GetCalls(int(parsedId)).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(c []*models.Call) {
				taskService.Execute(c)
			},
		)

		return nil
	})
}
