package main

import "fmt"

func main(){
	Worker := struct {
		Name string
		ID int
	}{
		Name: "Aurelio",
		ID: 2,
	}
	fmt.Println(Worker)
}