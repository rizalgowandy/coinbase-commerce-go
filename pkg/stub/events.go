package stub

import (
	"context"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
)

func NewEvents() *Events {
	return &Events{}
}

type Events struct{}

func (c *Events) Show(ctx context.Context, req *entity.ShowEventReq) (*entity.ShowEventResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateEventResource()
	data.ID = req.EventID

	return &entity.ShowEventResp{
		Data: data,
	}, nil
}

func (c *Events) List(ctx context.Context, req *entity.ListEventsReq) (*entity.ListEventsResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	pagination := CreatePagination()
	if req.Order != "" {
		pagination.Order = req.Order
	}
	if req.Limit > 0 {
		pagination.Limit = req.Limit
	}

	data := make([]entity.EventResource, pagination.Limit)
	for i := 0; i < pagination.Limit; i++ {
		data[i] = CreateEventResource()
	}

	pagination.CursorRange = []string{data[0].ID, data[len(data)-1].ID}

	return &entity.ListEventsResp{
		Pagination: pagination,
		Data:       data,
	}, nil
}
