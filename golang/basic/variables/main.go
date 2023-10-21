package main

import "fmt"

var i int
var f float64
var s string
var b bool

type TestingStruct struct {}

var x int = 1
var y int32 = 2
var z int =  x + y

//VAMOS TRABALHAR COM ZERO EM GOLANG
func main(){
    fmt.Printf("%v %v %v %q\n", i, f, b, s) 
    fmt.Println(s)
    var value int
    testing(value)
    fmt.Println(TestingStruct)
}
func testing(test int){
    fmt.Printf("%v\n", test)
}
