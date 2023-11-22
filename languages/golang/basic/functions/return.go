package main

import "fmt"

func returnFunc() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func returnString(name string, year, date int) (string, int) {
	return func(name string, year, date int) (string, int) {
		return name, year - date
	}(name, year, date)
}

func main () {
	myReturn := returnFunc()

	x := myReturn(1, 2)

	fmt.Println(x)

	fmt.Println(returnString("Bob Drogado", 2023, 1980))
}