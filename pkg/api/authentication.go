package api

// Reference: https://commerce.coinbase.com/docs/api/#authentication

func DefaultHeaders(key string) map[string]string {
	return map[string]string{
		"X-CC-Api-Key": key,
		"X-CC-Version": Version,
	}
}
