package listener

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"github.com/smeshkov/kinso-interview/app/event"
)

const (
	queueCheckInterval = 10 * time.Second // check queue every 10 seconds
)

// Run - checks the queue every "queueCheckInterval".
func Run(c context.Context, queueAddr string, consumer func(events []*event.EventDTO) error) {
	ticker := time.NewTicker(queueCheckInterval)
	for {
		select {
		case <-c.Done():
			slog.Info("context has been closed, stopping listener")
			return
		case <-ticker.C:
			checkQueue(queueAddr, consumer)
		}
	}
}

func checkQueue(queueAddr string, consumer func(events []*event.EventDTO) error) {
	log := slog.With("queue_address", queueAddr)
	log.Debug("checking queue", "queue_address", queueAddr)

	b, err := os.ReadFile(queueAddr)
	if err != nil {
		log.Error("failed to read file", "error", err)
		return
	}

	events := []*event.EventDTO{}

	err = json.Unmarshal(b, &events)
	if err != nil {
		log.Error("failed to unmarshal json", "error", err)
		return
	}

	if len(events) == 0 {
		log.Debug("no events in queue")
		return
	}

	err = consumer(events)
	if err != nil {
		log.Error("failed to unmarshal json", "error", err)
		return
	}

	log.Debug("processed events", "count", len(events))
}
