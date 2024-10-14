package main

import (
	"fmt"
	"github.com/learn-languages/languages/golang/packages/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}