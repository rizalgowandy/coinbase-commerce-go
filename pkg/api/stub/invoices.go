package stub

import (
	"context"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
)

func NewInvoices() *Invoices {
	return &Invoices{}
}

type Invoices struct{}

func (c *Invoices) Create(ctx context.Context, req *entity.CreateInvoiceReq) (*entity.CreateInvoiceResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateInvoiceResource()
	data.BusinessName = req.BusinessName
	data.CustomerEmail = req.CustomerEmail
	data.CustomerName = req.CustomerName
	data.LocalPrice = req.LocalPrice
	data.Memo = req.Memo

	return &entity.CreateInvoiceResp{
		Data: data,
	}, nil
}

func (c *Invoices) Show(ctx context.Context, req *entity.ShowInvoiceReq) (*entity.ShowInvoiceResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateInvoiceResource()
	if req.InvoiceID != "" {
		data.ID = req.InvoiceID
	}
	if req.InvoiceCode != "" {
		data.Code = req.InvoiceCode
	}

	return &entity.ShowInvoiceResp{
		Data: data,
	}, nil
}

func (c *Invoices) List(ctx context.Context, req *entity.ListInvoicesReq) (*entity.ListInvoicesResp, error) {
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

	data := make([]entity.InvoiceResource, pagination.Limit)
	for i := 0; i < pagination.Limit; i++ {
		data[i] = CreateInvoiceResource()
	}

	pagination.CursorRange = []string{data[0].ID, data[len(data)-1].ID}

	return &entity.ListInvoicesResp{
		Pagination: pagination,
		Data:       data,
	}, nil
}

func (c *Invoices) Void(ctx context.Context, req *entity.VoidInvoiceReq) (*entity.VoidInvoiceResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateInvoiceResource()
	if req.InvoiceID != "" {
		data.ID = req.InvoiceID
	}
	if req.InvoiceCode != "" {
		data.Code = req.InvoiceCode
	}
	data.Status = enum.InvoiceStatusVoid

	return &entity.VoidInvoiceResp{
		Data: data,
	}, nil
}

func (c *Invoices) Resolve(ctx context.Context, req *entity.ResolveInvoiceReq) (*entity.ResolveInvoiceResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateInvoiceResource()
	if req.InvoiceID != "" {
		data.ID = req.InvoiceID
	}
	if req.InvoiceCode != "" {
		data.Code = req.InvoiceCode
	}
	data.Status = enum.InvoiceStatusPaid

	return &entity.ResolveInvoiceResp{
		Data: data,
	}, nil
}
