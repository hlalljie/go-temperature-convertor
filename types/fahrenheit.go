package types

import "fmt"

type Fahrenheit float64

// Converts fahrenheit to string
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f F", f)
}

// Converts fahrenheit to celsius
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Converts fahrenheit to kelvin
func (f Fahrenheit) Kelvin() Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}
