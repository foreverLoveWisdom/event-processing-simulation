package main

import (
	"fmt"
	"time"

	"github.com/golangci/golangci-lint/pkg/golinters/intrange"
)

type Event struct {
	Timestamp time.Time
	Payload   string
}

func main() {
	events := []Event{}

	// Simulating writing 1 million events
	for _, i := range intrange.New(0, 1000000) {

		fmt.Println("Processing event", i)
		events = append(events, Event{
			Timestamp: time.Now(),
			Payload:   fmt.Sprintf("Event %d", i),
		})
	}

	// Simulating querying for an event by timestamp(e.g, range query)
	start := time.Now()

	for _, event := range events {
		fmt.Println("Querying event: ", event.Payload)

		if event.Timestamp.Before(time.Now()) {
			// Simulating processing the event
			if time.Since(start) > time.Second {
				fmt.Println("Querytime exceeded threshold")
				break
			}
		}
	}

	// Print a success message
	fmt.Println("Success")
}
