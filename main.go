package main

import (
	"fmt"
	"temperatureconversion/types"
)

func main() {
	fmt.Println("Temperature Conversion")
	fmt.Println("----------------------")
	choice := 0
	for choice != 4 {
		choice = selectStartingTemperatureType()
		switch choice {
		case 1:
			var temperature types.Celsius = types.Celsius(enterTemperature())
			fmt.Println("----------------------")
			fmt.Println(temperature, "° Celsius is:")
			fmt.Println("----------------------")
			fmt.Println(temperature.Fahrenheit(), "° Fahrenheit")
			fmt.Println(temperature.Kelvin(), "° Kelvin")
			fmt.Println("----------------------")
		case 2:
			var temperature types.Fahrenheit = types.Fahrenheit(enterTemperature())
			fmt.Println("----------------------")
			fmt.Println(temperature, "° Fahrenheit is:")
			fmt.Println("----------------------")
			fmt.Println(temperature.Celsius(), "° Celsius")
			fmt.Println(temperature.Kelvin(), "° Kelvin")
			fmt.Println("----------------------")
		case 3:
			var temperature types.Kelvin = types.Kelvin(enterTemperature())
			fmt.Println("----------------------")
			fmt.Println(temperature, "° Kelvin is:")
			fmt.Println("----------------------")
			fmt.Println(temperature.Celsius(), "° Celsius")
			fmt.Println(temperature.Fahrenheit(), "° Fahrenheit")
			fmt.Println("----------------------")
		default:
			fmt.Println("Please Enter a valid choice from 1 to 4")
		}
		fmt.Println()

	}

	fmt.Println("Goodbye!")

}

// Enter temperature
func enterTemperature() float64 {
	var temperature float64
	fmt.Print("Enter temperature: ")
	fmt.Scan(&temperature)
	return temperature
}

// Select starting temperature type (Celsius, Fahrenheit, or Kelvin)
func selectStartingTemperatureType() int {
	choice := 0
	fmt.Println("Select starting temperature type:")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("4. Exit")
	fmt.Println("----------------------")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	return choice
}
