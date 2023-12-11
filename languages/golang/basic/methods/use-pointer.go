package main

import "fmt"

type User struct{
	Name	string
	Age 	int
	Status	bool
}

func (u *User) MyUser(){
	u.Name = "Teste"
}

func (u User) AgeUser(){
	u.Age = 21
}
func main(){
	aurelio := User{
		Name: "Aurelio",
		Age: 31,
		Status: true,
	}

	aurelio.MyUser()
	fmt.Println(aurelio.Name)

	aurelio.AgeUser()
	fmt.Println(aurelio.Age)

}