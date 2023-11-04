package main

import "fmt"

func main() {
	sliceTest := []int{10,20,30,40,50}

	fmt.Printf("Len: %d, Capacity: %d and Array: %v\n", len(sliceTest), cap(sliceTest),sliceTest)

	fmt.Printf("Len: %d, Capacity: %d and Array: %v\n", len(sliceTest[:0]), cap(sliceTest[:0]),sliceTest[:0])
	
	fmt.Printf("Len: %d, Capacity: %d and Array: %v\n", len(sliceTest[0:3]), cap(sliceTest[0:3]),sliceTest[0:3])

	fmt.Printf("Len: %d, Capacity: %d and Array: %v\n", len(sliceTest[:4]), cap(sliceTest[:4]),sliceTest[:4])

	fmt.Printf("Len: %d, Capacity: %d and Array: %v\n", len(sliceTest[:5]), cap(sliceTest[0:5]),sliceTest[:5])
	
}