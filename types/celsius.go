package types

import "fmt"

type Celsius float64

// Converts celsius to string
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f C", c)
}

// Converts celsius to kelvin
func (c Celsius) Kelvin() Kelvin {
	return Kelvin(c + 273.15)
}

// Converts celsius to fahrenheit
func (c Celsius) Fahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
