package main

import "fmt"

func main() {
	first_element := true

	fmt.Println("Делятся на 3 без остатка:")
	for i := 1; i < 101; i++ {
		if (i % 3) == 0 {
			if !first_element {
				fmt.Print(", ")
			} else {
				first_element = false
			}
			fmt.Print(i)
		}
	}
}
