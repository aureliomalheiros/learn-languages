package main

import "fmt"

var i int
var f float64
var s string
var b bool

type TestingStruct struct {}

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
