package main

import (
	"fmt"
	"context"
)

func main(){
	ctx := context.Background()

	fmt.Println("Context: \t, ctx")
	fmt.Println("Context err: \t", ctx.Err())
	fmt.Printf("Context type: \t%T\n", ctx)

}