package api

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
)

func NewCharges(cfg Config) *Charges {
	return &Charges{
		client: NewRestyClient(cfg),
	}
}

type Charges struct {
	client *resty.Client
}

func (c *Charges) Create(ctx context.Context, req *entity.CreateChargeReq) (*entity.CreateChargeResp, error) {
	url := "/charges"

	var (
		content    entity.CreateChargeResp
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

func (c *Charges) Show(ctx context.Context, req *entity.ShowChargeReq) (*entity.ShowChargeResp, error) {
	url := "/charges/{identifier}"

	var (
		content    entity.ShowChargeResp
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

func (c *Charges) List(ctx context.Context, req *entity.ListChargesReq) (*entity.ListChargesResp, error) {
	url := "/charges"

	var (
		content    entity.ListChargesResp
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

func (c *Charges) Cancel(ctx context.Context, req *entity.CancelChargeReq) (*entity.CancelChargeResp, error) {
	url := "/charges/{identifier}/cancel"

	var (
		content    entity.CancelChargeResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("identifier", req.Identifier()).
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

func (c *Charges) Resolve(ctx context.Context, req *entity.ResolveChargeReq) (*entity.ResolveChargeResp, error) {
	url := "/charges/{identifier}/resolve"

	var (
		content    entity.ResolveChargeResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("identifier", req.Identifier()).
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
