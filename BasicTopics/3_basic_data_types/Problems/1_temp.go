/*
Problem:
Temperature Conversion:

Write a program that converts a temperature in Celsius to Fahrenheit. Prompt the user to enter a temperature in Celsius, perform the conversion, and print the result in Fahrenheit.
*/
//first import necessary package
//take function main
//take variable for celsius and fahrenheit
//by using formula build the logic for this program.
package main

import "fmt"

func main() {

	var ftemp float64 = 0
	var ctemp float64 = 0
	ftemp = (ctemp * 1.8) + 32

	fmt.Printf("enter temperature in celsisus:")
	fmt.Scan("%f", &ctemp)
	//ftemp = (ctemp * 1.8) + 32
	fmt.Printf("temperature in farenheit:%2f", ftemp)

}
