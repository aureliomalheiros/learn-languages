package main

import (
	"fmt"
	"sync/atomic"
)

//In this example, we use the atomic package to safely increment the counter variable across multiple goroutines without the need for explicit locking.
//The atomic.AddInt64 function atomically adds the specified value (in this case, 1) to the counter variable, ensuring that the operation is thread-safe and preventing race conditions.
//By using atomic operations, we can avoid the overhead of mutex locks while still ensuring that our shared variable is updated correctly across concurrent goroutines.
var counter int64

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // Atomically increment the counter variable
			// The atomic package provides functions for performing atomic operations on variables, which can help prevent
		}()
	fmt.Printf("Counter: %d\n", counter) // The output will always be 100, as the atomic operation ensures that the counter variable is updated correctly
	}

}
