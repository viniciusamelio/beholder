package router

import (
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/resources/templates"
	"beholder-api/internal/services"
	"context"

	"github.com/labstack/echo/v4"
)

func Router(envRepo *repositories.EnvironmentRepository, sessionRepo *repositories.SessionRepository, callRepo *repositories.CallRepository, taskService services.TaskService) {
	r := echo.New()
	EnvironmentRoutes(r, *envRepo)
	SessionRouter(r, *sessionRepo, taskService)
	CallRouter(r, *callRepo)
	r.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	r.Static("/static", "static")
	r.GET("/app/home", func(c echo.Context) error {
		err := templates.Home().Render(context.Background(), c.Response().Writer)
		if err != nil {
			return err
		}
		return nil
	})
	r.Logger.Fatal(r.Start(":8080"))
}
