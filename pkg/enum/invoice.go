package enum

type InvoiceStatus string

const (
	InvoiceStatusOpen   InvoiceStatus = "OPEN"
	InvoiceStatusViewed InvoiceStatus = "VIEWED"
	InvoiceStatusVoid   InvoiceStatus = "VOID"
	InvoiceStatusPaid   InvoiceStatus = "PAID"
)
