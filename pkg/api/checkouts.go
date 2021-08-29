package api

import (
	"context"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
	"github.com/go-resty/resty/v2"
)

func NewCheckouts(cfg Config) *Checkouts {
	return &Checkouts{
		client: NewRestyClient(cfg),
	}
}

type Checkouts struct {
	client *resty.Client
}

func (c *Checkouts) List(ctx context.Context, req *entity.ListCheckoutsReq) (*entity.ListCheckoutsResp, error) {
	url := "/checkouts"

	var (
		content    entity.ListCheckoutsResp
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

func (c *Checkouts) Show(ctx context.Context, req *entity.ShowCheckoutReq) (*entity.ShowCheckoutResp, error) {
	url := "/checkouts/{identifier}"

	var (
		content    entity.ShowCheckoutResp
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

func (c *Checkouts) Create(ctx context.Context, req *entity.CreateCheckoutReq) (*entity.CreateCheckoutResp, error) {
	url := "/checkouts"

	var (
		content    entity.CreateCheckoutResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Post(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr.Error
	}

	return &content, nil
}

func (c *Checkouts) Update(ctx context.Context, req *entity.UpdateCheckoutReq) (*entity.UpdateCheckoutResp, error) {
	url := "/checkouts/{identifier}"

	var (
		content    entity.UpdateCheckoutResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("identifier", req.Identifier()).
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Put(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr.Error
	}

	return &content, nil
}

func (c *Checkouts) Delete(ctx context.Context, req *entity.DeleteCheckoutReq) (*entity.DeleteCheckoutResp, error) {
	url := "/checkouts/{identifier}"

	var (
		content    entity.DeleteCheckoutResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("identifier", req.Identifier()).
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Delete(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr.Error
	}

	return &content, nil
}
