package main

import "fmt"

func main(){
	age := map[string]int {
		"Aurelio": 31,
		"Ana Carolina": 27,
		"Bob": 0,
	}
	age["Bob"] += 1
	fmt.Println(age)
	fmt.Printf("%v\n", age["Aurelio"])
}