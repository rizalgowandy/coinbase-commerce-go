package stub

import (
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
)

func CreatePagination() entity.PaginationResp {
	return entity.PaginationResp{
		Order:         "desc",
		StartingAfter: nil,
		EndingBefore:  nil,
		Total:         1,
		Yielded:       1,
		Limit:         1,
		PreviousURI:   nil,
		NextURI:       nil,
		CursorRange: []string{
			"start_cursor",
			"end_cursor",
		},
	}
}
