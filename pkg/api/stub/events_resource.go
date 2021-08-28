package stub

import (
	"time"

	"github.com/benalucorp/coinbase-commerce-go/pkg/api"
	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
	"github.com/benalucorp/coinbase-commerce-go/pkg/enum"
	"github.com/segmentio/ksuid"
)

func CreateEventResource() entity.EventResource {
	uuid := ksuid.New().String()

	return entity.EventResource{
		ID:         "stub_id-" + uuid,
		Resource:   enum.ResourceEvent,
		Type:       enum.EventTypeChargeConfirmed,
		APIVersion: api.Version,
		CreatedAt:  time.Now(),
		Data:       entity.ChargeResource{},
	}
}
