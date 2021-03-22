package main

import "fmt"

func div3(start, stop int) []int {
	var result []int
	for i := start; i < stop+1; i++ {
		if (i % 3) == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	start := 1
	stop := 100

	my_result := div3(start, stop)

	fmt.Println("Делятся на 3 без остатка:")
	fmt.Println(my_result)
}
