package coinbase

import (
	"context"
	"errors"
	"time"

	"github.com/benalucorp/coinbase-commerce-go/pkg/api"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/go-resty/resty/v2"
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

// NewClient creates a client to interact with Coinbase Commerce API.
func NewClient(cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	r := resty.New().
		SetHostURL("https://api.commerce.coinbase.com").
		SetHeaders(map[string]string{
			"X-CC-Api-Key": cfg.Key,
			"X-CC-Version": "2018-03-22",
		}).
		SetTimeout(cfg.Timeout).
		SetRetryCount(cfg.RetryCount).
		SetRetryMaxWaitTime(cfg.RetryMaxWaitTime).
		SetDebug(cfg.Debug)

	return &Client{
		charges: api.NewCharges(r),
	}, nil
}

type Client struct {
	charges *api.Charges
}

// CreateCharge charge a customer with certain amount of currency.
// To get paid in cryptocurrency, you need to create a charge object and
// provide the user with a cryptocurrency address to which they must send cryptocurrency.
// Once a charge is created a customer must broadcast a payment
// to the blockchain before the charge expires.
// Reference: https://commerce.coinbase.com/docs/api/#create-a-charge
func (c Client) CreateCharge(ctx context.Context, req *entity.CreateChargeReq) (*entity.CreateChargeResp, error) {
	if c.charges == nil {
		return nil, errors.New("client: initialize first")
	}

	if req == nil {
		return nil, errors.New("payload: missing")
	}

	return c.charges.Create(ctx, req)
}

// ShowCharge retrieves the details of a charge that has been previously created.
// Supply the unique charge code or id that was returned when the charge was created.
// This information is also returned when a charge is first created.
// Reference: https://commerce.coinbase.com/docs/api/#show-a-charge
func (c Client) ShowCharge(ctx context.Context, req *entity.ShowChargeReq) (*entity.ShowChargeResp, error) {
	if c.charges == nil {
		return nil, errors.New("client: initialize first")
	}

	if req == nil {
		return nil, errors.New("payload: missing")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return c.charges.Show(ctx, req)
}
