package main

import "fmt"

type User struct {
	Name 		string
	Username 	string
	Status 		bool
}

type UserStatus interface {
	Desactive()
}

func (u *User) Desactive(){
	u.Status = false
}

func main(){
	test := User{
		Name: "My_User",
		Username: "myuser",
		Status: true,
	}

	fmt.Println(test)

	test.Desactive()
	
	fmt.Printf("%v\n", test.Status)
}