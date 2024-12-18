package main

import (
	"fmt"
	"time"
)

type Event struct {
	Timestamp time.Time
	Payload   string
}

func New(start, end int) []int {
	result := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		result[i-start] = i
	}

	return result
}

func main() {
	events := []Event{}
	noOfEvents := 1000000

	// Simulating writing 1 million events
	for _, i := range New(0, noOfEvents) {
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
