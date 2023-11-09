package main

import "fmt"

type Triangle struct {X, Y, Z int}


func main(){
	p := Triangle{1, 2, 3}
	fmt.Printf("%v\n", p)
}