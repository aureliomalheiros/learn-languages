## Convert values


In the example `mismatched-types-int16-and-int32.go` a solution:

```go
package main

import "fmt"

func main (){

    var x int16 = 1
    var y int32 = 2
    var z int   = int(x) + int(y)

    fmt.Println(z)

}
```


The problem to convert float to int is to do the following `int(x)`, you can see that it discards all values after de comma.
