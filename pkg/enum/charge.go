package enum

type ChargeStatus string

const (
	ChargeStatusNew           ChargeStatus = "NEW"
	ChargeStatusPending       ChargeStatus = "PENDING"
	ChargeStatusCompleted     ChargeStatus = "COMPLETED"
	ChargeStatusExpired       ChargeStatus = "EXPIRED"
	ChargeStatusUnresolved    ChargeStatus = "UNRESOLVED"
	ChargeStatusResolved      ChargeStatus = "RESOLVED"
	ChargeStatusCanceled      ChargeStatus = "CANCELED"
	ChargeStatusRefundPending ChargeStatus = "REFUND PENDING"
	ChargeStatusRefunded      ChargeStatus = "REFUNDED"
)

type ChargeUnresolvedContext string

const (
	ChargeUnresolvedContextUnderpaid ChargeUnresolvedContext = "UNDERPAID"
	ChargeUnresolvedContextOverpaid  ChargeUnresolvedContext = "OVERPAID"
	ChargeUnresolvedContextDelayed   ChargeUnresolvedContext = "DELAYED"
	ChargeUnresolvedContextMultiple  ChargeUnresolvedContext = "MULTIPLE"
	ChargeUnresolvedContextManual    ChargeUnresolvedContext = "MANUAL"
	ChargeUnresolvedContextOther     ChargeUnresolvedContext = "OTHER"
)
