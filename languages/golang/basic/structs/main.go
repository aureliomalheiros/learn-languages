package main

import "fmt"

type Test struct {
	Name string
	Age int
	Salary, Number int
}

func main(){
	aurelio := Test{
		Name: "Aurelio", 
		Age: 30,
	}

	fmt.Println(aurelio)
	fmt.Println(aurelio.Name)
	fmt.Println(aurelio.Age)
}