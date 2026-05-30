package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex  	sync.Mutex
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock() // Lock the mutex to ensure exclusive access to the counter variable
			counter++
			mutex.Unlock() // Unlock the mutex after modifying the counter variable
		}()
	}
	fmt.Printf("Counter: %d\n", counter) // The output may be unpredictable due to the race condition
}
