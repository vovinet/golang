package main

import "fmt"

func min(x []int) int {
	min := x[0]

	for i := 1; i < len(x); i++ {
		if min > x[i] {
			min = x[i]
		}
	}

	return min
}

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	min := min(x)

	fmt.Println("Исходный массив: ", x)

	fmt.Println("Самый маленький элемент массива: ", min)
}
