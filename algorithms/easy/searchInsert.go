package main

func searchInsert(nums []int, target int) int {
	vari := 0
	for i, v := range nums {
		if v == target {
			vari = i
		} else if v < target && i < len(nums)-1 && nums[i+1] > target {
			vari = i + 1
		} else if v+1 == target || nums[len(nums)-1] < target && i == len(nums)-1 {
			vari = i + 1
		}
	}
	return vari
}

func main() {
	arr := []int{2, 3, 4, 7, 8, 9}
	searchInsert(arr, 11)
}
