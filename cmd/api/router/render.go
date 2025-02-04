package router

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderTempl(c echo.Context, template templ.Component) error {
	err := template.Render(context.Background(), c.Response().Writer)
	if err != nil {
		return err
	}
	return nil
}
