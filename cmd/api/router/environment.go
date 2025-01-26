package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/utils"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func EnvironmentRoutes(r *echo.Echo, repo repositories.EnvironmentRepository) *echo.Group {
	g := r.Group("/environment")

	g.POST("", func(c echo.Context) error {
		input := dtos.CreateEnvironmentDto{}
		c.Bind(&input)
		validate := validator.New()
		err := validate.Struct(input)
		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}
		repo.Create(input.ToModel()).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, *f.Code(), f.Message())
			},
			func(e *models.Environment) {
				Response(c, 201, e)
			},
		)
		return nil
	})

	g.PUT("/:id", func(c echo.Context) error {
		id := c.Param("id")
		input := dtos.UpdateEnvironmentDto{}
		c.Bind(&input)
		validate := validator.New()
		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ErrorResponse(c, 400, "Invalid id")
		}
		input.ID = int(parsedId)
		err = validate.Struct(input)
		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}
		repo.Update(input.ID, input.ToModel()).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, *f.Code(), f.Message())
			},
			func(e *models.Environment) {
				Response(c, 200, e)
			},
		)
		return nil
	})

	return g
}
