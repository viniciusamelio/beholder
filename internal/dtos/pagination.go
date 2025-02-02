package dtos

import "github.com/guregu/null"

var DefaultPagination = PaginationDto{
	Take: null.IntFrom(10),
	Skip: null.IntFrom(0),
}

type PaginationDto struct {
	Take null.Int `json:"take"`
	Skip null.Int `json:"skip"`
}
