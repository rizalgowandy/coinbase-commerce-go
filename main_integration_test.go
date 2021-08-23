package coinbase

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/benalucorp/coinbase-commerce-go/pkg/api"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
	"github.com/kokizzu/gotro/L"
	"github.com/stretchr/testify/assert"
)

// How to run all integration test:
// $ KEY=REAL_API_KEY go test -v . -run . -integration

var (
	integration bool
	client      *Client
)

func TestMain(m *testing.M) {
	flag.BoolVar(&integration, "integration", false, "enable integration test")
	flag.Parse()

	if !integration {
		os.Exit(m.Run())
	}

	var err error
	client, err = NewClient(api.Config{
		Key:   os.Getenv("KEY"),
		Debug: true,
	})
	if L.IsError(err, "client: create failure") {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestClient_CreateCharge_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
		req *entity.CreateChargeReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Empty request",
			args: args{
				ctx: context.Background(),
				req: &entity.CreateChargeReq{
					Name:        "",
					Description: "",
					LocalPrice:  entity.CreateChargePrice{},
					PricingType: "",
					Metadata:    entity.CreateChargeMetadata{},
					RedirectURL: "",
					CancelURL:   "",
				},
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &entity.CreateChargeReq{
					Name:        "The Sovereign Individual",
					Description: "Mastering the Transition to the Information Age",
					LocalPrice: entity.CreateChargePrice{
						Amount:   "100.00",
						Currency: "USD",
					},
					PricingType: enum.PricingTypeFixedPrice,
					Metadata: entity.CreateChargeMetadata{
						CustomerID:   "id_1005",
						CustomerName: "Satoshi Nakamoto",
					},
					RedirectURL: "https://charge/completed/page",
					CancelURL:   "https://charge/canceled/page",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
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

func TestClient_ShowCharge_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
		req *entity.ShowChargeReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Missing identifier",
			args: args{
				ctx: context.Background(),
				req: &entity.ShowChargeReq{
					ChargeCode: "",
					ChargeID:   "",
				},
			},
			wantErr: true,
		},
		{
			name: "Show using charge code",
			args: args{
				ctx: context.Background(),
				req: &entity.ShowChargeReq{
					ChargeCode: func() string {
						c := client
						got, err := c.charges.Create(context.Background(), &entity.CreateChargeReq{
							Name:        "The Sovereign Individual",
							Description: "Mastering the Transition to the Information Age",
							LocalPrice: entity.CreateChargePrice{
								Amount:   "100.00",
								Currency: "USD",
							},
							PricingType: enum.PricingTypeFixedPrice,
							Metadata: entity.CreateChargeMetadata{
								CustomerID:   "id_1005",
								CustomerName: "Satoshi Nakamoto",
							},
							RedirectURL: "https://charge/completed/page",
							CancelURL:   "https://charge/canceled/page",
						})
						if err != nil {
							return ""
						}
						return got.Data.Code
					}(),
					ChargeID: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Show using charge id",
			args: args{
				ctx: context.Background(),
				req: &entity.ShowChargeReq{
					ChargeCode: "",
					ChargeID: func() string {
						c := client
						got, err := c.charges.Create(context.Background(), &entity.CreateChargeReq{
							Name:        "The Sovereign Individual",
							Description: "Mastering the Transition to the Information Age",
							LocalPrice: entity.CreateChargePrice{
								Amount:   "100.00",
								Currency: "USD",
							},
							PricingType: enum.PricingTypeFixedPrice,
							Metadata: entity.CreateChargeMetadata{
								CustomerID:   "id_1005",
								CustomerName: "Satoshi Nakamoto",
							},
							RedirectURL: "https://charge/completed/page",
							CancelURL:   "https://charge/canceled/page",
						})
						if err != nil {
							return ""
						}
						return got.Data.ID
					}(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
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

func TestClient_ListCharges_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
		req *entity.ListChargesReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Next page",
			args: args{
				ctx: context.Background(),
				req: &entity.ListChargesReq{
					PaginationReq: func() entity.PaginationReq {
						c := client
						got, err := c.ListCharges(context.Background(), &entity.ListChargesReq{
							PaginationReq: entity.PaginationReq{
								Limit: 2,
							},
						})
						if err != nil {
							return entity.PaginationReq{}
						}
						return got.Pagination.NextPaginationReq()
					}(),
				},
			},
			wantErr: false,
		},
		{
			name: "Prev page",
			args: args{
				ctx: context.Background(),
				req: &entity.ListChargesReq{
					PaginationReq: func() entity.PaginationReq {
						c := client
						got, err := c.ListCharges(context.Background(), &entity.ListChargesReq{
							PaginationReq: entity.PaginationReq{
								Limit: 2,
							},
						})
						if err != nil {
							return entity.PaginationReq{}
						}
						got, err = c.ListCharges(context.Background(), &entity.ListChargesReq{
							PaginationReq: got.Pagination.NextPaginationReq(),
						})
						if err != nil {
							return entity.PaginationReq{}
						}
						return got.Pagination.PrevPaginationReq()
					}(),
				},
			},
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &entity.ListChargesReq{
					PaginationReq: entity.PaginationReq{
						Limit: 5,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
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

func TestClient_CancelCharge_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
		req *entity.CancelChargeReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Missing identifier",
			args: args{
				ctx: context.Background(),
				req: &entity.CancelChargeReq{
					ChargeCode: "",
					ChargeID:   "",
				},
			},
			wantErr: true,
		},
		{
			name: "Cancel using charge code",
			args: args{
				ctx: context.Background(),
				req: &entity.CancelChargeReq{
					ChargeCode: func() string {
						c := client
						got, err := c.charges.Create(context.Background(), &entity.CreateChargeReq{
							Name:        "The Sovereign Individual",
							Description: "Mastering the Transition to the Information Age",
							LocalPrice: entity.CreateChargePrice{
								Amount:   "100.00",
								Currency: "USD",
							},
							PricingType: enum.PricingTypeFixedPrice,
							Metadata: entity.CreateChargeMetadata{
								CustomerID:   "id_1005",
								CustomerName: "Satoshi Nakamoto",
							},
							RedirectURL: "https://charge/completed/page",
							CancelURL:   "https://charge/canceled/page",
						})
						if err != nil {
							return ""
						}
						return got.Data.Code
					}(),
					ChargeID: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Cancel using charge id",
			args: args{
				ctx: context.Background(),
				req: &entity.CancelChargeReq{
					ChargeCode: "",
					ChargeID: func() string {
						c := client
						got, err := c.charges.Create(context.Background(), &entity.CreateChargeReq{
							Name:        "The Sovereign Individual",
							Description: "Mastering the Transition to the Information Age",
							LocalPrice: entity.CreateChargePrice{
								Amount:   "100.00",
								Currency: "USD",
							},
							PricingType: enum.PricingTypeFixedPrice,
							Metadata: entity.CreateChargeMetadata{
								CustomerID:   "id_1005",
								CustomerName: "Satoshi Nakamoto",
							},
							RedirectURL: "https://charge/completed/page",
							CancelURL:   "https://charge/canceled/page",
						})
						if err != nil {
							return ""
						}
						return got.Data.ID
					}(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
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

func TestClient_ResolveCharge_Integration(t *testing.T) {
	if !integration {
		return
	}

	type args struct {
		ctx context.Context
		req *entity.ResolveChargeReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Missing identifier",
			args: args{
				ctx: context.Background(),
				req: &entity.ResolveChargeReq{
					ChargeCode: "",
					ChargeID:   "",
				},
			},
			wantErr: true,
		},
		{
			name: "Resolve using charge code",
			args: args{
				ctx: context.Background(),
				req: &entity.ResolveChargeReq{
					ChargeCode: func() string {
						c := client
						got, err := c.charges.Create(context.Background(), &entity.CreateChargeReq{
							Name:        "The Sovereign Individual",
							Description: "Mastering the Transition to the Information Age",
							LocalPrice: entity.CreateChargePrice{
								Amount:   "100.00",
								Currency: "USD",
							},
							PricingType: enum.PricingTypeFixedPrice,
							Metadata: entity.CreateChargeMetadata{
								CustomerID:   "id_1005",
								CustomerName: "Satoshi Nakamoto",
							},
							RedirectURL: "https://charge/completed/page",
							CancelURL:   "https://charge/resolveed/page",
						})
						if err != nil {
							return ""
						}
						return got.Data.Code
					}(),
					ChargeID: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Resolve using charge id",
			args: args{
				ctx: context.Background(),
				req: &entity.ResolveChargeReq{
					ChargeCode: "",
					ChargeID: func() string {
						c := client
						got, err := c.charges.Create(context.Background(), &entity.CreateChargeReq{
							Name:        "The Sovereign Individual",
							Description: "Mastering the Transition to the Information Age",
							LocalPrice: entity.CreateChargePrice{
								Amount:   "100.00",
								Currency: "USD",
							},
							PricingType: enum.PricingTypeFixedPrice,
							Metadata: entity.CreateChargeMetadata{
								CustomerID:   "id_1005",
								CustomerName: "Satoshi Nakamoto",
							},
							RedirectURL: "https://charge/completed/page",
							CancelURL:   "https://charge/resolveed/page",
						})
						if err != nil {
							return ""
						}
						return got.Data.ID
					}(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
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
