package main

import (
	"fmt"
	"time"
)

//This is classic example Data Race Condition

var counter int

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			counter++ // Incrementing the counter without synchronization can lead to a race condition
		}()
		time.Sleep(time.Second) // Wait for the goroutines to finish - Fix the problem
		fmt.Printf("Counter: %d\n", counter) // The output may be unpredictable due to the race condition
	}
}
