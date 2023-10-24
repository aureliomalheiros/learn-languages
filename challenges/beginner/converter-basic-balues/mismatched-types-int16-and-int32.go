package main

import "fmt"

func main(){
    var x int16 = 1
    var y int32 = 2
    var z int = x + y //solution var z int = int(x) + int(y)
    
    fmt.Println(z)
}
