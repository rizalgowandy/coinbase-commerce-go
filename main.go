package coinbase

import (
	"context"

	"github.com/benalucorp/coinbase-commerce-go/pkg/api"
	"github.com/benalucorp/coinbase-commerce-go/pkg/api/stub"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
)

// NewClient creates a client to interact with Coinbase Commerce API.
func NewClient(cfg api.Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		charges:     api.NewCharges(cfg),
		chargesStub: stub.NewCharges(),
	}, nil
}

// Client is the main client to interact with Coinbase Commerce API.
type Client struct {
	charges     api.ChargesItf
	chargesStub api.ChargesItf
}

// CreateCharge charge a customer with certain amount of currency.
// To get paid in cryptocurrency, you need to create a charge object and
// provide the user with a cryptocurrency address to which they must send cryptocurrency.
// Once a charge is created a customer must broadcast a payment
// to the blockchain before the charge expires.
// Reference: https://commerce.coinbase.com/docs/api/#create-a-charge
func (c Client) CreateCharge(ctx context.Context, req *entity.CreateChargeReq) (*entity.CreateChargeResp, error) {
	if stub.Ok(ctx) {
		return c.chargesStub.Create(ctx, req)
	}
	return c.charges.Create(ctx, req)
}

// ShowCharge retrieves the details of a charge that has been previously created.
// Supply the unique charge code or id that was returned when the charge was created.
// This information is also returned when a charge is first created.
// Reference: https://commerce.coinbase.com/docs/api/#show-a-charge
func (c Client) ShowCharge(ctx context.Context, req *entity.ShowChargeReq) (*entity.ShowChargeResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if stub.Ok(ctx) {
		return c.chargesStub.Show(ctx, req)
	}
	return c.charges.Show(ctx, req)
}

// ListCharges lists all the charges.
// Reference: https://commerce.coinbase.com/docs/api/#list-charges
func (c Client) ListCharges(ctx context.Context, req *entity.ListChargesReq) (*entity.ListChargesResp, error) {
	if stub.Ok(ctx) {
		return c.chargesStub.List(ctx, req)
	}
	return c.charges.List(ctx, req)
}

// CancelCharge cancels a charge that has been previously created.
// Supply the unique charge code or id that was returned when the charge was created.
// This information is also returned when a charge is first created.
//
// Note:
// Only new charges can be successfully canceled.
// Once payment is detected, charge can no longer be canceled.
//
// Reference: https://commerce.coinbase.com/docs/api/#cancel-a-charge
func (c Client) CancelCharge(ctx context.Context, req *entity.CancelChargeReq) (*entity.CancelChargeResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if stub.Ok(ctx) {
		return c.chargesStub.Cancel(ctx, req)
	}
	return c.charges.Cancel(ctx, req)
}

// ResolveCharge resolves a charge that has been previously marked as unresolved.
// Supply the unique charge code or id that was returned when the charge was created.
// This information is also returned when a charge is first created.
//
// Note:
// Only unresolved charges can be successfully resolved
//
// Reference: https://commerce.coinbase.com/docs/api/#resolve-a-charge
func (c Client) ResolveCharge(ctx context.Context, req *entity.ResolveChargeReq) (*entity.ResolveChargeResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if stub.Ok(ctx) {
		return c.chargesStub.Resolve(ctx, req)
	}
	return c.charges.Resolve(ctx, req)
}
