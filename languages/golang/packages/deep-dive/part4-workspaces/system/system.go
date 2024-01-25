package main

import (
	"fmt"
	"github.com/aureliomalheiros/learn-languages/languages/golang/packages/deep-dive/part4-workspaces/math"
)

func main(){
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}