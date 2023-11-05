package main

import "fmt"

func createMap() map[string]int{
	return map[string]int{ "user1": 1000, "user2": 2000}
}

func main(){
	score := createMap()

	fmt.Println(score)
}