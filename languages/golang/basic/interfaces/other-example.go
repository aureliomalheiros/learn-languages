package main

import "fmt"

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string{
	return d.Name + "au au"
}

func (c Cat) Speak() string{
	return c.Name + "miau miau"
}


type Animal interface{
	Speak() string
}

//Methods
func AnimalSpeak(a Animal) string{
	return a.Speak()
}

func main(){
	dog_1 := Dog{"Bob"}
	dog_2 := Dog{"Trup"}


	cat_1 := Cat{"Chico"}
	cat_2 := Cat{"Bill"}

	fmt.Println(AnimalSpeak(dog_1),"/n")
	fmt.Println(AnimalSpeak(dog_2),"/n")
	fmt.Println(AnimalSpeak(cat_1),"/n")
	fmt.Println(AnimalSpeak(cat_2),"/n")


}