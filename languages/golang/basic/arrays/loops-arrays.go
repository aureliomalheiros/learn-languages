package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 4, 5, 6}

	for index, number := range numbers {
		fmt.Printf("Index: %d and number: %d\n", index, number)
	}
}