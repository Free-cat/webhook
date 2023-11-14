package model

import "time"

type Webhook struct {
	Uuid      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
}
