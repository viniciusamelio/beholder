// Code generated by github.com/go-surreal/som, DO NOT EDIT.
package where

import (
	models "beholder-api/internal/application/models"
	lib "beholder-api/internal/gen/som/internal/lib"
)

var Call = newCall[models.Call](lib.NewKey[models.Call]())

func newCall[M any](key lib.Key[M]) call[M] {
	return call[M]{
		Body:        lib.NewStringPtr[M](lib.Field(key, "body")),
		CalledAt:    lib.NewTime[M](lib.Field(key, "called_at")),
		CreatedAt:   lib.NewTimePtr[M](lib.Field(key, "created_at")),
		EnvUID:      lib.NewIntPtr[M, *int](lib.Field(key, "env_uid")),
		ID:          lib.NewID[M](lib.Field(key, "id"), "call"),
		Key:         key,
		Method:      lib.NewString[M](lib.Field(key, "method")),
		Name:        lib.NewString[M](lib.Field(key, "name")),
		Path:        lib.NewStringPtr[M](lib.Field(key, "path")),
		QueryParams: lib.NewStringPtr[M](lib.Field(key, "query_params")),
		SessionUID:  lib.NewIntPtr[M, *int](lib.Field(key, "session_uid")),
		UID:         lib.NewInt[M, int](lib.Field(key, "uid")),
		UserID:      lib.NewStringPtr[M](lib.Field(key, "user_id")),
	}
}

type call[M any] struct {
	lib.Key[M]
	ID          *lib.ID[M]
	UID         *lib.Int[M, int]
	EnvUID      *lib.IntPtr[M, *int]
	SessionUID  *lib.IntPtr[M, *int]
	Name        *lib.String[M]
	Path        *lib.StringPtr[M]
	Body        *lib.StringPtr[M]
	QueryParams *lib.StringPtr[M]
	UserID      *lib.StringPtr[M]
	Method      *lib.String[M]
	CalledAt    *lib.Time[M]
	CreatedAt   *lib.TimePtr[M]
}

func (n call[M]) Session() session[M] {
	return newSession[M](lib.Field(n.Key, "session"))
}

func (n call[M]) Env() environment[M] {
	return newEnvironment[M](lib.Field(n.Key, "env"))
}

type callEdges[M any] struct {
	lib.Filter[M]
	lib.Key[M]
}
