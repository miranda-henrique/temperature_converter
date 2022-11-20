package main

import (
	"fmt"
)

func main() {
	var inputTempType int
	var inputChecker bool = false
	var outputTempType int
	var outputChecker bool = false
	var temp float64

	const CELSIUS int = 1
	const FAHRENHEIT int = 2
	const KELVIN int = 3

	fmt.Println("Temperature converter\n")

	for inputChecker == false {
		fmt.Print("Please enter the current temperature scale: \n1 - Celsius \n2 - Fahrenheit \n3 - Kelvin \n\n Option: ")
		fmt.Scanln(&inputTempType)

		if inputTempType != 1 && inputTempType != 2 && inputTempType != 3 {
			fmt.Println("\n#### Please enter a valid option ####\n")
		} else {
			inputChecker = true
		}
	}

	for outputChecker == false {
		fmt.Print("\n\nWhat scale will we convert to? \n1 - Celsius \n2 - Fahrenheit \n3 - Kelvin \n\n Option: ")
		fmt.Scanln(&outputTempType)

		if outputTempType != 1 && outputTempType != 2 && outputTempType != 3 {
			fmt.Println("\n#### Please enter a valid option ####\n")
		} else {
			outputChecker = true
		}
	}

	fmt.Print("Enter temperature: ")
	fmt.Scanln(&temp)

	if inputTempType == outputTempType {
		fmt.Println("Same scales. No need to convert =)")

	} else if inputTempType == CELSIUS {
		if outputTempType == FAHRENHEIT {
			celsiusToFahrenheit(temp)
		} else if outputTempType == KELVIN {
			celsiusToKelvin(temp)
		}

	} else if inputTempType == FAHRENHEIT {
		if outputTempType == CELSIUS {
			fahrenheitToCelsius(temp)
		} else if outputTempType == KELVIN {
			fahrenheitToKelvin(temp)
		}

	} else {
		if outputTempType == CELSIUS {
			kelvinToCelsius(temp)
		} else if outputTempType == FAHRENHEIT {
			kelvinToFahrenheit(temp)
		}
	}
}

func celsiusToFahrenheit(c float64) {
	f := ((9 * c) / 5) + 32

	fmt.Printf("%.2f ºCelsius = %.2f ºFahrenheit", c, f)
}

func celsiusToKelvin(c float64) {
	k := c + 273

	fmt.Printf("%.2f ºCelsius = %.2f ºKelvin", c, k)
}

func fahrenheitToCelsius(f float64) {
	c := ((f - 32) / 9) / 5

	fmt.Printf("%.2f ºFahrenheit = %.2f ºCelsius", f, c)
}

func fahrenheitToKelvin(f float64) {
	c := ((f - 32) / 9) / 5
	k := c + 273

	fmt.Printf("%.2f ºFahrenheit = %.2f ºKelvin", f, k)
}

func kelvinToCelsius(k float64) {
	c := k - 273

	fmt.Printf("%.2f ºKelvin = %.2f ºCelsius", k, c)
}

func kelvinToFahrenheit(k float64) {
	c := k - 273
	f := ((9 * c) / 5) + 32

	fmt.Printf("%.2f ºKelvin = %.2f ºFahrenheit", k, f)
}
