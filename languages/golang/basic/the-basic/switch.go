package main

import "fmt"

func main(){
	x := 1

	switch {
		case x > 1:
			fmt.Println("maior que 1")
		case x > 2:
			fmt.Println("maior que 2")
		case x == 1:
			fmt.Println("Agora  sim")
		default:
			fmt.Println("Sem isso")
	}

	switch y := 2; {
	case x > y:
		fmt.Println("Não foi")
	case x < y:
		fmt.Println("Agora foi, mas no outro")
	default:
		fmt.Println("Sem isso")
	}
}