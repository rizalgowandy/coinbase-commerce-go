package stub

import (
	"time"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
	"github.com/segmentio/ksuid"
)

func CreateInvoiceResource() entity.InvoiceResource {
	uuid := ksuid.New().String()

	return entity.InvoiceResource{
		ID:            "stub_id-" + uuid,
		Resource:      enum.ResourceInvoice,
		Code:          "stub_code-" + uuid,
		Status:        enum.InvoiceStatusOpen,
		BusinessName:  "Crypto Accounting LLC",
		CustomerName:  "Test Customer",
		CustomerEmail: "customer@test.com",
		Memo:          "Taxes and Accounting Services",
		LocalPrice: struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		}{
			Amount:   "100.00",
			Currency: "USD",
		},
		HostedURL: "https://commerce.coinbase.com/invoices/YKDXC4HE",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Charge:    entity.ChargeResource{},
	}
}
