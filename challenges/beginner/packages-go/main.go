package main

import (
	"fmt"
	"module/converter"
)
func main() {
	fmt.Println(converter.ConverterCelsiusToFahrenheit(10))
	fmt.Println(converter.ConverterFahrenheitToCelsius(20))
}