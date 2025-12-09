package main

import (
    "fmt"
    "os"
    "strconv"
    "backend-monorepo/math"
)

func main() {

    AStr := os.Getenv("A")
    BStr := os.Getenv("B")
    
    A, err := strconv.Atoi(AStr)
    if err != nil {
        fmt.Println("A must be an interger")
        return 
    }

    B, err := strconv.Atoi(BStr)
    if err != nil {
        fmt.Println("B must be an interger")
        return
    }

    m := math.Math{A: A, B: B}
    fmt.Println(m.Add())
}
