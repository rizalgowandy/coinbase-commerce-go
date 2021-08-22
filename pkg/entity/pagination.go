package entity

import (
	"strconv"
)

// Reference: https://commerce.coinbase.com/docs/api/#pagination

// PaginationReq is a parameter to return list of data with pagination.
// PaginationReq is optional, most fields automatically filled by system.
// If you already have a response with pagination,
// you can generate pagination request directly to traverse next or prev page.
type PaginationReq struct {
	// Order of the resources in the response. desc (default), asc.
	// Order is optional.
	Order string `json:"order"`
	// Limit number of results per call. Accepted values: 0 - 100. Default 25
	// Limit is optional.
	Limit int `json:"limit"`
	// StartingAfter is a cursor for use in pagination.
	// StartingAfter is a resource ID that defines your place in the list.
	// StartingAfter is optional.
	StartingAfter *string `json:"starting_after"`
	// EndingBefore is cursor for use in pagination.
	// EndingBefore is a resource ID that defines your place in the list.
	// EndingBefore is optional.
	EndingBefore *string `json:"ending_before"`
}

func (p PaginationReq) QueryParams() map[string]string {
	res := map[string]string{}
	if p.Order != "" {
		res["order"] = p.Order
	}
	if p.Limit > 0 {
		res["limit"] = strconv.Itoa(p.Limit)
	}
	if p.StartingAfter != nil {
		res["starting_after"] = *p.StartingAfter
	}
	if p.EndingBefore != nil {
		res["ending_before"] = *p.EndingBefore
	}
	return res
}

type PaginationResp struct {
	Order         string  `json:"order"`
	StartingAfter *string `json:"starting_after"`
	EndingBefore  *string `json:"ending_before"`
	Total         int     `json:"total"`
	Yielded       int     `json:"yielded"`
	Limit         int     `json:"limit"`
	PreviousURI   *string `json:"previous_uri"`
	NextURI       *string `json:"next_uri"`
	// CursorRange returns cursors for starting after and ending before.
	// Format: [starting_after, ending_before].
	CursorRange []string `json:"cursor_range"`
}

// HasPrevPage returns true if prev page exists and can be traversed.
func (p *PaginationResp) HasPrevPage() bool {
	return p.PreviousURI != nil
}

// HasNextPage returns true if next page exists and can be traversed.
func (p *PaginationResp) HasNextPage() bool {
	return p.NextURI != nil
}

// PrevPageCursor returns cursor to be used as ending before value.
func (p *PaginationResp) PrevPageCursor() *string {
	if len(p.CursorRange) < 1 {
		return nil
	}
	return &p.CursorRange[0]
}

// NextPageCursor returns cursor to be used as starting after value.
func (p *PaginationResp) NextPageCursor() *string {
	if len(p.CursorRange) < 2 {
		return nil
	}
	return &p.CursorRange[1]
}

// PrevPaginationReq returns pagination request for the prev page result.
func (p *PaginationResp) PrevPaginationReq() PaginationReq {
	return PaginationReq{
		Order:         p.Order,
		Limit:         p.Limit,
		StartingAfter: nil,
		EndingBefore:  p.PrevPageCursor(),
	}
}

// NextPaginationReq returns pagination request for the next page result.
func (p *PaginationResp) NextPaginationReq() PaginationReq {
	return PaginationReq{
		Order:         p.Order,
		Limit:         p.Limit,
		StartingAfter: p.NextPageCursor(),
		EndingBefore:  nil,
	}
}
