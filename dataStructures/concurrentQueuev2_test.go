package datastructures

import (
	"fmt"
	"sync"
	"testing"
)

func Test_ConcurrentQueuev2(t *testing.T) {
	q := NewConcurrentQueue()
	var wg sync.WaitGroup

	// Start producer
	wg.Add(1)
	go Producer(q, &wg)

	// Start consumer
	wg.Add(1)
	go Consumer(q, &wg)

	wg.Wait()
	fmt.Println("Done")
}