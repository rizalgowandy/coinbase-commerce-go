package coinbase

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
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
	client, err = NewClient(Config{
		Key:              os.Getenv("KEY"),
		Timeout:          0,
		RetryCount:       0,
		RetryMaxWaitTime: 0,
		Debug:            true,
	})
	if err != nil {
		fmt.Println("FAIL", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestClient_Charge_Integration(t *testing.T) {
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
		want    *entity.CreateChargeResp
		wantErr bool
	}{
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
					PricingType: "fixed_price",
					Metadata: entity.CreateChargeMetadata{
						CustomerID:   "id_1005",
						CustomerName: "Satoshi Nakamoto",
					},
					RedirectURL: "https://charge/completed/page",
					CancelURL:   "https://charge/canceled/page",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client
			got, err := c.Charges.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			L.Describe(got)
			assert.NotNil(t, got)
		})
	}
}
