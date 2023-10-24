package main

import "fmt"

func factorialTail(n int, acc int) int {
    if n == 0 {
        return acc
    }
    fmt.Println(n, acc)
    return factorialTail(n - 1, n * acc)
}

func main (){
    factorialTail(10, 10)
}
