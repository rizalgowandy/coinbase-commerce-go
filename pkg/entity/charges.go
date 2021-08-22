package entity

// Reference: https://commerce.coinbase.com/docs/api/#create-a-charge

type CreateChargeReq struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	LocalPrice  CreateChargePrice    `json:"local_price"`
	PricingType string               `json:"pricing_type"`
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
