package main

import "fmt"

func main(){
	x := 10

	expressionFunc := func(x int) int{
		return x * 10
	}
	fmt.Printf("%v\n", expressionFunc(x))
	fmt.Println(x, expressionFunc(x))
}