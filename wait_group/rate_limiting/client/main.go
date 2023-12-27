package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup

	// Iterate through the total number of requests in batches of 'max'
	for i := 0; i < total; i += max {
		limit := max
		if i+max > total {
			limit = total - i
		}

		// Add the number of goroutines that will be spawned in this batch to the WaitGroup
		wg.Add(limit)

		// Launch goroutines to make network requests
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()

				// Establish a TCP connection to localhost on port 8080
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("could not dial: %v", err)
				}

				// Read data from the connection
				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("could not read from conn: %v", err)
				}

				// Check if the received data is "success"; if not, log an error
				if string(bs) != "success" {
					log.Fatalf("request error, request: %d", i+1+j)
				}

				// If the received data is "success", print a success message
				fmt.Printf("request %d: success\n", i+1+j)
			}(j)
		}

		// Wait for all goroutines in this batch to complete before moving to the next batch
		wg.Wait()
	}
}
