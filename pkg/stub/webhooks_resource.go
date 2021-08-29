package stub

import (
	"time"

	"github.com/rizalgowandy/coinbase-commerce-go/pkg/entity"
	"github.com/segmentio/ksuid"
)

func CreateWebhooksResource() entity.WebhookResource {
	return entity.WebhookResource{
		ID:           int(ksuid.New().Time().UnixNano()),
		ScheduledFor: time.Now().Add(10 * time.Second),
		Event:        CreateEventResource(),
	}
}
