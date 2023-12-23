package main

import (
	"fmt"
	"context"
)

func main(){
	ctx := context.WithValue(context.Background(), "language", "Go")

	fooValue(ctx, "language")

	ctx = context.WithValue(ctx, "dog", "Gaston")
	fooValue(ctx, "dog")

	fooValue(ctx, "color")
}

func fooValue(ctx context.Context, s string){
	if v := ctx.Value(s); v != nil {
		fmt.Println("Found value;", v)
		return
	}
	fmt.Println("Key not found: ", s)
}