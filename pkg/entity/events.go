package entity

import (
	"errors"
)

// Reference: https://commerce.coinbase.com/docs/api/#list-events

type ListEventsReq struct {
	PaginationReq
}

type ListEventsResp struct {
	Pagination PaginationResp  `json:"pagination"`
	Data       []EventResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#show-an-event

type ShowEventReq struct {
	EventID string `json:"event_id"`
}

func (s ShowEventReq) Validate() error {
	if s.EventID == "" {
		return errors.New("payload: at least one of [id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Event code has higher priority than id.
func (s ShowEventReq) Identifier() string {
	return s.EventID
}

type ShowEventResp struct {
	Data EventResource `json:"data"`
}
