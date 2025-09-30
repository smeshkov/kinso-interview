package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
)

// EventDTO represents the desired JSON structure.
type EventDTO struct {
	EventID   string `json:"eventId"`
	UserID    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	Source    string `json:"source"`
	RawData   string `json:"rawData"`
}

// generateEvent creates a single EventDTO with mock data.
func generateEvent(userIDs []string) EventDTO {
	sources := []string{"Gmail", "Slack", "WhatsApp", "LinkedIn"}

	// Use a UUID for the unique EventID
	eventID := uuid.New().String()

	// Randomly select a UserID from the pre-generated pool
	userID := userIDs[rand.Intn(len(userIDs))]

	// Generate a recent, but random, UTC timestamp in ISO 8601 format
	now := time.Now().UTC()
	randomOffsetSeconds := time.Duration(rand.Intn(86400*7)) * time.Second // Up to 7 days in the past
	createdAt := now.Add(-randomOffsetSeconds).Format("2006-01-02T15:04:05.000000Z")

	// Select a random source
	source := sources[rand.Intn(len(sources))]

	// Generate RawData as a JSON string
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
	const totalEvents = 100
	const uniqueUsers = 20 // Number of unique user IDs to generate

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a pool of a few unique user IDs to ensure repetition
	userIDs := make([]string, uniqueUsers)
	for i := 0; i < uniqueUsers; i++ {
		userIDs[i] = uuid.New().String()[:8] // Truncate for a simpler look
	}

	// Create a slice to hold the events
	events := make([]EventDTO, totalEvents)

	// Populate the slice
	for i := range totalEvents {
		events[i] = generateEvent(userIDs)
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
