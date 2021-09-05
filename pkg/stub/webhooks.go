package stub

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/kokizzu/gotro/L"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/api"
	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
	"github.com/robfig/cron/v3"
)

func NewWebhooks(cfg api.Config) *Webhooks {
	ctrl := cron.New(
		cron.WithParser(cron.NewParser(
			cron.SecondOptional |
				cron.Minute |
				cron.Hour |
				cron.Dom |
				cron.Month |
				cron.Dow |
				cron.Descriptor,
		)),
	)
	ctrl.Start()

	client := resty.New().
		SetTimeout(cfg.Timeout).
		SetRetryCount(cfg.RetryCount).
		SetRetryMaxWaitTime(cfg.RetryMaxWaitTime).
		SetDebug(cfg.Debug)

	return &Webhooks{
		ctrl:    ctrl,
		client:  client,
		entries: sync.Map{},
		debug:   cfg.Debug,
	}
}

type Webhooks struct {
	ctrl    *cron.Cron
	client  *resty.Client
	entries sync.Map
	debug   bool
}

func (w *Webhooks) Register(ctx context.Context, req *entity.WebhookReq) error {
	signature, err := api.CreateWebhookSignature(ctx, &req.Resource, req.SharedSecretKey)
	if err != nil {
		return err
	}

	spec := fmt.Sprintf("%d %d * * * *", req.Resource.ScheduledFor.Second(), req.Resource.ScheduledFor.Minute())

	id := req.Resource.Event.ID
	key := id + "-" + spec

	if w.debug {
		L.Describe("stub: webhook will be sent at " + req.Resource.ScheduledFor.String())
	}

	entryID, err := w.ctrl.AddFunc(spec, func() {
		res, ok := w.entries.LoadAndDelete(key)
		if !ok {
			return
		}
		defer w.ctrl.Remove(res.(cron.EntryID))

		resp, errReq := w.client.R().
			SetContext(ctx).
			SetHeaders(map[string]string{
				"X-Request-ID":           id,
				"X-CC-Webhook-Signature": signature,
				"Content-Type":           "application/json",
			}).
			SetBody(req.Resource).
			Post(req.URL)
		if L.IsError(errReq, "stub: send webhook failure for "+id) {
			return
		}
		if resp.IsError() {
			L.IsError(errors.New(resp.Status()), "stub: webhook rejected for "+id)
		}
	})
	if err != nil {
		return err
	}

	w.entries.Store(key, entryID)
	return nil
}
