package entity

import (
	"errors"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
)

// Reference: https://commerce.coinbase.com/docs/api/#list-checkouts

type ListCheckoutsReq struct {
	PaginationReq
}

type ListCheckoutsResp struct {
	Pagination PaginationResp     `json:"pagination"`
	Data       []CheckoutResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#show-a-checkout

type ShowCheckoutReq struct {
	CheckoutID string `json:"checkout_id"`
}

func (s ShowCheckoutReq) Validate() error {
	if s.CheckoutID == "" {
		return errors.New("payload: at least one of [id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
// Checkout code has higher priority than id.
func (s ShowCheckoutReq) Identifier() string {
	return s.CheckoutID
}

type ShowCheckoutResp struct {
	Data CheckoutResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#create-a-checkout

type CreateCheckoutReq struct {
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	LocalPrice    CreateCheckoutPrice          `json:"local_price"`
	PricingType   enum.PricingType             `json:"pricing_type"`
	RequestedInfo []enum.CheckoutRequestedInfo `json:"requested_info"`
}

type CreateCheckoutPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type CreateCheckoutResp struct {
	Data CheckoutResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#update-a-checkout

type UpdateCheckoutReq struct {
	CheckoutID    string                       `json:"checkout_id"`
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	LocalPrice    UpdateCheckoutPrice          `json:"local_price"`
	PricingType   enum.PricingType             `json:"pricing_type"`
	RequestedInfo []enum.CheckoutRequestedInfo `json:"requested_info"`
}

func (u *UpdateCheckoutReq) Validate() error {
	if u.CheckoutID == "" {
		return errors.New("payload: at least one of [id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
func (u *UpdateCheckoutReq) Identifier() string {
	return u.CheckoutID
}

type UpdateCheckoutPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type UpdateCheckoutResp struct {
	Data CheckoutResource `json:"data"`
}

// Reference: https://commerce.coinbase.com/docs/api/#delete-a-checkout

type DeleteCheckoutReq struct {
	CheckoutID string `json:"checkout_id"`
}

func (d DeleteCheckoutReq) Validate() error {
	if d.CheckoutID == "" {
		return errors.New("payload: at least one of [id] must be supplied")
	}
	return nil
}

// Identifier returns identifier for current request.
func (d DeleteCheckoutReq) Identifier() string {
	return d.CheckoutID
}

type DeleteCheckoutResp struct{}
