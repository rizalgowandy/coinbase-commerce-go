package api

import (
	"context"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/go-resty/resty/v2"
)

func NewCharges(client *resty.Client) *Charges {
	return &Charges{
		client: client,
	}
}

type Charges struct {
	client *resty.Client
}

func (a *Charges) Create(ctx context.Context, req *entity.CreateChargeReq) (*entity.CreateChargeResp, error) {
	url := "/charges"

	var (
		content    entity.CreateChargeResp
		contentErr entity.ErrResp
	)

	_, err := a.client.R().
		SetContext(ctx).
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

func (a *Charges) Show(ctx context.Context, req *entity.ShowChargeReq) (*entity.ShowChargeResp, error) {
	url := "/charges/{identifier}"

	var (
		content    entity.ShowChargeResp
		contentErr entity.ErrResp
	)

	_, err := a.client.R().
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

func (a *Charges) List(ctx context.Context, req *entity.ListChargesReq) (*entity.ListChargesResp, error) {
	url := "/charges"

	var (
		content    entity.ListChargesResp
		contentErr entity.ErrResp
	)

	_, err := a.client.R().
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
