package main

import "fmt"

func returnFunc() func() int{
	var x int
	return func() int{
		x++
		return x * x
	}
}

func main(){
	f := returnFunc()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}