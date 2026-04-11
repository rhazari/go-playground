package datastructures

import (
	"fmt"
	"sync"
	"time"
)

// Producer sends items to the queue (channel).
func producer(queue chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		queue <- i // Enqueue
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
	close(queue) // Close the channel when done producing
}

// Consumer receives items from the queue (channel).
func consumer(queue <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range queue { // Dequeue until channel is closed
		fmt.Printf("Consumed: %d\n", item)
		time.Sleep(200 * time.Millisecond) // Simulate processing
	}
}