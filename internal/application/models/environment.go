package models

import (
	"time"

	"github.com/go-surreal/som"
)

type Environment struct {
	som.Node
	UID       int        `json:"id" som:"id"`
	Name      string     `json:"name" som:"name"`
	Tags      []string   `json:"tags" som:"tags"`
	BaseUrl   string     `json:"base_url" som:"base_url"`
	CreatedAt *time.Time `json:"created_at" som:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" som:"updated_at"`

	Sessions *[]*Session `json:"sessions" som:"->session.environment"`
}
