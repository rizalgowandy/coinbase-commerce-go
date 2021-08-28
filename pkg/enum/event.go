package enum

type EventType string

const (
	// EventTypeChargeCreated means a new charge is created.
	EventTypeChargeCreated EventType = "charge:created"

	// EventTypeChargeConfirmed	means charge has been confirmed and the associated payment is completed.
	EventTypeChargeConfirmed EventType = "charge:confirmed"

	// EventTypeChargeFailed means failed to complete.
	EventTypeChargeFailed EventType = "charge:failed"

	// EventTypeChargeDelayed means received a payment after it had been expired.
	EventTypeChargeDelayed EventType = "charge:delayed"

	// EventTypeChargePending means charge has been detected but has not been confirmed yet.
	EventTypeChargePending EventType = "charge:pending"

	// EventTypeChargeResolved means charge has been resolved.
	EventTypeChargeResolved EventType = "charge:resolved"
)
