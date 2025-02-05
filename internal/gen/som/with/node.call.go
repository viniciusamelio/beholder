// Code generated by github.com/go-surreal/som, DO NOT EDIT.
package with

import models "beholder-api/internal/application/models"

var Call = call[models.Call]("")

type call[M any] string

func (n call[M]) fetch(M) {}

func (n call[M]) Session() session[M] {
	return session[M](keyed(n, "session"))
}

func (n call[M]) Env() environment[M] {
	return environment[M](keyed(n, "env"))
}
