package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/benalucorp/coinbase-commerce-go"
	"github.com/benalucorp/coinbase-commerce-go/pkg/api"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
	"github.com/benalucorp/coinbase-commerce-go/pkg/stub"
	"github.com/kokizzu/gotro/L"
)

func main() {
	client, err := coinbase.NewClient(api.Config{
		Key:   "API_KEY",
		Debug: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Set stub to send webhook response
	ctx := stub.SetWebhookReq(context.Background(), &entity.WebhookReq{
		URL:             "http://127.0.0.1:9000/test",
		SharedSecretKey: "WEBHOOK_KEY",
		Resource:        stub.CreateWebhooksResource(),
	})
	resp, err := client.CreateCharge(ctx, &entity.CreateChargeReq{
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
		log.Fatal(err)
	}
	L.Describe(resp.Data.ID)

	// Create server.
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var req entity.WebhookResource
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		receivedSignature := r.Header.Get("X-CC-Webhook-Signature")
		sharedSecretKey := "WEBHOOK_KEY"
		err = coinbase.CompareWebhookSignature(r.Context(), &req, receivedSignature, sharedSecretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Success.
		L.Describe(r.Header)
		L.Describe(req)
		w.WriteHeader(http.StatusOK)
	})
	_ = http.ListenAndServe(":9000", nil)
}
