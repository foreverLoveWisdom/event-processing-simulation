package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendRequest(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	_, err := client.Get("http://localhost:8080/api")
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
}

func main() {
	var wg sync.WaitGroup
	client := &http.Client{Timeout: time.Second * 10}

	noOfRequests := 25000
	start := time.Now()

	for i := 0; i < noOfRequests; i++ {
		wg.Add(1)

		go sendRequest(&wg, client)
	}

	wg.Wait()
	fmt.Printf("Completed %d requests in %v\n", noOfRequests, time.Since(start))
}

// figure out why what does timeout for the client and server means
// why timeout server
// why all 12 cores CPU 100% when time out
// what is going here
// good opportunity to practice I/O and concurrency
