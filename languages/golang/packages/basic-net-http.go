package main

import (
    "net/http"
    "io"
    "fmt"
)

func main() {
    request, err := http.Get("https://google.com.br")
    
    if err != nil{
        panic(err)
    }

    result, err := io.ReadAll(request.Body)
    
    if err != nil {
        panic(err)
    }

    fmt.Println(string(result))
    
    //Close request, because we can use many resource

    request.Body.Close()
}
