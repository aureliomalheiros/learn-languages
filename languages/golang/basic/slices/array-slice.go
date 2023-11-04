package main

import "fmt"

func main(){
	week := [...]string{
		"Monday", "Tuesday", 
		"Wednesday", "Thursday", 
		"Friday", "Saturday", 
		"Sunday"}
	
	for index, dayWeek := range week {
		fmt.Printf("Index: %d Day Week %s\n", index, dayWeek)
	}
	
	fmt.Println("\n\n")

	businessWeek 	:= week[0:5]
	weekend			:= week[5:] 

	fmt.Printf("Lenght: %d, Capacity: %d Array: %v\n",  len(businessWeek), cap(businessWeek), businessWeek)
	fmt.Printf("Lenght: %d, Capacity: %d Array: %v\n",  len(weekend), cap(weekend), weekend)
}