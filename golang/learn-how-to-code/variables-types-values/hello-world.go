package main

import "fmt"

func main(){
    n,_ := fmt.Println("Hello World")
	fmt.Println(n)
    call()
    
}
// Call function

func call(){
	fmt.Println("Call func")
}
