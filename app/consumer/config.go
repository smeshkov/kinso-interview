package consumer

import (
	"time"

	eventdto "github.com/smeshkov/kinso-interview/app/event"
)

var (
	sourcePriority = map[string]float64{
		"Slack":    90,
		"WhatsApp": 90,
		"Gmail":    60,
		"LinkedIn": 30,
	}
	typePriority = map[string]float64{
		"alert":        30,
		"message":      10,
		"notification": 10,
	}
)

// Received in the last 15 min: 20
// Received > 24 hours ago: -10
func recencyPriority(ts time.Time) float64 {
	now := time.Now()
	diff := now.Sub(ts)
	if diff.Minutes() < 15 {
		return 20
	}
	if diff.Hours() > 24 {
		return -10
	}
	return 10
}

func GetPriority(event *eventdto.EventDTO) float64 {
	p := sourcePriority[event.Source]
	p += typePriority[event.Source]

	t, err := time.Parse(time.RFC3339, event.CreatedAt)
	if err != nil {
		t = time.Time{}
	}

	p += recencyPriority(t)
	return p
}
