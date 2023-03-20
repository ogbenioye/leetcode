package main

import "fmt"

func twoSums(nums []int, target int) []int {

	var indicies []int

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				indicies = append(indicies, i, j)
				fmt.Println(indicies)
				return indicies
			}

		}

	}
	return indicies
}

//make a copy of that array, take a value from first array, add to another
func main() {
	nums := []int{1, 4, 6, 2}

	twoSums(nums, 6)
}
