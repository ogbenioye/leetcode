package main

import "fmt"

/*
	same algorithm as what is in "spiral.go" but in a different direction
*/

func spiral(matrix [][]int) int {
	//rows
	top, bottom := 0, len(matrix)-1
	//columns
	left, right := 0, len(matrix[0])-1

	direction := "down"

	result := make([]int, 0)

	for top <= bottom && left <= right {
		if direction == "down" {
			for i := top; i <= bottom; i++ {
				result = append(result, matrix[i][left])
			}
			left++
			direction = "right"
		} else if direction == "right" {
			for i := left; i <= right; i++ {
				result = append(result, matrix[bottom][i])
			}
			bottom--
			direction = "up"
		} else if direction == "up" {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][right])
			}
			right--
			direction = "left"
		} else if direction == "left" {
			for i := right; i >= left; i-- {
				result = append(result, matrix[top][i])
			}
			top++
			direction = "down"
		}
	}
	fmt.Println(result)
	return result[len(result)-1]
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	out := spiral(matrix)
	// if "house" == "hose" {
	fmt.Println(out)
	// }
}
