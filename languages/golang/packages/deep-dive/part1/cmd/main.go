package main

import (
	"fmt"
	"github.com/aureliomalheiros/learn-languages/languages/golang/packages/deep-dive/math"
)

func main(){
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}