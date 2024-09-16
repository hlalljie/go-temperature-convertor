package types

import "fmt"

type Kelvin float64

// Converts kelvin to string
func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f K", k)
}

// Converts kelvin to celsius
func (k Kelvin) Celsius() Celsius {
	return Celsius(k - 273.15)
}

// Converts kelvin to fahrenheit
func (k Kelvin) Fahrenheit() Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}
