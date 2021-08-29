package api

import (
	"context"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
	"github.com/go-resty/resty/v2"
)

func NewInvoices(cfg Config) *Invoices {
	return &Invoices{
		client: NewRestyClient(cfg),
	}
}

type Invoices struct {
	client *resty.Client
}

func (c *Invoices) List(ctx context.Context, req *entity.ListInvoicesReq) (*entity.ListInvoicesResp, error) {
	url := "/invoices"

	var (
		content    entity.ListInvoicesResp
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

func (c *Invoices) Show(ctx context.Context, req *entity.ShowInvoiceReq) (*entity.ShowInvoiceResp, error) {
	url := "/invoices/{identifier}"

	var (
		content    entity.ShowInvoiceResp
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

func (c *Invoices) Create(ctx context.Context, req *entity.CreateInvoiceReq) (*entity.CreateInvoiceResp, error) {
	url := "/invoices"

	var (
		content    entity.CreateInvoiceResp
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

func (c *Invoices) Void(ctx context.Context, req *entity.VoidInvoiceReq) (*entity.VoidInvoiceResp, error) {
	url := "/invoices/{identifier}/void"

	var (
		content    entity.VoidInvoiceResp
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

func (c *Invoices) Resolve(ctx context.Context, req *entity.ResolveInvoiceReq) (*entity.ResolveInvoiceResp, error) {
	url := "/invoices/{identifier}/resolve"

	var (
		content    entity.ResolveInvoiceResp
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
