package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/utils"

	"github.com/labstack/echo/v4"
)

func SessionRouter(r *echo.Echo, repo repositories.SessionRepository) {
	g := r.Group("/session")

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
}
