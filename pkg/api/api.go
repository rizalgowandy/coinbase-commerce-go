package api

import (
	"context"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
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

type ChargesItf interface {
	Create(ctx context.Context, req *entity.CreateChargeReq) (*entity.CreateChargeResp, error)
	Show(ctx context.Context, req *entity.ShowChargeReq) (*entity.ShowChargeResp, error)
	List(ctx context.Context, req *entity.ListChargesReq) (*entity.ListChargesResp, error)
	Cancel(ctx context.Context, req *entity.CancelChargeReq) (*entity.CancelChargeResp, error)
	Resolve(ctx context.Context, req *entity.ResolveChargeReq) (*entity.ResolveChargeResp, error)
}

type CheckoutsItf interface {
	List(ctx context.Context, req *entity.ListCheckoutsReq) (*entity.ListCheckoutsResp, error)
	Show(ctx context.Context, req *entity.ShowCheckoutReq) (*entity.ShowCheckoutResp, error)
	Create(ctx context.Context, req *entity.CreateCheckoutReq) (*entity.CreateCheckoutResp, error)
	Update(ctx context.Context, req *entity.UpdateCheckoutReq) (*entity.UpdateCheckoutResp, error)
	Delete(ctx context.Context, req *entity.DeleteCheckoutReq) (*entity.DeleteCheckoutResp, error)
}
