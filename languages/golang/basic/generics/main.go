package main

import "fmt"

func Sum[Generic int | float32](m map[string]Generic)Generic {
	var sum Generic
	for _, v := range m {
		sum += v
	}
	return sum
}

func main(){
	m := map[string]int{"User1": 1000, "User2": 2000}
	m2 := map[string]float32{"User3": 3000, "User4": 4000}

	fmt.Println(Sum(m))
	fmt.Println(Sum(m2))
}