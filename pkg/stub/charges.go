package stub

import (
	"context"
	"time"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/api"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
)

func NewCharges(cfg api.Config) *Charges {
	return &Charges{
		webhooks: NewWebhooks(cfg),
	}
}

type Charges struct {
	webhooks *Webhooks
}

func (c *Charges) Create(ctx context.Context, req *entity.CreateChargeReq) (*entity.CreateChargeResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateChargeResource()
	data.Name = req.Name
	data.Pricing.Local = req.LocalPrice
	data.Description = req.Description
	data.PricingType = req.PricingType

	// Create webhook if requested.
	webhookReq := GetWebhookReq(ctx)
	if webhookReq != nil {
		webhookReq.Resource.Event.Data = data
		if err := c.webhooks.Register(ctx, webhookReq); err != nil {
			return nil, err
		}
	}

	return &entity.CreateChargeResp{
		Data: data,
	}, nil
}

func (c *Charges) Show(ctx context.Context, req *entity.ShowChargeReq) (*entity.ShowChargeResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateChargeResource()
	if req.ChargeID != "" {
		data.ID = req.ChargeID
	}
	if req.ChargeCode != "" {
		data.Code = req.ChargeCode
	}

	return &entity.ShowChargeResp{
		Data: data,
	}, nil
}

func (c *Charges) List(ctx context.Context, req *entity.ListChargesReq) (*entity.ListChargesResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	pagination := CreatePagination()
	if req.Order != "" {
		pagination.Order = req.Order
	}
	if req.Limit > 0 {
		pagination.Limit = req.Limit
	}

	data := make([]entity.ChargeResource, pagination.Limit)
	for i := 0; i < pagination.Limit; i++ {
		data[i] = CreateChargeResource()
	}

	pagination.CursorRange = []string{data[0].ID, data[len(data)-1].ID}

	return &entity.ListChargesResp{
		Pagination: pagination,
		Data:       data,
	}, nil
}

func (c *Charges) Cancel(ctx context.Context, req *entity.CancelChargeReq) (*entity.CancelChargeResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateChargeResource()
	if req.ChargeID != "" {
		data.ID = req.ChargeID
	}
	if req.ChargeCode != "" {
		data.Code = req.ChargeCode
	}
	data.Timeline = append(data.Timeline, struct {
		Time    time.Time                    `json:"time"`
		Status  enum.ChargeStatus            `json:"status"`
		Context enum.ChargeUnresolvedContext `json:"context,omitempty"`
	}{
		Time:    time.Now(),
		Status:  enum.ChargeStatusCanceled,
		Context: "",
	})

	return &entity.CancelChargeResp{
		Data: data,
	}, nil
}

func (c *Charges) Resolve(ctx context.Context, req *entity.ResolveChargeReq) (*entity.ResolveChargeResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateChargeResource()
	if req.ChargeID != "" {
		data.ID = req.ChargeID
	}
	if req.ChargeCode != "" {
		data.Code = req.ChargeCode
	}
	data.ConfirmedAt = time.Now().Add(time.Minute)
	data.Timeline = append(data.Timeline, []struct {
		Time    time.Time                    `json:"time"`
		Status  enum.ChargeStatus            `json:"status"`
		Context enum.ChargeUnresolvedContext `json:"context,omitempty"`
	}{
		{
			Time:    time.Now(),
			Status:  enum.ChargeStatusUnresolved,
			Context: enum.ChargeUnresolvedContextDelayed,
		},
		{
			Time:    time.Now().Add(time.Minute),
			Status:  enum.ChargeStatusResolved,
			Context: "",
		},
	}...)

	return &entity.ResolveChargeResp{
		Data: data,
	}, nil
}
