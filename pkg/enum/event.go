package enum

type Event string

const (
	// EventChargeCreated means a new charge is created.
	EventChargeCreated Event = "charge:created"

	// EventChargeConfirmed	means charge has been confirmed and the associated payment is completed.
	EventChargeConfirmed Event = "charge:confirmed"

	// EventChargeFailed means failed to complete.
	EventChargeFailed Event = "charge:failed"

	// EventChargeDelayed means received a payment after it had been expired.
	EventChargeDelayed Event = "charge:delayed"

	// EventChargePending means charge has been detected but has not been confirmed yet.
	EventChargePending Event = "charge:pending"

	// EventChargeResolved means charge has been resolved.
	EventChargeResolved Event = "charge:resolved"
)
