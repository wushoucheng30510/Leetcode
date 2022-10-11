package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 5}, 3))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
