package models

import (
	"time"
)

type Response struct {
	ID      int     `json:"id"`
	Status  int     `json:"status"`
	CallID  int     `json:"call_id"`
	Body    *string `json:"body"`
	Headers *string `json:"headers"`

	CreatedAt *time.Time `json:"created_at"`
	SentAt    time.Time  `json:"sent_at"`
}
