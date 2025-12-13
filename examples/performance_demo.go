package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/tonnytg/webreq"
)

func main() {
	fmt.Println("=== WebReq Performance Demo ===\n")

	// Demo 1: Simple GET request
	fmt.Println("1. Simple GET Request:")
	simpleGet()
	fmt.Println()

	// Demo 2: Sequential requests showing connection reuse
	fmt.Println("2. Sequential Requests (connection reuse):")
	sequentialRequests()
	fmt.Println()

	// Demo 3: Concurrent requests showing connection pooling
	fmt.Println("3. Concurrent Requests (connection pooling):")
	concurrentRequests()
	fmt.Println()

	// Demo 4: Custom context with timeout
	fmt.Println("4. Custom Context with Timeout:")
	contextWithTimeout()
	fmt.Println()

	// Demo 5: Request cancellation
	fmt.Println("5. Request Cancellation:")
	requestCancellation()
}

func simpleGet() {
	start := time.Now()

	req := webreq.NewRequest("GET")
	req.SetURL("https://httpbin.org/get")
	req.SetTimeout(10)

	_, err := req.Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("   Request completed in %v\n", time.Since(start))
	fmt.Printf("   Status code: %d\n", req.StatusCode)
}

func sequentialRequests() {
	start := time.Now()
	count := 5

	for i := 0; i < count; i++ {
		req := webreq.NewRequest("GET")
		req.SetURL("https://httpbin.org/get")
		req.SetTimeout(10)

		_, err := req.Execute()
		if err != nil {
			log.Printf("Request %d error: %v\n", i+1, err)
			continue
		}
		fmt.Printf("   Request %d completed (Status: %d)\n", i+1, req.StatusCode)
	}

	elapsed := time.Since(start)
	fmt.Printf("   Total time for %d requests: %v\n", count, elapsed)
	fmt.Printf("   Average per request: %v\n", elapsed/time.Duration(count))
}

func concurrentRequests() {
	start := time.Now()
	count := 10
	var wg sync.WaitGroup
	results := make(chan int, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			req := webreq.NewRequest("GET")
			req.SetURL("https://httpbin.org/get")
			req.SetTimeout(10)

			_, err := req.Execute()
			if err != nil {
				log.Printf("Request %d error: %v\n", num, err)
				return
			}
			results <- req.StatusCode
		}(i + 1)
	}

	wg.Wait()
	close(results)

	successCount := 0
	for range results {
		successCount++
	}

	elapsed := time.Since(start)
	fmt.Printf("   Completed %d concurrent requests in %v\n", successCount, elapsed)
	fmt.Printf("   Average per request: %v\n", elapsed/time.Duration(successCount))
}

func contextWithTimeout() {
	// Create a context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()

	req := webreq.NewRequest("GET")
	req.SetURL("https://httpbin.org/delay/2") // Endpoint that delays 2 seconds

	_, err := req.ExecuteWithContext(ctx)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("   Request completed in %v\n", time.Since(start))
	fmt.Printf("   Status code: %d\n", req.StatusCode)
}

func requestCancellation() {
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after 500ms
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("   Cancelling request...")
		cancel()
	}()

	start := time.Now()

	req := webreq.NewRequest("GET")
	req.SetURL("https://httpbin.org/delay/5") // Endpoint that delays 5 seconds

	_, err := req.ExecuteWithContext(ctx)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("   Request cancelled after %v\n", elapsed)
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Request completed in %v (unexpected)\n", elapsed)
	}
}
