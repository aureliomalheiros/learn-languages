package main

import "fmt"

type Person struct {
	Name string
	Age int
	ID int
}

type Worker struct {
	Person
	Function string
}

func main(){
	person1 := Worker{
		Person: Person{
			Name: "person1", 
			Age: 30,
			ID: 10001000,
		},
		Function: "SRE",
	}

	person2 := Person {
		Name: "Person2",
		Age: 29,
		ID: 20002000,
	}

	fmt.Println(person1)
	fmt.Println(person2)
}