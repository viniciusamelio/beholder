package models

import (
	"time"

	"github.com/go-surreal/som"
)

type Session struct {
	som.Node
	UID       int        `json:"id" som:"uid"`
	EnvUID    int        `json:"env_uid" som:"env_uid"`
	UserID    string     `json:"user_id" som:"user_id"`
	CreatedAt *time.Time `json:"created_at" som:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" som:"updated_at"`

	Env *Environment `json:"env" som:"->env"`
}
