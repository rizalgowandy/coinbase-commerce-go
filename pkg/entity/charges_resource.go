package entity

import (
	"time"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/enum"
)

// Reference: https://commerce.coinbase.com/docs/api/#charge-resource

type ChargeResource struct {
	ID          string        `json:"id"`
	Resource    enum.Resource `json:"resource"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	LogoURL     string        `json:"logo_url"`
	HostedURL   string        `json:"hosted_url"`
	CreatedAt   time.Time     `json:"created_at"`
	ExpiresAt   time.Time     `json:"expires_at"`
	ConfirmedAt time.Time     `json:"confirmed_at"`
	Checkout    struct {
		ID string `json:"id"`
	} `json:"checkout"`
	Timeline []struct {
		Time    time.Time                    `json:"time"`
		Status  enum.ChargeStatus            `json:"status"`
		Context enum.ChargeUnresolvedContext `json:"context,omitempty"`
	} `json:"timeline"`
	Metadata    map[string]string `json:"metadata"`
	PricingType enum.PricingType  `json:"pricing_type"`
	Pricing     struct {
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
	} `json:"pricing"`
	PaymentThreshold struct {
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
	} `json:"payment_threshold"`
	AppliedThreshold struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"applied_threshold"`
	AppliedThresholdType string `json:"applied_threshold_type"`
	Payments             []struct {
		Network       string `json:"network"`
		TransactionID string `json:"transaction_id"`
		Status        string `json:"status"`
		Value         struct {
			Local struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"local"`
			Crypto struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"crypto"`
		} `json:"value"`
		Block struct {
			Height                   int    `json:"height"`
			Hash                     string `json:"hash"`
			ConfirmationsAccumulated int    `json:"confirmations_accumulated"`
			ConfirmationsRequired    int    `json:"confirmations_required"`
		} `json:"block"`
	} `json:"payments"`
	Addresses struct {
		Bitcoin     string `json:"bitcoin"`
		BitcoinCash string `json:"bitcoin_cash"`
		Ethereum    string `json:"ethereum"`
		Litecoin    string `json:"litecoin"`
		Dogecoin    string `json:"dogecoin"`
		USDC        string `json:"usdc"`
		Dai         string `json:"dai"`
	} `json:"addresses"`
}
