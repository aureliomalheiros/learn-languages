package main

import "fmt"

type Worker struct {
    Name   string
    Salary int
}

func main() {
    var aurelio Worker
    aurelioPointer := &aurelio

    fmt.Println(aurelioPointer)
}