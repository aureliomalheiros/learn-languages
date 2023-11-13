package main

import "fmt"

type Number interface {
	int | float32
}

func Operation[T Number](numbers [5]T) T {
	var mathOp T 

	for _, value := range numbers{
		mathOp += value
	}
	return mathOp
}

func main(){
	arrayInt 	:= [5]int{1, 2, 3, 4, 5}
	arrayFloat 	:= [5]float32{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(Operation(arrayInt))
	fmt.Println(Operation(arrayFloat))
}