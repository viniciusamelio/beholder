package router

import (
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/services"

	"github.com/labstack/echo/v4"
)

func Router(envRepo *repositories.EnvironmentRepository, sessionRepo *repositories.SessionRepository, requestRepo *repositories.RequestRepository, taskService services.TaskService) {
	r := echo.New()
	EnvironmentRoutes(r, *envRepo)
	SessionRouter(r, *sessionRepo, taskService)
	FrontendRouter(r, *envRepo, *sessionRepo)
	RequestRouter(r, *requestRepo)
	r.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	r.Static("/static", "static")

	r.Logger.Fatal(r.Start(":8080"))
}
