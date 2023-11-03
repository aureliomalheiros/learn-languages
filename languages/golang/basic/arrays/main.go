package main

import "fmt"


//All values of the 
func main(){
	//Declaration with type, but without values
	var a[10] int
	//Declarations with type and inicialization values
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	//Declaration variable without type (Type Inference Declaration), but inilization values
	c := [2]float32{1.1, 2.1}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}