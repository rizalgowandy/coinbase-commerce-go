package entity

// Reference: https://commerce.coinbase.com/docs/api/#pagination

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

type PaginationResp struct {
	Order         string   `json:"order"`
	StartingAfter *string  `json:"starting_after"`
	EndingBefore  *string  `json:"ending_before"`
	Total         int      `json:"total"`
	Yielded       int      `json:"yielded"`
	Limit         int      `json:"limit"`
	PreviousURI   *string  `json:"previous_uri"`
	NextURI       *string  `json:"next_uri"`
	CursorRange   []string `json:"cursor_range"`
}
