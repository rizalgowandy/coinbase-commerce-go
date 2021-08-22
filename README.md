[![Go Doc](https://pkg.go.dev/badge/github.com/benalucorp/coinbase-commerce-go?status.svg)](https://pkg.go.dev/github.com/benalucorp/coinbase-commerce-go?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/benalucorp/coinbase-commerce-go)](https://goreportcard.com/report/github.com/benalucorp/coinbase-commerce-go)
[![Build Status](https://github.com/benalucorp/coinbase-commerce-go/workflows/Go/badge.svg?branch=main)](https://github.com/benalucorp/coinbase-commerce-go/actions?query=branch%3Amain)
[![Sourcegraph](https://sourcegraph.com/github.com/benalucorp/coinbase-commerce-go/-/badge.svg)](https://sourcegraph.com/github.com/benalucorp/coinbase-commerce-go?badge)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/benalucorp/coinbase-commerce-go)](https://www.tickgit.com/browse?repo=github.com/benalucorp/coinbase-commerce-go)

![gdk](https://socialify.git.ci/benalucorp/coinbase-commerce-go/image?description=1&descriptionEditable=Accept%20cryptocurrency%20using%20Coinbase%20Commerce%20API.&font=Inter&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F1885080%3Fs%3D280%26v%3D4&owner=1&pattern=Floating%20Cogs&theme=Light)

## Getting Started

Most requests to the Commerce API must be authenticated with an API key. You can create an API key in your Settings page after creating a [Coinbase Commerce](https://commerce.coinbase.com/signup) account.

## Installation

```shell
go get -v github.com/benalucorp/coinbase-commerce-go
```

## Quick Start

```go
package main

import (
	"context"
	"log"

	"github.com/benalucorp/coinbase-commerce-go"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
)

func main() {
	client, err := coinbase.NewClient(coinbase.Config{
		Key:              "REPLACE_WITH_YOUR_API_KEY",
		Timeout:          0,    // Default: 1 min.
		RetryCount:       0,    // Default: 0 = disable.
		RetryMaxWaitTime: 0,    // Default: 2 secs.
		Debug:            true, // Turn on on development for easier debugging.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a charge.
	createResp, err := client.CreateCharge(context.Background(), &entity.CreateChargeReq{
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
	log.Printf("%+v", createResp)

	// Show a charge.
	showResp, err := client.ShowCharge(context.Background(), &entity.ShowChargeReq{
		ChargeCode: createResp.Data.Code, // Provide either code or id.
		ChargeID:   "",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", showResp)
}
```

## Supported API

- Charges:
   - Create a charge: https://commerce.coinbase.com/docs/api/#create-a-charge
   - Show a charge: https://commerce.coinbase.com/docs/api/#show-a-charge
