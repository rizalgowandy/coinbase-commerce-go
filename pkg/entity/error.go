package entity

import (
	"fmt"
)

// Reference: https://commerce.coinbase.com/docs/api/#errors

type ErrResp struct {
	Error    ErrDetailResp `json:"error"`
	Warnings []string      `json:"warnings"`
}

func (e ErrResp) Valid() bool {
	return e.Error.Type != "" || e.Error.Message != ""
}

type ErrDetailResp struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e ErrDetailResp) Code() string {
	return e.Type
}

func (e ErrDetailResp) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("coinbase: %s", e.Message)
	}
	return fmt.Sprintf("coinbase: %s", e.Code())
}
