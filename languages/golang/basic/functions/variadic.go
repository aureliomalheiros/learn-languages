package main

import "fmt"

func variadicFunc(values ...int) int {
	sum := 0

	for _, value := range values{
		sum += value
	}
	return sum
}

func main(){
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(variadicFunc(mySlice...))
	fmt.Println(variadicFunc(1, 2))
	fmt.Println(variadicFunc(1, 2, 3))
	fmt.Println(variadicFunc(1, 2, 3, 4))
	fmt.Println(variadicFunc(1, 2, 3, 4, 5))
}