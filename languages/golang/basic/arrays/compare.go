package main

import "fmt"

func main(){
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{2, 1}
	//d := [2]float32{2.1, 2.3}	

	//It is not possible compare int with float
	/*
	./erros.go:15:8: invalid operation: 
	c == d (mismatched types [2]int and [2]float32)
	*/

	fmt.Println(
		a == b, 
		a == c, 
		b == c,
	)

}