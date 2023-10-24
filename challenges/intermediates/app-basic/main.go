package main

import (
	"log"
	"os"
	"line-command/app"
)

func main() {
	application := app.Generation()	
	if erro :=  application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
