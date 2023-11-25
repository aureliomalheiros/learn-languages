package main

import(
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Initialize my server in port 8000")
	http.ListenAndServe(":8000", nil)
	
}