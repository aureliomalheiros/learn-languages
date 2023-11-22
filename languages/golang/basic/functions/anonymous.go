package main

import "fmt"

func main(){
	myValue := 2
	x := 3
	func(x int){
		fmt.Println(x*myValue)
	}(x)

	y := 234234234
	
	func(y int){
		fmt.Println(y*y)
	}(y)
	
}