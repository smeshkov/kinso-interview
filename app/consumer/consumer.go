package consumer

import (
	"context"
	"fmt"

	"github.com/smeshkov/kinso-interview/app/event"
)

// interface to implement by all the other consumers
type Consumer interface {
	IsSupported(event *event.EventDTO) bool
	Consume(c context.Context, event *event.EventDTO) error
	Next() Consumer
}

// chain of responsibility,
// each consumer in the chain handles only certain source
var chain Consumer = newConsumer(&GmailConsumer{
	newConsumer(&SlackConsumer{
		newConsumer(&WhatsAppConsumer{
			newConsumer(&LinkedInConsumer{
				newConsumer(nil),
			}),
		}),
	}),
})

// method to call in the listener
func Consume(c context.Context, events []*event.EventDTO) error {
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

func (h *baseConsumer) IsSupported(event *event.EventDTO) bool {
	return false
}

func (h *baseConsumer) Consume(c context.Context, event *event.EventDTO) error {
	return nil
}

func (h *baseConsumer) Next() Consumer {
	return h.next
}
