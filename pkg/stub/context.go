package stub

import (
	"context"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
)

type contextKey string

const (
	CtxKeyStub              = contextKey("Coinbase-Commerce-Stub")
	CtxKeyStubErrDetailResp = contextKey("Coinbase-Commerce-Stub-ErrDetailResp")
	CtxKeyStubWebhookReq    = contextKey("Coinbase-Commerce-Stub-WebhookReq")
)

func Enable(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, CtxKeyStub, true)
}

func Ok(ctx context.Context) bool {
	if ctx == nil {
		return false
	}
	res, ok := ctx.Value(CtxKeyStub).(bool)
	if !ok {
		return false
	}
	return res
}

func SetErrDetailResp(ctx context.Context, err entity.ErrDetailResp) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = Enable(ctx)
	return context.WithValue(ctx, CtxKeyStubErrDetailResp, err)
}

func GetErrDetailResp(ctx context.Context) entity.ErrDetailResp {
	if ctx == nil {
		return entity.ErrDetailResp{}
	}
	res, ok := ctx.Value(CtxKeyStubErrDetailResp).(entity.ErrDetailResp)
	if !ok {
		return entity.ErrDetailResp{}
	}
	return res
}

func SetWebhookReq(ctx context.Context, req *entity.WebhookReq) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = Enable(ctx)
	return context.WithValue(ctx, CtxKeyStubWebhookReq, req)
}

func GetWebhookReq(ctx context.Context) *entity.WebhookReq {
	if ctx == nil {
		return nil
	}
	res, ok := ctx.Value(CtxKeyStubWebhookReq).(*entity.WebhookReq)
	if !ok {
		return nil
	}
	return res
}
