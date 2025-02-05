package models

import (
	"time"

	"github.com/go-surreal/som"
)

type Call struct {
	som.Node
	UID         int        `json:"uid" som:"uid"`
	EnvUID      *int       `json:"env_uid" som:"env_uid"`
	SessionUID  *int       `json:"session_uid" som:"session_uid"`
	Name        string     `json:"name" som:"name"`
	Path        *string    `json:"path" som:"path"`
	Body        *string    `json:"body" som:"body"`
	QueryParams *string    `json:"query_params" som:"query_params"`
	UserID      *string    `json:"user_id" som:"user_id"`
	Method      string     `json:"method" som:"method"`
	CalledAt    time.Time  `json:"called_at" som:"called_at"`
	CreatedAt   *time.Time `json:"created_at" som:"created_at"`

	Session *Session     `json:"session" som:"->session"`
	Env     *Environment `json:"env" som:"->environment"`
}
