package main

import "fmt"

//When use pointers function
//Following example

func corretSum(c, d int) int {
	return c + d
}

//In this function I modify value
func incorretSum(e, f int) int {
	e = 50
	return e + f
}

//Using pointers
func pointerSum(g, h *int) int {
	*g = 100
	return *g + *h
}

func main(){
	a := 10
	b := 20

	corretSumFunction := corretSum(a, b)
	fmt.Println(corretSumFunction)

	incorretSumFunction := incorretSum(a, b)
	fmt.Println(incorretSumFunction)

	pointerSumFunction := pointerSum(&a, &b)
	fmt.Println(pointerSumFunction)

	//Print values 
	fmt.Printf("Values: %v %v\n", a, b)
}