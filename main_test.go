package coinbase

import (
	"context"
	"testing"

	"github.com/benalucorp/coinbase-commerce-go/pkg/api"
	"github.com/benalucorp/coinbase-commerce-go/pkg/api/stub"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/kokizzu/gotro/L"
	"github.com/stretchr/testify/assert"
)

func TestClient_CancelCharge(t *testing.T) {
	type fields struct {
		charges     api.ChargesItf
		chargesStub api.ChargesItf
	}
	type args struct {
		ctx context.Context
		req *entity.CancelChargeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Stub error",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.SetErrDetailResp(context.Background(), entity.ErrDetailResp{
					Type:    "bad_request",
					Message: "stub: error triggered",
				}),
				req: &entity.CancelChargeReq{
					ChargeCode: "code",
					ChargeID:   "id",
				},
			},
			wantErr: true,
		},
		{
			name: "Stub success",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.Enable(context.Background()),
				req: &entity.CancelChargeReq{
					ChargeCode: "code",
					ChargeID:   "id",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				charges:     tt.fields.charges,
				chargesStub: tt.fields.chargesStub,
			}
			got, err := c.CancelCharge(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_CreateCharge(t *testing.T) {
	type fields struct {
		charges     api.ChargesItf
		chargesStub api.ChargesItf
	}
	type args struct {
		ctx context.Context
		req *entity.CreateChargeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Stub error",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.SetErrDetailResp(context.Background(), entity.ErrDetailResp{
					Type:    "bad_request",
					Message: "stub: error triggered",
				}),
				req: &entity.CreateChargeReq{},
			},
			wantErr: true,
		},
		{
			name: "Stub success",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.Enable(context.Background()),
				req: &entity.CreateChargeReq{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				charges:     tt.fields.charges,
				chargesStub: tt.fields.chargesStub,
			}
			got, err := c.CreateCharge(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_ListCharges(t *testing.T) {
	type fields struct {
		charges     api.ChargesItf
		chargesStub api.ChargesItf
	}
	type args struct {
		ctx context.Context
		req *entity.ListChargesReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Stub error",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.SetErrDetailResp(context.Background(), entity.ErrDetailResp{
					Type:    "bad_request",
					Message: "stub: error triggered",
				}),
				req: &entity.ListChargesReq{},
			},
			wantErr: true,
		},
		{
			name: "Stub success",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.Enable(context.Background()),
				req: &entity.ListChargesReq{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				charges:     tt.fields.charges,
				chargesStub: tt.fields.chargesStub,
			}
			got, err := c.ListCharges(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCharges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_ResolveCharge(t *testing.T) {
	type fields struct {
		charges     api.ChargesItf
		chargesStub api.ChargesItf
	}
	type args struct {
		ctx context.Context
		req *entity.ResolveChargeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Stub error",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.SetErrDetailResp(context.Background(), entity.ErrDetailResp{
					Type:    "bad_request",
					Message: "stub: error triggered",
				}),
				req: &entity.ResolveChargeReq{
					ChargeCode: "code",
					ChargeID:   "id",
				},
			},
			wantErr: true,
		},
		{
			name: "Stub success",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.Enable(context.Background()),
				req: &entity.ResolveChargeReq{
					ChargeCode: "code",
					ChargeID:   "id",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				charges:     tt.fields.charges,
				chargesStub: tt.fields.chargesStub,
			}
			got, err := c.ResolveCharge(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_ShowCharge(t *testing.T) {
	type fields struct {
		charges     api.ChargesItf
		chargesStub api.ChargesItf
	}
	type args struct {
		ctx context.Context
		req *entity.ShowChargeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Stub error",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.SetErrDetailResp(context.Background(), entity.ErrDetailResp{
					Type:    "bad_request",
					Message: "stub: error triggered",
				}),
				req: &entity.ShowChargeReq{
					ChargeCode: "code",
					ChargeID:   "id",
				},
			},
			wantErr: true,
		},
		{
			name: "Stub success",
			fields: fields{
				charges:     nil,
				chargesStub: stub.NewCharges(),
			},
			args: args{
				ctx: stub.Enable(context.Background()),
				req: &entity.ShowChargeReq{
					ChargeCode: "code",
					ChargeID:   "id",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				charges:     tt.fields.charges,
				chargesStub: tt.fields.chargesStub,
			}
			got, err := c.ShowCharge(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		cfg api.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Missing configuration",
			args: args{
				cfg: api.Config{},
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				cfg: api.Config{
					Key: "sample",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got, err)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
