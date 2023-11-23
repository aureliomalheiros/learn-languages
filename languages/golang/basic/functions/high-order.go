package main

import "fmt"

func apply(numbers []int, f func(int) int) []int{
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = f(n)
	}
	return result
}

func main(){
	mySliceNums := []int{1, 2, 3, 4, 5}
	//Anonymous func

	doubled := apply(mySliceNums, func(n int) int {
		return n * 2
	})

	fmt.Println(doubled)
}