package main

import (
	"fmt"
	"sync"
)

//This example I demonstrates a race condition
//The main goroutine and the anonymous goroutine are accessing the shared variable data without any synchronization mechanism, which can lead to unpredictable results. Depending on the scheduling of the goroutines, the output may be either "The value is 0." or "The value is 1.", and it may change between runs.

func main() {
	var memoryAccess sync.Mutex
	var value int

	go func() { 
		 memoryAccess.Lock() // Lock the mutex to ensure exclusive access to the shared variable
		 value++
		 memoryAccess.Unlock() // Unlock the mutex after modifying the shared variable
	}() // Immediately Invoked Function Expression (IIFE)

	memoryAccess.Lock() // Lock the mutex to ensure exclusive access to the shared variable
	if value == 0 {
		fmt.Printf("The value is %v.\n", value)
	} else {
		fmt.Printf("The value is %v.\n", value)
	}
	memoryAccess.Unlock() // Unlock the mutex after reading the shared variable
}
