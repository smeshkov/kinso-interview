package event

import "github.com/smeshkov/kinso-interview/app/storage"

type EventDTO struct {
	EventID   string  `json:"eventId"`
	UserID    string  `json:"userId"`
	CreatedAt string  `json:"createdAt"`
	Source    string  `json:"source"` // Gmail, Slack, WhatsApp, LinkedIn and etc
	Priority  float64 `json:"priority"`
	RawData   string  `json:"rawData"` // raw data in JSON format
}

func ToEvent(dto *EventDTO) *storage.Event {
	return &storage.Event{
		ID:        dto.EventID,
		UserID:    dto.UserID,
		CreatedAt: dto.CreatedAt,
		Source:    dto.Source,
		Weight:    dto.Priority,
		RawData:   dto.RawData,
	}
}

func ToDTO(event *storage.Event) *EventDTO {
	return &EventDTO{
		EventID:   event.ID,
		UserID:    event.UserID,
		CreatedAt: event.CreatedAt,
		Source:    event.Source,
		Priority:  event.Weight,
		RawData:   event.RawData,
	}
}

func ToDTOs(events []*storage.Event) []*EventDTO {
	dtos := make([]*EventDTO, len(events))
	for i, event := range events {
		dtos[i] = ToDTO(event)
	}
	return dtos
}
