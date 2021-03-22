package main

import "fmt"

func conv(metres float64) float64 {
	var metersInInch = 0.3048
	output := metres * (1 / metersInInch)
	return output
}

func main() {
	fmt.Print("Enter Length (m): ")
	var input float64
	fmt.Scanf("%f", &input)

	inches := conv(input)

	fmt.Printf("Length is %f foot", inches)
}
