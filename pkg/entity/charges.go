package entity

import (
	"errors"

	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
)

// Reference: https://commerce.coinbase.com/docs/api/#create-a-charge

type CreateChargeReq struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	LocalPrice  CreateChargePrice    `json:"local_price"`
	PricingType enum.PricingType     `json:"pricing_type"`
	Metadata    CreateChargeMetadata `json:"metadata"`
	RedirectURL string               `json:"redirect_url"`
	CancelURL   string               `json:"cancel_url"`
}

type CreateChargeMetadata struct {
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
}

type CreateChargePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type CreateChargeResp struct {
	Data ChargeResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#show-a-charge

// ShowChargeReq only requires one of the charge cod or charge id is filled.
// If you have already filled one, the other one can be left empty.
type ShowChargeReq struct {
	ChargeCode string `json:"charge_code"`
	ChargeID   string `json:"charge_id"`
}

func (s ShowChargeReq) Validate() error {
	if s.ChargeCode == "" && s.ChargeID == "" {
		return errors.New("payload: at least one of [code, id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Charge code has higher priority than id.
func (s ShowChargeReq) Identifier() string {
	if s.ChargeCode != "" {
		return s.ChargeCode
	}
	return s.ChargeID
}

type ShowChargeResp struct {
	Data ChargeResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#list-charges

type ListChargesReq struct {
	PaginationReq
}

type ListChargesResp struct {
	Pagination PaginationResp   `json:"pagination"`
	Data       []ChargeResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#cancel-a-charge

type CancelChargeReq struct {
	ChargeCode string `json:"charge_code"`
	ChargeID   string `json:"charge_id"`
}

func (c CancelChargeReq) Validate() error {
	if c.ChargeCode == "" && c.ChargeID == "" {
		return errors.New("payload: at least one of [code, id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Charge code has higher priority than id.
func (c CancelChargeReq) Identifier() string {
	if c.ChargeCode != "" {
		return c.ChargeCode
	}
	return c.ChargeID
}

type CancelChargeResp struct {
	Data ChargeResource `json:"data"`
}
