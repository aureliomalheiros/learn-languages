package main

import "fmt"

var x int  //not using the variable
var y int = 1

// it's possible 

var (
    a int       = 2
    b string    = "aureli"
    c float32   = 3.2
    d bool      = true
)

func main(){    
    //short variable declaration
    //Use in declaration local variables

    myName  := "Aurelio Malheiros"
    date    := 31

    fmt.Println(myName)
    fmt.Println(date)

    //Multiple variabels
    
    s, t, u := 2, true, "buenas, my friend"

    fmt.Printf("%v\n%v\n%s\n", s, t, u)
    fmt.Println(y)
}