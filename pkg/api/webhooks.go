package api

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
)

func CreateWebhookSignature(_ context.Context, req *entity.WebhookResource, sharedSecretKey string) (string, error) {
	// Create raw request.
	rawReq, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	// Sign raw request.
	h := hmac.New(sha256.New, []byte(sharedSecretKey))
	if _, err = h.Write(rawReq); err != nil {
		return "", err
	}
	signedReq := h.Sum(nil)

	// Encode signed request.
	encodedReq := make([]byte, hex.EncodedLen(len(signedReq)))
	hex.Encode(encodedReq, signedReq)

	// Return encoded signed request as string.
	return string(encodedReq), nil
}

func CompareWebhookSignature(_ context.Context, req *entity.WebhookResource, receivedSignature, sharedSecretKey string) error {
	// Decode received signature.
	decodedSignature, err := hex.DecodeString(receivedSignature)
	if err != nil {
		return err
	}

	// Create raw request.
	rawReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Sign raw request.
	h := hmac.New(sha256.New, []byte(sharedSecretKey))
	if _, err = h.Write(rawReq); err != nil {
		return err
	}
	signedReq := h.Sum(nil)

	// Compare created signature with the received signature.
	if ok := hmac.Equal(signedReq, decodedSignature); !ok {
		return errors.New("coinbase: no signatures found matching the expected signature for payload")
	}
	return nil
}
