package main

import (
	"fmt"
)

// Convert Celcius to Farenheit
func convertC2F(temperature float64, unit string) string{
	var conversion float64
	if unit == "C" || unit == "c"{
		conversion = (temperature*9/5)+32
	}
	return fmt.Sprintf("%.2f %s is %.2f F", temperature, unit, conversion)
}

// Convert Farenheit to Celcius
func convertF2C(temperature float64, unit string) string{
	var conversion float64
	if unit == "F" || unit == "f"{
		conversion = (temperature-32)*5/9
	}
	return fmt.Sprintf("%.2f %s is %.2f C", temperature, unit, conversion)
}


func main(){
	var temperature float64 // Declare a variable to store the input for the temperature
	var unit string // Declare a variable to store the input for the unit
	fmt.Print("Enter temperature value: ")
	fmt.Scanln(&temperature)
		fmt.Print("Enter C or F: ")
	fmt.Scanln(&unit)
	if unit != "C" && unit != "F"{
		fmt.Println("Error: Invalid Unit")
		return
	} else if unit == "F" || unit == "f"{
		fmt.Println(convertF2C(temperature, unit))
	} else{
		fmt.Println(convertC2F(temperature, unit))
	}	
}