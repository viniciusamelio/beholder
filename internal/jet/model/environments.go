//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Environments struct {
	ID          *int32 `sql:"primary_key"`
	Name        string
	Description *string
	Tags        *string
	BaseURL     *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
