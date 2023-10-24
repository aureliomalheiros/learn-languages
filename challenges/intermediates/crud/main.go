package main


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"crud/server"
)
func main(){

	router := mux.NewRouter()
	// Method Create

	//Method POST
	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)
	//Method GET
	router.HandleFunc("/user", server.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.FindUser).Methods(http.MethodGet)
	//Method PUT
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)
	//Method DELETE
	router.HandleFunc("/user/{id}", server.DeleteUser).Methods(http.MethodDelete)


	fmt.Println("Listening port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	
}