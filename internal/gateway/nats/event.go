package nats

import (
	"encoding/json"
	"fmt"

	"github.com/sterligov/banner-rotator/internal/config"
	"github.com/sterligov/banner-rotator/internal/model"
)

type EventGateway struct {
	nats    *Nats
	subject string
}

func NewEventGateway(cfg *config.Config, nats *Nats) *EventGateway {
	return &EventGateway{
		nats:    nats,
		subject: cfg.Queue.Subject,
	}
}

func (eg *EventGateway) Publish(e model.Event) error {
	je, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}

	if err := eg.nats.conn.Publish(eg.subject, je); err != nil {
		return fmt.Errorf("nats publish: %w", err)
	}

	return nil
}
