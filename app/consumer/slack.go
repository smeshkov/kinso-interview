package consumer

import (
	"context"
	"log/slog"

	"github.com/smeshkov/kinso-interview/app/ctx"
	eventdto "github.com/smeshkov/kinso-interview/app/event"
)

type SlackConsumer struct {
	*baseConsumer
}

func (h *SlackConsumer) IsSupported(event *eventdto.EventDTO) bool {
	return event.Source == "Slack"
}

func (h *SlackConsumer) Consume(c context.Context, event *eventdto.EventDTO) error {
	// do something with the event

	ok := ctx.DB.Put(eventdto.ToEvent(event))
	if !ok {
		slog.Debug("event already exists", "event_id", event.EventID)
	}

	return nil
}
