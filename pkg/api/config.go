package api

import (
	"errors"
	"time"
)

type Config struct {
	// Key is the authentication API key.
	// Most requests to the Commerce API must be authenticated with an API key.
	// You can create an API key in your Settings page after creating a Coinbase Commerce account.
	// Reference: https://commerce.coinbase.com/docs/api/#authentication
	Key string
	// Timeout describes total waiting time before a request is treated as timeout.
	// Default: 1 min.
	Timeout time.Duration
	// RetryCount describes total number of retry in case error occurred.
	// Set 0 to disable retry mechanism.
	// Default: 3.
	RetryCount int
	// RetryMaxWaitTime describes total waiting time between each retry.
	// Default: 2 second.
	RetryMaxWaitTime time.Duration
	// Debug describes the client to enter debug mode.
	Debug bool
}

func (c *Config) Validate() error {
	if c.Key == "" {
		return errors.New("config: invalid key")
	}
	if c.Timeout <= 0 {
		c.Timeout = time.Minute
	}
	if c.RetryCount < 0 {
		c.RetryCount = 3
	}
	if c.RetryMaxWaitTime <= 0 {
		c.RetryMaxWaitTime = 2 * time.Second
	}
	return nil
}
