package api

const (
	HostURL = "https://api.commerce.coinbase.com"
)

func DefaultHeaders(key string) map[string]string {
	return map[string]string{
		"X-CC-Api-Key": key,
		"X-CC-Version": "2018-03-22",
	}
}
