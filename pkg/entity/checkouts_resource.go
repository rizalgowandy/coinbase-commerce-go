package entity

import (
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
)

// Reference: https://commerce.coinbase.com/docs/api/#checkout-resource

type CheckoutResource struct {
	ID            string                       `json:"id"`
	Resource      enum.Resource                `json:"resource"`
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	LogoURL       string                       `json:"logo_url"`
	RequestedInfo []enum.CheckoutRequestedInfo `json:"requested_info"`
	PricingType   enum.PricingType             `json:"pricing_type"`
	LocalPrice    struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"local_price"`
	BrandColor   string `json:"brand_color"`
	BrandLogoURL string `json:"brand_logo_url"`
}
