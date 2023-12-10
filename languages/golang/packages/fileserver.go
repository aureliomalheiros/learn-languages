package main

import (
    "fmt"
    "net/http"
    "log"
)

func main(){
    fmt.Println("Server running port 8000")
    fileServer := http.FileServer(http.Dir("./public"))
    mux := http.NewServeMux()
    mux.Handle("/", fileServer)
    log.Fatal(http.ListenAndServe(":8000", mux))

}
