# Study notes about Golang

## Basic condepts

### Camel Case

[Camel Case](https://en.wikipedia.org/wiki/Camel_case)

Golang use Camel Case to style.

> ![Important] Style
> Best Practices 
> - QuoteRuneToASCII or parseRequestLine
> Not recomend
> quote_rune_to_ASCII or parse_request_line


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
> _When not use function main in the file root_


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

Define variable:

var _name_ type = _expression_



- It's possible use values _signed_ and _unsigned_.
    -  int8, int16, int32 e int64

- In Golang, it is not allowed to declared variable and not use it. For these cases we should use _ after variable. The example below shows this case:

```go
package main

import "fmt"

func main(){
    var x int  //not using the variable
    var y int = 1

    fmt.Println(y)
}
```

```bash
go run main.go
# command-line-arguments
./main.go:9:9: x declared and not used
```

> ![IMPORTANT] 
> When use global variable
> When declare global variabels (package level) the compiler is erro-free.

#### Declaration flavors
- Global Variavels
```go
var x int = 1 //atribute value
var y int //without value
func main(){ 

}
```
- Local variabels
```go
func main(){
    // These local variable can only be seen in this function
    var x int
    var buzzy = true
}
```
- Short declaration variable
```go
    // Declare
    // Atribut value
    // Define type 
func main(){
    x := 3
}
```
## Value Zero

> ![NOTE]
> Here use documentation about package **fmt**
> [fmt package](https://pkg.go.dev/fmt)

[Value Zero](variables/variable-zero.go)


In my studies about zero:

- When not assign value the variable in Golang assign zero to it;
- Values `bools` adding of the value `false` [false = 0 e true = 1]
- Em variaveis sem retorno precisamos usar
- An uninitialized string in Go comes with the value: ""

## Named type

In Golang it's possible create a type, for example


```go
type name string
```

### Variable lifetime

Donavan(2017). _The lifetime of a variable is the time interval during which it exists while the program runs._ (p. 73) The language programming Golang

### Pointers

> ![NOTE] 
> 
> I create folder to files pointers, but this subject is within variables

> ![IMPORTANT]
> 
> Best articles about Pointers

Using pointers in Go is useful when you need to make values mutable or when you want to pass large values to functions without copying them. In my code [pointer-function.go](basic/pointers/pointer-function.go) I'm not trying to say that the sums are incorrect; it's just a demonstration.

[Go 101](https://go101.org/article/pointer.html)