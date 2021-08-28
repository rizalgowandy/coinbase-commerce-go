package stub

import (
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
	"github.com/segmentio/ksuid"
)

func CreateCheckoutResource() entity.CheckoutResource {
	uuid := ksuid.New().String()

	return entity.CheckoutResource{
		ID:          "id-" + uuid,
		Resource:    enum.ResourceCheckout,
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		LogoURL:     "https://res.cloudinary.com/commerce/image/upload/v1629866084/p6fysdmmql1wxgbmnjfw.png",
		RequestedInfo: []enum.CheckoutRequestedInfo{
			enum.CheckoutRequestedInfoName,
			enum.CheckoutRequestedInfoEmail,
		},
		PricingType: enum.PricingTypeFixedPrice,
		LocalPrice: struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		}{
			Amount:   "100.00",
			Currency: "USD",
		},
		BrandColor:   "#456D9C",
		BrandLogoURL: "https://res.cloudinary.com/commerce/image/upload/v1629866084/p6fysdmmql1wxgbmnjfw.png",
	}
}
