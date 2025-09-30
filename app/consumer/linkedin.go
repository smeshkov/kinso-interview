package consumer

import (
	"context"
	"log/slog"

	"github.com/smeshkov/kinso-interview/app/ctx"
	eventdto "github.com/smeshkov/kinso-interview/app/event"
)

type LinkedInConsumer struct {
	*baseConsumer
}

func (h *LinkedInConsumer) IsSupported(event *eventdto.EventDTO) bool {
	return event.Source == "LinkedIn"
}

func (h *LinkedInConsumer) Consume(c context.Context, event *eventdto.EventDTO) error {
	// do something with the event

	ok := ctx.DB.Put(eventdto.ToEvent(event))
	if !ok {
		slog.Debug("event already exists", "event_id", event.EventID)
	}

	return nil
}
