package main

import (
	"fmt"
	"log"
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

	for _, i := range New(0, noOfEvents) {
		log.Println("Processing event", i)
		events = append(events, Event{
			Timestamp: time.Now(),
			Payload:   fmt.Sprintf("Event %d", i),
		})
	}

	start := time.Now()

	for _, event := range events {
		log.Println("Querying event: ", event.Payload)

		if event.Timestamp.Before(time.Now()) {
			if time.Since(start) > time.Second {
				log.Println("Querytime exceeded threshold")
				break
			}
		}
	}

	log.Println("Success")
}
