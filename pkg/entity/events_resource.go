package entity

import (
	"time"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
)

type EventResource struct {
	ID         string         `json:"id"`
	Resource   enum.Resource  `json:"resource"`
	Type       enum.EventType `json:"type"`
	APIVersion string         `json:"api_version"`
	CreatedAt  time.Time      `json:"created_at"`
	Data       ChargeResource `json:"data"`
}
