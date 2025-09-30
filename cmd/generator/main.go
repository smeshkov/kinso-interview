package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid" // Requires: go get github.com/google/uuid
)

// EventDTO represents the desired JSON structure.
type EventDTO struct {
	EventID   string `json:"eventId"`
	UserID    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	Source    string `json:"source"`  // Gmail, Slack, WhatsApp, LinkedIn and etc
	RawData   string `json:"rawData"` // raw data in JSON format
}

// generateEvent creates a single EventDTO object with mock data.
func generateEvent() EventDTO {
	sources := []string{"Gmail", "Slack", "WhatsApp", "LinkedIn"}

	// 1. Generate unique IDs
	eventID := uuid.New().String()
	userID := uuid.New().String()[:8] // Truncate for a simpler look

	// 2. Generate a recent, but random, UTC timestamp in ISO 8601 format (ending in Z)
	now := time.Now().UTC()
	randomOffsetSeconds := time.Duration(rand.Intn(86400*7)) * time.Second // Up to 7 days in the past
	createdAt := now.Add(-randomOffsetSeconds).Format("2006-01-02T15:04:05.000000Z")

	// 3. Select a random source
	source := sources[rand.Intn(len(sources))]

	// 4. Generate RawData as a JSON string
	rawDataContent := map[string]interface{}{
		"type":        []string{"message", "notification", "alert"}[rand.Intn(3)],
		"status":      []string{"read", "unread", "archived"}[rand.Intn(3)],
		"data_length": rand.Intn(451) + 50, // 50 to 500
	}
	rawDataBytes, _ := json.Marshal(rawDataContent)
	rawData := string(rawDataBytes)

	return EventDTO{
		EventID:   eventID,
		UserID:    userID,
		CreatedAt: createdAt,
		Source:    source,
		RawData:   rawData,
	}
}

func main() {
	const count = 100

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a slice to hold the events
	events := make([]EventDTO, count)

	// Populate the slice with 100 generated events
	for i := 0; i < count; i++ {
		events[i] = generateEvent()
	}

	// Marshal the slice into a pretty-printed JSON byte array
	jsonOutput, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("_data/events.json", jsonOutput, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
