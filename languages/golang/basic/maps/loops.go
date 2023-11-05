package main

import "fmt"

func main(){
	months := make(map[int]string)

	myMonths := [12]string {
		"January", "February", "March",
		"April", "May", "June",
		"July", "August", "September",
		"Octuber", "November", "December",
	}

	for number, myMonth := range myMonths{
		months[number+1] = myMonth
	}
	
	for numberMonth, month := range months{
		fmt.Printf("Month: %s is number: %d\n", month, numberMonth)
	}
}