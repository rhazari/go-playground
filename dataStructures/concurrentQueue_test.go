package datastructures

import (
	"fmt"
	"sync"
	"testing"
)

func Test_ConcurrentQueue(t *testing.T) {
	queue := make(chan int, 3) // Buffered channel (queue size 3)
	var wg sync.WaitGroup

	// Start producer
	wg.Add(1)
	go producer(queue, &wg)

	// Start consumer
	wg.Add(1)
	go consumer(queue, &wg)

	wg.Wait() // Wait for both to finish
	fmt.Println("Done")
}