package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendRequest(client *http.Client, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Response: %d\n", resp.StatusCode)
}

func stressTest(url string, requestsPerSecond int, duration int) {
	client := &http.Client{}

	interval := time.Duration(1e9 / requestsPerSecond)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var wg sync.WaitGroup

	start := time.Now()
	for time.Since(start).Seconds() < float64(duration) {
		select {
		case <-ticker.C:
			wg.Add(1)
			go sendRequest(client, url, &wg)
		}
	}

	wg.Wait()
}

func main() {
	url := flag.String("url", "", "Server URL to stress")
	rps := flag.Int("rps", 10, "Number of request per seconds")
	duration := flag.Int("duration", 10, "Duration in seconds")

	flag.Parse()

	if *url == "" {
		fmt.Println("Please specify an url with --url")
		return
	}

	fmt.Printf("Lauching stress test on %s wwith %d rps during %d seconds.\n", *url, *rps, *duration)
	stressTest(*url, *rps, *duration)
}
