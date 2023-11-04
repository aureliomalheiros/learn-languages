package main

import "fmt"

func main(){
	var mySlice 	[]int
	var otherSlice 	[]int

	mySlice = append(mySlice, 10)

	fmt.Printf("length: %d, capacity: %d, array: %v\n", len(mySlice), cap(mySlice), mySlice)

	mySlice = append(mySlice, 40)

	fmt.Printf("length: %d, capacity: %d, array: %v\n", len(mySlice), cap(mySlice), mySlice)

	mySlice = append(mySlice, 50)

	fmt.Printf("length: %d, capacity: %d, array: %v\n", len(mySlice), cap(mySlice), mySlice)
	
	mySlice = append(mySlice, 60)

	fmt.Printf("length: %d, capacity: %d, array: %v\n", len(mySlice), cap(mySlice), mySlice)

	mySlice = append(mySlice, 70)

	fmt.Printf("length: %d, capacity: %d, array: %v\n", len(mySlice), cap(mySlice), mySlice)

	for x := 0; x <= 100; x+=10 {
		otherSlice = append(otherSlice, x)
	}
	fmt.Printf("length: %d, capacity: %d, array: %v\n", len(otherSlice), cap(otherSlice), otherSlice)

}