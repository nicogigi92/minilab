package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Function that sends a GET request to the kube-apiserver API
func makeRequest(clientset *kubernetes.Clientset, wg *sync.WaitGroup, mu *sync.Mutex, errCount *int, throttleCount *int, totalCount *int, responseTimes *[]time.Duration) {
	defer wg.Done()

	start := time.Now()
	_, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{}) // Request to list Nodes
	elapsed := time.Since(start)

	mu.Lock()

	*totalCount++ // Increment the total number of requests sent
	*responseTimes = append(*responseTimes, elapsed)
	if err != nil {
		*errCount++ // Increment the number of failed requests if an error occurs
	}

	if elapsed > time.Second {
		// 	If the response took more than one second, check if it's a case of throttling without an explicit error
		*throttleCount++ // Increment for silent throttling
	}

	mu.Unlock()
}

func calculateAverageResponseTime(responseTimes []time.Duration) float64 {
	if len(responseTimes) == 0 {
		return 0
	}

	total := 0.0
	for _, t := range responseTimes {
		total += t.Seconds()
	}
	return total / float64(len(responseTimes))
}

func main() {
	// Command-line arguments to configure the benchmark
	requestsPerSecond := flag.Int("rps", 10, "Number of requests per second")
	duration := flag.Int("duration", 10, "Benchmark duration in seconds")
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "Path to the kubeconfig file")
	flag.Parse()

	// Load the Kubernetes configuration
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	// Create the Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Variables to manage the benchmark
	var wg sync.WaitGroup
	var mu sync.Mutex
	errorCount := 0
	throttleCount := 0
	totalCount := 0
	var responseTimes []time.Duration
	ticker := time.NewTicker(time.Second / time.Duration(*requestsPerSecond))
	stopChan := make(chan bool)

	// Start the benchmark
	go func() {
		for i := 0; i < *duration; i++ {
			for j := 0; j < *requestsPerSecond; j++ {
				wg.Add(1)
				go makeRequest(clientset, &wg, &mu, &errorCount, &throttleCount, &totalCount, &responseTimes)
			}
			<-ticker.C
		}
		stopChan <- true
	}()

	// Wait for the benchmark to finish
	<-stopChan
	wg.Wait()

	averageResponseTime := calculateAverageResponseTime(responseTimes)

	// Benchmark results
	fmt.Printf("Benchmark finished!\n")
	fmt.Printf("Total number of requests sent: %d\n", totalCount)
	fmt.Printf("Total number of failed requests: %d\n", errorCount)
	fmt.Printf("Number of requests with throttling detected: %d\n", throttleCount)
	fmt.Printf("Average response time: %f\n", averageResponseTime)
}
