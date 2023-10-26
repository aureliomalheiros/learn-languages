package main

import "fmt"

func main() {
	a := 20
	fmt.Println("Initial value a")
	fmt.Println(a)
	fmt.Println(&a)

	//* pointer to address a 
	var pointer *int = &a

	fmt.Println(pointer)

	//Value of the pointer == 20
	*pointer = 25

	fmt.Println(a)
	fmt.Println(pointer)

	//Value b
	b := &a
	fmt.Println("Initial value b")
	fmt.Println(*b)
	*b = 30
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(a)
}