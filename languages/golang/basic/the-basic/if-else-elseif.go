package main

import "fmt"

func main(){
	x := 1
	y := 1
	z := 2

	if x == z {
		fmt.Println("OK")
	} else if x == y && y != z {
		fmt.Println("Hum, Ok")
		if x < z {
			fmt.Println("Yeap")
		}
		if x > y || x != z {
			fmt.Println("Very good")
		}
	} else {
		fmt.Println("Finish")
	}
}