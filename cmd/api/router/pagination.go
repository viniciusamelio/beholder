package router

import (
	"beholder-api/internal/dtos"
	"strconv"

	"github.com/guregu/null"
	"github.com/labstack/echo/v4"
)

type PaginationContext struct {
	echo.Context
	Pagination dtos.PaginationDto
}

func PaginationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		skip, err := strconv.Atoi(c.QueryParam("skip"))
		if err != nil {
			skip = int(dtos.DefaultPagination.Skip.Int64)
		}
		take, err := strconv.Atoi(c.QueryParam("take"))
		if err != nil {
			take = int(dtos.DefaultPagination.Take.Int64)
		}

		if take > 100 {
			return ErrorResponse(c, 400, "Take argument must be at most 100")
		}

		return next(&PaginationContext{
			c,
			dtos.PaginationDto{
				Skip: null.IntFrom(int64(skip)),
				Take: null.IntFrom(int64(take)),
			},
		})
	}
}
