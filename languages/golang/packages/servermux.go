package main

import (
    "fmt"
    "net/http"
)

type blog struct{}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Blog"))
}

func main(){
    
    fmt.Println("Server Running in ports 8000 and 8001")

    mux := http.NewServeMux()
    mux.HandleFunc("/",  HelloHandler)
    mux.Handle("/blog", blog{})
    http.ListenAndServe(":8000", mux)

}


func HelloHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello World!"))
}

