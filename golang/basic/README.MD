# Study notes about Golang

## Basic condepts

---
### Simple package and how use

Default all project existing file `main.go` in this file has the root of the project.


```go
package main

import "fmt"

func main(){
    fmt.Println("Buenas, my friend")
}
```

### Line for line

```package main``` 

All file in Golang exist package, but the package *main* exist only in the root.

[Example](simple-packages-and-main)

```import fmt```

I use print message in the terminarl, the function that does thie is **fmt.Println("")**


```func main() {}```

Here  is the function that every core file should have.


```fmt.Println("Buenas, my friend")```

And finish, function of the package "fmt" and println phrase in the terminal.

> [!Important]
> When not declared package main in my folder root?
> 

_When not use function main in the file root_


```bash
$ go run main.go 
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
``````
```

_When not use package main in the files root_

```bash
$ go run main.go
package command-line-arguments is not a main package
```


---

## Variables

