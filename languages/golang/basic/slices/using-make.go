package main

import "fmt"

func main(){
	makeSlice := make([]int, 9, 10)
	// The length cannot be great than the capacity
	fmt.Printf("Lenght: %d, Capacity: %d, Array: %v\n", len(makeSlice), cap(makeSlice), makeSlice)
}