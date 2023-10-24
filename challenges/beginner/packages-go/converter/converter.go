package converter


type Celsius float64
type Fahrenheit float64

const (
	AbsluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BolingC Celsius = 100
)
func ConverterCelsiusToFahrenheit(c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5 + 32)
}
func ConverterFahrenheitToCelsius(f Fahrenheit) Celsius{
	return Celsius((f-32)*5/9)
}