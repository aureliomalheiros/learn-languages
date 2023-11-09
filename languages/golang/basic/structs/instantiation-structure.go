package main

import "fmt"

type Person struct {
	Name 	string
	ID		int
}

func main(){
	var aurelio Person

	aurelio.Name = "Aurelio"
	aurelio.ID	 = 123123123

	fmt.Println(aurelio.Name)
	fmt.Println(aurelio.ID)
	fmt.Printf("%v\n", aurelio)
	
}