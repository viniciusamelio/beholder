// Code generated by github.com/go-surreal/som, DO NOT EDIT.

package constant

import "beholder-api/internal/gen/som/internal/lib"

func String[M any](val string) *lib.String[M] {
	return lib.NewString[M](lib.NewVarKey[M](val))
}
