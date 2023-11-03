package main

import "fmt"

func main(){
	var myArray [5]int = [5]int{1, 2, 3, 4, 10}
	//Print last value of the array
	fmt.Println(myArray[len(myArray)-1])
	//Print number of elements in the array
	fmt.Println(len(myArray))
}