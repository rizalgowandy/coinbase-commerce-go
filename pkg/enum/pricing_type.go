package enum

type PricingType string

const (
	NoPrice PricingType = "no_price"

	// FixedPrice is the default to create charge.
	FixedPrice PricingType = "fixed_price"
)
