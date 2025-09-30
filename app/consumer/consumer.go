package consumer

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/smeshkov/kinso-interview/app/ctx"
	eventdto "github.com/smeshkov/kinso-interview/app/event"
)

// interface to implement by all the other consumers
type Consumer interface {
	IsSupported(event *eventdto.EventDTO) bool
	Consume(c context.Context, event *eventdto.EventDTO) error
	Next() Consumer
}

// chain of responsibility,
// each consumer in the chain handles only certain source
var chain Consumer = newConsumer(&ConsumerImpl{newConsumer(nil)})

// method to call in the listener
func Consume(c context.Context, events []*eventdto.EventDTO) error {
	for _, event := range events {
		var h Consumer = chain

		var isSupported bool

		// chain of responsibility
		for h != nil {
			if h.IsSupported(event) {
				isSupported = true
				err := h.Consume(c, event)
				if err != nil {
					return err
				}
			}
			h = h.Next()
		}

		if !isSupported {
			return fmt.Errorf("no consumer found for event: %v", event)
		}
	}

	return nil
}

// utility for quick instantiation of consumers
type baseConsumer struct {
	next Consumer
}

func newConsumer(next Consumer) *baseConsumer {
	return &baseConsumer{next: next}
}

func (h *baseConsumer) IsSupported(event *eventdto.EventDTO) bool {
	return false
}

func (h *baseConsumer) Consume(c context.Context, event *eventdto.EventDTO) error {
	return nil
}

func (h *baseConsumer) Next() Consumer {
	return h.next
}

type ConsumerImpl struct {
	*baseConsumer
}

func (h *ConsumerImpl) IsSupported(event *eventdto.EventDTO) bool {
	return true
}

func (h *ConsumerImpl) Consume(c context.Context, event *eventdto.EventDTO) error {
	// do something with the event

	p := GetPriority(event)
	e := eventdto.ToEvent(event)
	e.Weight = p

	ok := ctx.DB.Put(e)
	if !ok {
		slog.Debug("event already exists", "event_id", event.EventID)
	}

	return nil
}
