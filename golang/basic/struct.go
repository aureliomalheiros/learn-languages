package main

import "fmt"

type people struct {
	name string
	age int
}

// Embedding
type studant struct {
	people // embedding struct people in the studant
	enrol int
}
func main (){
	person1 := people {
		name: "Person1",
		age: 30,
	} 
	person2 := people {
		name: "Person2",
		age: 26,
	}
	person3 := studant {
		people: people{
			name: "Person3",
			age: 23,
		},
		enrol: 1234,
	}
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	//To display only one entry of a struct

	fmt.Println(person1.age)
	fmt.Println(person1.name)
	fmt.Println("My name is ", person1.name,"and I have ", person1.name, "years old" )
}