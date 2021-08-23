package api

import (
	"github.com/go-resty/resty/v2"
)

const (
	HostURL = "https://api.commerce.coinbase.com"
)

func DefaultHeaders(key string) map[string]string {
	return map[string]string{
		"X-CC-Api-Key": key,
		"X-CC-Version": "2018-03-22",
	}
}

func NewClient(cfg Config) *resty.Client {
	return resty.New().
		SetHostURL(HostURL).
		SetHeaders(DefaultHeaders(cfg.Key)).
		SetTimeout(cfg.Timeout).
		SetRetryCount(cfg.RetryCount).
		SetRetryMaxWaitTime(cfg.RetryMaxWaitTime).
		SetDebug(cfg.Debug)
}
