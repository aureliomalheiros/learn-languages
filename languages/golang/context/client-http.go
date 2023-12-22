package main

import (
	"net/http"
	"time"
	"io"
	"os"
	"context"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req,  err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}