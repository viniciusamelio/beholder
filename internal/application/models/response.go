package models

import (
	"time"

	"github.com/go-surreal/som"
)

type Response struct {
	som.Node
	UID     int     `json:"uid" som:"uid"`
	Status  int     `json:"status" som:"status"`
	CallUID int     `json:"call_uid" som:"call_uid"`
	Body    *string `json:"body" som:"body"`
	Headers *string `json:"headers" som:"headers"`

	CreatedAt *time.Time `json:"created_at" som:"created_at"`
	SentAt    time.Time  `json:"sent_at" som:"sent_at"`

	Call *Call `json:"call" som:"->call"`
}
