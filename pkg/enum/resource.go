package enum

type Resource string

const (
	ResourceCharge   Resource = "charge"
	ResourceCheckout Resource = "checkout"
	ResourceInvoice  Resource = "invoice"
	ResourceEvent    Resource = "event"
)
