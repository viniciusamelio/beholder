// Code generated by github.com/go-surreal/som, DO NOT EDIT.
package where

import (
	models "beholder-api/internal/application/models"
	lib "beholder-api/internal/gen/som/internal/lib"
)

var Response = newResponse[models.Response](lib.NewKey[models.Response]())

func newResponse[M any](key lib.Key[M]) response[M] {
	return response[M]{
		Body:      lib.NewStringPtr[M](lib.Field(key, "body")),
		CallUID:   lib.NewInt[M, int](lib.Field(key, "call_uid")),
		CreatedAt: lib.NewTimePtr[M](lib.Field(key, "created_at")),
		Headers:   lib.NewStringPtr[M](lib.Field(key, "headers")),
		ID:        lib.NewID[M](lib.Field(key, "id"), "response"),
		Key:       key,
		SentAt:    lib.NewTime[M](lib.Field(key, "sent_at")),
		Status:    lib.NewInt[M, int](lib.Field(key, "status")),
		UID:       lib.NewInt[M, int](lib.Field(key, "uid")),
	}
}

type response[M any] struct {
	lib.Key[M]
	ID        *lib.ID[M]
	UID       *lib.Int[M, int]
	Status    *lib.Int[M, int]
	CallUID   *lib.Int[M, int]
	Body      *lib.StringPtr[M]
	Headers   *lib.StringPtr[M]
	CreatedAt *lib.TimePtr[M]
	SentAt    *lib.Time[M]
}

func (n response[M]) Call() call[M] {
	return newCall[M](lib.Field(n.Key, "call"))
}

type responseEdges[M any] struct {
	lib.Filter[M]
	lib.Key[M]
}
