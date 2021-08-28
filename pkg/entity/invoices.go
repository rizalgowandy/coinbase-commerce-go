package entity

import (
	"errors"
)

// Reference: https://commerce.coinbase.com/docs/api/#list-invoices

type ListInvoicesReq struct {
	PaginationReq
}

type ListInvoicesResp struct {
	Pagination PaginationResp    `json:"pagination"`
	Data       []InvoiceResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#show-an-invoice

// ShowInvoiceReq only requires one of the invoice cod or invoice id is filled.
// If you have already filled one, the other one can be left empty.
type ShowInvoiceReq struct {
	InvoiceCode string `json:"invoice_code"`
	InvoiceID   string `json:"invoice_id"`
}

func (s ShowInvoiceReq) Validate() error {
	if s.InvoiceCode == "" && s.InvoiceID == "" {
		return errors.New("payload: at least one of [code, id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Invoice code has higher priority than id.
func (s ShowInvoiceReq) Identifier() string {
	if s.InvoiceCode != "" {
		return s.InvoiceCode
	}
	return s.InvoiceID
}

type ShowInvoiceResp struct {
	Data InvoiceResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#create-an-invoice

type CreateInvoiceReq struct {
	BusinessName  string             `json:"business_name"`
	CustomerEmail string             `json:"customer_email"`
	CustomerName  string             `json:"customer_name"`
	LocalPrice    CreateInvoicePrice `json:"local_price"`
	Memo          string             `json:"memo"`
}

type CreateInvoicePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type CreateInvoiceResp struct {
	Data InvoiceResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#void-an-invoice

type VoidInvoiceReq struct {
	InvoiceCode string `json:"invoice_code"`
	InvoiceID   string `json:"invoice_id"`
}

func (c VoidInvoiceReq) Validate() error {
	if c.InvoiceCode == "" && c.InvoiceID == "" {
		return errors.New("payload: at least one of [code, id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Invoice code has higher priority than id.
func (c VoidInvoiceReq) Identifier() string {
	if c.InvoiceCode != "" {
		return c.InvoiceCode
	}
	return c.InvoiceID
}

type VoidInvoiceResp struct {
	Data InvoiceResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#resolve-an-invoice

type ResolveInvoiceReq struct {
	InvoiceCode string `json:"invoice_code"`
	InvoiceID   string `json:"invoice_id"`
}

func (r ResolveInvoiceReq) Validate() error {
	if r.InvoiceCode == "" && r.InvoiceID == "" {
		return errors.New("payload: at least one of [code, id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Invoice code has higher priority than id.
func (r ResolveInvoiceReq) Identifier() string {
	if r.InvoiceCode != "" {
		return r.InvoiceCode
	}
	return r.InvoiceID
}

type ResolveInvoiceResp struct {
	Data InvoiceResource `json:"data"`
}
