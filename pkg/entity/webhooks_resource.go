package entity

import (
	"time"
)

type WebhookResource struct {
	ID           int           `json:"id"`
	ScheduledFor time.Time     `json:"scheduled_for"`
	Event        EventResource `json:"event"`
}
