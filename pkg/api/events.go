package api

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
)

func NewEvents(cfg Config) *Events {
	return &Events{
		client: NewRestyClient(cfg),
	}
}

type Events struct {
	client *resty.Client
}

func (c *Events) List(ctx context.Context, req *entity.ListEventsReq) (*entity.ListEventsResp, error) {
	url := "/events"

	var (
		content    entity.ListEventsResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetQueryParams(req.QueryParams()).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr.Error
	}

	return &content, nil
}

func (c *Events) Show(ctx context.Context, req *entity.ShowEventReq) (*entity.ShowEventResp, error) {
	url := "/events/{identifier}"

	var (
		content    entity.ShowEventResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("identifier", req.Identifier()).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr.Error
	}

	return &content, nil
}
