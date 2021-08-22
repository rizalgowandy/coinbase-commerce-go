package enum

type PricingType string

const (
	PricingTypeNoPrice PricingType = "no_price"

	// FixedPrice is the default to create charge.
	PricingTypeFixedPrice PricingType = "fixed_price"
)
