package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/resources/templates"
	"beholder-api/internal/utils"

	"github.com/labstack/echo/v4"
)

func FrontendRouter(e *echo.Echo, envRepo repositories.EnvironmentRepository) {

	g := e.Group("/app")

	g.GET("", PaginationMiddleware(func(c echo.Context) error {
		context := c.(*PaginationContext)
		envRepo.Get(context.Pagination).
			Fold(
				func(f utils.Failure) {
					ErrorResponse(c, *f.Code(), f.Message())
				},
				func(e *[]*models.Environment) {
					RenderTempl(c, templates.Home(e))
				},
			)
		return nil
	}))

}
