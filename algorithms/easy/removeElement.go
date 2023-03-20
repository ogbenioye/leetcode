package main

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}

	removeElement(arr, 2)
}
