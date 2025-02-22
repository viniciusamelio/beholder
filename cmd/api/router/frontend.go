package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/resources/templates"
	"beholder-api/internal/resources/views"
	"beholder-api/internal/utils"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FrontendRouter(e *echo.Echo, envRepo repositories.EnvironmentRepository, sessionRepo repositories.SessionRepository) {

	g := e.Group("/app")

	g.GET("", PaginationMiddleware(func(c echo.Context) error {
		context := c.(*PaginationContext)
		envRepo.Get(context.Pagination).
			Fold(
				func(f utils.Failure) {
					ErrorResponse(c, *f.Code(), f.Message())
				},
				func(e *[]*models.Environment) {
					RenderTempl(c, views.Home(e, []templates.BreadcrumbItem{
						{
							Name:   "Environments",
							Active: true,
						},
					}))
				},
			)
		return nil
	}))

	g.GET("/env/:id/sessions", PaginationMiddleware(
		func(c echo.Context) error {
			context := c.(*PaginationContext)
			uid, err := strconv.Atoi(context.Param("id"))
			if err != nil {
				ErrorResponse(c, 400, fmt.Sprintf("Invalid id: %s", context.Param("id")))
			}
			envRepo.GetDetailed(uid, context.Pagination).
				Fold(
					func(f utils.Failure) {
						ErrorResponse(c, *f.Code(), f.Message())
					},
					func(e *models.Environment) {
						envRepo.Get(dtos.DefaultPagination).Fold(
							func(f utils.Failure) {
								ErrorResponse(c, *f.Code(), f.Message())
							},
							func(envs *[]*models.Environment) {
								RenderTempl(c, views.EnvSessions(e, envs, []templates.BreadcrumbItem{
									{
										Name:   "Environments",
										Active: false,
										Href:   "/app",
									},
									{
										Name:   e.Name,
										Active: true,
									},
								}))
							},
						)
					},
				)
			return nil
		},
	))

}
