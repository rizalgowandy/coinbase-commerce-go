package entity

// Reference: https://commerce.coinbase.com/docs/api/#securing-webhooks

type WebhookReq struct {
	URL             string
	SharedSecretKey string
	Resource        WebhookResource
}
