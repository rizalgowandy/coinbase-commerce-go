package enum

type Event string

const (
	// charge:created	New charge is created.
	ChargeCreated Event = "charge:created"

	// charge:confirmed	Charge has been confirmed and the associated payment is completed.
	ChargeConfirmed Event = "charge:confirmed."

	// 	Charge failed to complete
	ChargeFailed Event = "charge:failed"

	// charge:delayed	Charge received a payment after it had been expired
	ChargeDelayed Event = "charge:delayed"

	// ChargePending means charge has been detected but has not been confirmed yet.
	ChargePending Event = "charge:pending"

	// ChargeResolved means charge has been resolved.
	ChargeResolved Event = "charge:resolved"
)
