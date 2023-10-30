# Study Notes on Golang

## Basic Concepts

### Camel Case

[Golang](https://golang.org/) uses Camel Case for styling variable and function names.

> [!Important]
> **Best Practices:**
> - `QuoteRuneToASCII` or `parseRequestLine`
>
> **Not Recommended:**
> - `quote_rune_to_ASCII` or `parse_request_line`

### Reserved Words

```go
break     default      func   interface   select
case      defer        go     map         struct 
chan      else         goto   package     switch
const     fallthrough  if     range       type
continue  for          import return      var
```

### Types

```go
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
```

### Binary Operators

```go
* / % << >>& &^
+ - | ^
== ≠ < <= > >=
&&
||
```
> [!Important]
> The operator % is applied only to integer numbers.

---
## Hello World

### Simple Package and How to Use

By default, every Golang project has a file named `main.go`, which serves as the entry point of the project.

```go
package main

import "fmt"

func main(){
    fmt.Println("Hello, my friend")
}
```

### Line by Line Explanation

```package main```

Every Go file is associated with a package, and the `main` package is used only for the project's entry point.

[Example](simple-packages-and-main)

```import fmt```

This line imports the "fmt" package, which is used to print messages to the terminal.

```func main() {}```

The `main` function is the entry point of the Go program, and every core file should have this function.

```fmt.Println("Hello, my friend")```

This line uses the `fmt` package's `Println` function to print the message "Hello, my friend" to the terminal.

> [!Important]
> What happens when you don't declare the `main` package in the root folder?
> _You will encounter an error when the `main` function is missing in the root file._

```bash
$ go run main.go 
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```

_When the `main` package is missing in the root file:_

```bash
$ go run main.go
package command-line-arguments is not a main package
```

---

## Variables

### Starting with the Basics

To define a variable in Go:

```
var name type = expression
```

- You can use both signed and unsigned values:
    -  int8, int16, int32, int64

- In Go, it's not allowed to declare a variable and not use it. In such cases, you should use an underscore `_` after the variable. Here's an example:

```go
package main

import "fmt"

func main(){
    var x int  // not using the variable
    var y int = 1

    fmt.Println(y)
}
```

```bash
go run main.go
# command-line-arguments
./main.go:9:9: x declared and not used
```

> [!Important]
> When to use global variables?
> When declaring global variables (at the package level), the compiler won't produce errors.

#### Declaration Variations
- Global Variables
```go
var x int = 1 // assigns a value
var y int // without a value
func main(){ 

}
```
- Local Variables
```go
func main(){
    // These local variables can only be seen in this function
    var x int
    var buzzy = true
}
```
- Short Declaration of Variables
```go
func main(){
    x := 3 // declares, assigns a value, and defines the type
}
```

## Zero Value

> [!Note]
> You can find information about the `fmt` package in the [fmt package documentation](https://pkg.go.dev/fmt).

[Zero Value](variables/variable-zero.go)

In my studies about zero values:

- When you don't assign a value to a variable in Go, it is assigned a zero value.
- For boolean variables, `false` corresponds to 0, and `true` corresponds to 1.
- In variables with no return, you need to use an underscore `_`.
- An uninitialized string in Go comes with the value: "".

## Named Types

In Go, you can create named types. For example:

```go
type name string
```

### Variable Lifetime

According to Donavan (2017), "the lifetime of a variable is the time interval during which it exists while the program runs" (p. 73).

### Pointers

> [!Note] 
> I created a folder for pointer-related files, but this topic is within the variables section.

> [!Important]
> Best articles about Pointers

Using pointers in Go is useful when you need to make values mutable or when you want to pass large values to functions without copying them. In my code [pointer-function.go](basic/pointers/pointer-function.go), I'm not trying to say that the sums are incorrect; it's just a demonstration.

[Go 101](https://go101.org/article/pointer.html)

---