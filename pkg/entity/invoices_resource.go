package entity

import (
	"time"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
)

type InvoiceResource struct {
	ID            string             `json:"id"`
	Resource      enum.Resource      `json:"resource"`
	Code          string             `json:"code"`
	Status        enum.InvoiceStatus `json:"status"`
	BusinessName  string             `json:"business_name"`
	CustomerName  string             `json:"customer_name"`
	CustomerEmail string             `json:"customer_email"`
	Memo          string             `json:"memo"`
	LocalPrice    struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"local_price"`
	HostedURL string         `json:"hosted_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Charge    ChargeResource `json:"charge"`
}
