package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
)

func main(){
	fmt.Println("Server running in port :8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Println("Starting Request...")
	defer log.Println("Finish request.")
	select {
	case <- time.After(5 * time.Second):
		log.Println("Processed request")
		w.Write([]byte("Processed request"))
	case <- ctx.Done():
		log.Println("Cancel request!")
		http.Error(w, "Client cancelled request", http.StatusRequestTimeout)
	}
}