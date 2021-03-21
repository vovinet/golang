package main

import "fmt"

func main() {
	var metersInInch = 0.3048

	fmt.Print("Enter Length (m): ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * (1 / metersInInch)

	fmt.Printf("Length is %.2f foot", output)
}
