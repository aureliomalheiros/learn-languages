package main

import "fmt"

type User struct {
	Name string
	Age int
	Status bool
}

func (u User) Active(){
	u.Status = true
	fmt.Printf("User %s active\n", u.Name)
}
func main(){
	aurelio := User {
		Name: "Aurelio", 
		Age: 31,
		Status: false,
	}

	aurelio.Active()
}