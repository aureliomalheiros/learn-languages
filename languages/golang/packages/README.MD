# Packages

## net/http

The package is many used in projects with language Golang, I write the basic concepts.

Here is a simple example of an HTTP server in Go:

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8000", nil)
}
```

I created function `helloHandler` thats writes "Hello World!" to the HTTP response. In function main its possible see route "/" using `http.HandleFunc`. I finished, using `http.ListenAndServe` and server running in port 8000

**Whats handler?**

In package `net/http` handler is an object that implements the `http.Handler` interface. This is interfae requires a method called ServerHTTP which takes a `ResponseWriter` and a `Request` as parameters.

The diferente between these two lines of code lies:


```go
http.ListenAndServe(":8000", nil)
```

> The argument is nil, which means the server will use the default request multiplex from the `net/http` package.

```go
http.ListenAndServe(":8000", mux)
```

> The arugment is mux, which is an instance of `http.ServeMux`. This means the server will use the mux request multiplex to route request to diferente handlers based on the request URL.


[net/http](https://pkg.go.dev/net/http)
