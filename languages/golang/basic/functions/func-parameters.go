package main

import "fmt"

//Without return
func parametrFuncWithoutReturn(x, y int) {
	fmt.Println(x + y)
}

//With return
func parametrFuncWithReturn(x int, y int) int{
	return x + y
}

//With many returns
func funcManyReturns(x, y int,  z, w float32) (int, float32){
	return x + y, z + w
}

func main(){
	parametrFuncWithoutReturn(2, 2)

	fmt.Println(parametrFuncWithReturn(3, 3))

	fmt.Println(funcManyReturns(2, 3, 3.2, 2.1))
}