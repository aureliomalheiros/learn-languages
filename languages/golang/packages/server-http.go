package main

import(
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Initialize my server in port 8000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Server"))
	})
	http.ListenAndServe(":8000", nil)
	
}