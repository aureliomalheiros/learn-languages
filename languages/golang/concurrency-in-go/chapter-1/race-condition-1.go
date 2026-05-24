package main

import (
	"time"
	"fmt"
)

func main() {
	
	var data int

	go func() {
		data++
	}() //Immediately Invoked Function Expression (IIFE)

	time.Sleep(1 * time.Second) // Wait for the goroutine to finish - Fix the problem

	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	}else {
		fmt.Printf("The value is %v.\n", data)
	}
}
