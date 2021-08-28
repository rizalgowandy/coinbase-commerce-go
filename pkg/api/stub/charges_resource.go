package stub

import (
	"time"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
	"github.com/segmentio/ksuid"
)

func CreateChargeResource() entity.ChargeResource {
	uuid := ksuid.New().String()

	return entity.ChargeResource{
		ID:          "id-" + uuid,
		Resource:    enum.ResourceCharge,
		Code:        "code-" + uuid,
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		LogoURL:     "",
		HostedURL:   "https://commerce.coinbase.com/charges/FLJ8D3PW",
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(24 * time.Hour),
		ConfirmedAt: time.Time{},
		Checkout: struct {
			ID string `json:"id"`
		}{},
		Timeline: []struct {
			Time    time.Time                    `json:"time"`
			Status  enum.ChargeStatus            `json:"status"`
			Context enum.ChargeUnresolvedContext `json:"context,omitempty"`
		}{
			{
				Time:    time.Now(),
				Status:  enum.ChargeStatusNew,
				Context: enum.ChargeUnresolvedContextNone,
			},
		},
		Metadata:    struct{}{},
		PricingType: enum.PricingTypeFixedPrice,
		Pricing: struct {
			Local struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"local"`
			Bitcoin struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"bitcoin"`
			BitcoinCash struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"bitcoin_cash"`
			Ethereum struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"ethereum"`
			Litecoin struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"litecoin"`
			Dogecoin struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"dogecoin"`
			USDC struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"usdc"`
			Dai struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"dai"`
		}{
			Local: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "100.00",
				Currency: "USD",
			},
			Bitcoin: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "0.00198807",
				Currency: "BTC",
			},
			BitcoinCash: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{},
			Ethereum: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "0.030005000",
				Currency: "ETH",
			},
			Litecoin: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "0.52861107",
				Currency: "LTC",
			},
			Dogecoin: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "312.15857660",
				Currency: "DOGE",
			},
			USDC: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "100.000000",
				Currency: "USDC",
			},
			Dai: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "99.917617924021644154",
				Currency: "DAI",
			},
		},
		PaymentThreshold: struct {
			OverpaymentAbsoluteThreshold struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"overpayment_absolute_threshold"`
			OverpaymentRelativeThreshold  string `json:"overpayment_relative_threshold"`
			UnderpaymentAbsoluteThreshold struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"underpayment_absolute_threshold"`
			UnderpaymentRelativeThreshold string `json:"underpayment_relative_threshold"`
		}{
			OverpaymentAbsoluteThreshold: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "5.00",
				Currency: "USD",
			},
			OverpaymentRelativeThreshold: "0.005",
			UnderpaymentAbsoluteThreshold: struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   "5.00",
				Currency: "USD",
			},
			UnderpaymentRelativeThreshold: "0.005",
		},
		AppliedThreshold: struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		}{},
		AppliedThresholdType: "",
		Payments:             nil,
		Addresses: struct {
			Bitcoin     string `json:"bitcoin"`
			BitcoinCash string `json:"bitcoin_cash"`
			Ethereum    string `json:"ethereum"`
			Litecoin    string `json:"litecoin"`
			Dogecoin    string `json:"dogecoin"`
			USDC        string `json:"usdc"`
			Dai         string `json:"dai"`
		}{
			Bitcoin:     "BitcoinAddress",
			BitcoinCash: "BitcoinCashAddress",
			Ethereum:    "EthereumAddress",
			Litecoin:    "LitecoinAddress",
			Dogecoin:    "DogecoinAddress",
			USDC:        "USDCAddress",
			Dai:         "DaiAddress",
		},
	}
}
