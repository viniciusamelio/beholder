package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/utils"

	"github.com/labstack/echo/v4"
)

func RequestRouter(e *echo.Echo, repo repositories.RequestRepository) *echo.Group {
	g := e.Group("/request")

	g.POST("", func(c echo.Context) error {
		input := dtos.CreateRequestDto{}
		err := c.Bind(&input)

		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}

		repo.Create(
			input.ToModel(),
		).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, *f.Code(), f.Message())
			},
			func(r *models.Request) {
				Response(c, 201, r)
			},
		)

		return nil
	})

	g.GET("", PaginationMiddleware(func(c echo.Context) error {
		context := c.(*PaginationContext)
		repo.Get(context.Pagination).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, *f.Code(), f.Message())
			},
			func(r *[]*models.Request) {
				Response(c, 200, r)
			},
		)
		return nil
	}))

	return g
}
