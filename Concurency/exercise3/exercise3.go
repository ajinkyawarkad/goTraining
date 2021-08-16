// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.
package main

import (
	"fmt"
	"sync"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var m sync.Mutex

	go func() {
		for i := 0; i < 1000; i++ {
			m.Lock()
			scores["A"]++
			m.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			m.Lock()
			scores["B"]++
			m.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores)
}
