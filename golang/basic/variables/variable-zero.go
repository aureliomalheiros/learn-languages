package main

import "fmt"


var x int
var y bool
var z float32
var w string

type TestingStruct struct {}


func main(){
	fmt.Printf("x= %v,y= %v,z= %v,w= %v\n", x, y, z, w)
	fmt.Printf("String value: %q\n", w)

	testStruct := TestingStruct{}

	fmt.Println(testStruct)
}