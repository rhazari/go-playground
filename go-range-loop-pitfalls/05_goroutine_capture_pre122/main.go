package main

import (
	"fmt"
	"sync"
)

func main() {
	values := []string{"a", "b", "c"}

	var wg sync.WaitGroup

	fmt.Println("safe pattern: pass loop value as parameter")
	for _, v := range values {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(v)
	}
	wg.Wait()

	fmt.Println("Go 1.22+ also gives per-iteration variables with := in range,")
	fmt.Println("but this parameter pattern stays explicit and version-agnostic.")
}
