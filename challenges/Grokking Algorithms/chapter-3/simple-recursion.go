package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        fmt.Println("Finish")
        return 1
    }
    return n * factorial(n - 1)
}

func main(){
    factorial(1000)
}
