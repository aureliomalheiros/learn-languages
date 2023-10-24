package main

import "fmt"

type Celsius float64

func main(){
	temperatureCelsius 		:= Celsius(25.0)
	temperatureFahrenheit 	:= celsiusToFahrenheit(temperatureCelsius)

	fmt.Printf("%1.f Celsius same a %1.f Fahrenheit. \n", temperatureCelsius, temperatureFahrenheit)
}

func celsiusToFahrenheit(c Celsius) float64 {
	return float64(c)*9/5 + 32
}