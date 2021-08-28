package stub

import (
	"context"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
)

func NewCheckouts() *Checkouts {
	return &Checkouts{}
}

type Checkouts struct{}

func (c *Checkouts) Show(ctx context.Context, req *entity.ShowCheckoutReq) (*entity.ShowCheckoutResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateCheckoutResource()
	data.ID = req.CheckoutID

	return &entity.ShowCheckoutResp{
		Data: data,
	}, nil
}

func (c *Checkouts) Create(ctx context.Context, req *entity.CreateCheckoutReq) (*entity.CreateCheckoutResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateCheckoutResource()
	data.Name = req.Name
	data.Description = req.Description
	data.LocalPrice = req.LocalPrice
	data.PricingType = req.PricingType
	data.RequestedInfo = req.RequestedInfo

	return &entity.CreateCheckoutResp{
		Data: data,
	}, nil
}

func (c *Checkouts) List(ctx context.Context, req *entity.ListCheckoutsReq) (*entity.ListCheckoutsResp, error) {
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

	data := make([]entity.CheckoutResource, pagination.Limit)
	for i := 0; i < pagination.Limit; i++ {
		data[i] = CreateCheckoutResource()
	}

	pagination.CursorRange = []string{data[0].ID, data[len(data)-1].ID}

	return &entity.ListCheckoutsResp{
		Pagination: pagination,
		Data:       data,
	}, nil
}

func (c *Checkouts) Update(ctx context.Context, req *entity.UpdateCheckoutReq) (*entity.UpdateCheckoutResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	data := CreateCheckoutResource()
	data.ID = req.CheckoutID
	data.Name = req.Name
	data.Description = req.Description
	data.LocalPrice = req.LocalPrice
	data.PricingType = req.PricingType
	data.RequestedInfo = req.RequestedInfo

	return &entity.UpdateCheckoutResp{
		Data: data,
	}, nil
}

func (c *Checkouts) Delete(ctx context.Context, req *entity.DeleteCheckoutReq) (*entity.DeleteCheckoutResp, error) {
	if err := CreateErrResp(ctx); err.Valid() {
		return nil, err.Error
	}

	return &entity.DeleteCheckoutResp{}, nil
}
