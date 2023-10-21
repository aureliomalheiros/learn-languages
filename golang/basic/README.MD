# Study notes about Golang

## Basic condepts

### Reserved Words

```
break     default      func   interface   select
case      defer        go     map         struct 
chan      else         goto   package     switch
const     fallthrough  if     range       type
continue  for          import return      var
```

### Types 

```
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
```


### Binary Operators

```
* / % << >>& &^
+ - | ^
== ≠ < <= > >=
&&
||
```
> ![Important]
> The operator % is applied only to integer numbers.

---
## Hello World

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

### Starting with the basic


- It's possible use values _signed_ and _unsigned_.
    -  int8, int16, int32 e int64 -> 

- In Golang, it is not allowed to declared variable and not use it.

_This is good! It reduces clutter in the code_

## Value Zero
`value-zero.go`


In my studies about zero:

- When not assign value the variable in Golang assign zero to it;
- Values `bools` adding of the value `false` [false = 0 e true = 1]
- Em variaveis sem retorno precisamos usar
- An uninitialized string in Go comes with the value: ""