package main

import "fmt"

/*
	- Runtime: 1ms; Memory usage: 2.1MB (less than 100% of Go online submissions)
	- Faster than 70.37% of Go Online subbmissions as at 21-08-2023
	- Steps:
		Create a n x n matrix with zero as it elements.
		Replace each zero with the right number.
		To determine the right index for each number, insert numbers incrementally in a clockwise spiral direction

*/

func generateMatrix(n int) [][]int {
	/*
		Direction control:
		left to right = columns
		top to bottom = rows
	*/
	direction := "right"
	top, bottom := 0, n-1
	left, right := 0, n-1

	//to create the matrix and fill it with zeros
	matrix := make([][]int, 0)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	limit := n * n

	num := 1

	for num <= limit {
		if direction == "right" {
			for i := left; i <= right; i++ {
				matrix[top][i] = num
				num++
			}
			top++
			direction = "down"
		} else if direction == "down" {
			for i := top; i <= bottom; i++ {
				matrix[i][right] = num
				num++
			}
			right--
			direction = "left"
		} else if direction == "left" {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
			direction = "up"
		} else if direction == "up" {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
			direction = "right"
		}
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
