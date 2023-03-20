package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	//make a copy of nums
	//delete every element in nums
	//append the unique elements into nums

	newArray := nums[:]
	nums = nums[:0]

	mp := make(map[int]int)

	for _, v := range newArray {
		mp[v] = v
		_, ok := mp[v]

		if ok == false {
			mp[v] = v
		}
	}

	for _, v := range mp {
		nums = append(nums, v)
	}
	sort.Ints(nums)
	fmt.Println(nums)
	return len(nums)
}

func main() {
	nums := []int{0, 0, 1, 1}

	removeDuplicates(nums)
}
