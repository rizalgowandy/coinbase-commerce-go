package api

import (
	"github.com/go-resty/resty/v2"
)

func NewRestyClient(cfg Config) *resty.Client {
	return resty.New().
		SetHostURL(cfg.HostURL).
		SetHeaders(DefaultHeaders(cfg.Key)).
		SetTimeout(cfg.Timeout).
		SetRetryCount(cfg.RetryCount).
		SetRetryMaxWaitTime(cfg.RetryMaxWaitTime).
		SetDebug(cfg.Debug)
}
