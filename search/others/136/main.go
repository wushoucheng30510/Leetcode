package main

import "fmt"

func main() {
	input := [][]int{[]int{2, 2, 1}, []int{4, 1, 2, 1, 2}, []int{1}}
	for _, nums := range input {
		fmt.Println("Output: ", singleNumber2(nums))
	}
}

func singleNumber(nums []int) int {
	mapSearching := map[int]int{}
	for _, num := range nums {
		_, ok := mapSearching[num]
		if !ok {
			mapSearching[num] = 1
		} else {
			mapSearching[num] += 1
		}
	}
	var out int
	for key, count := range mapSearching {
		if count == 2 {
			continue
		} else {
			out = key
			break
		}
	}
	return out
}

func singleNumber2(nums []int) int {
	var sol int
	for i := 0; i < len(nums); i++ {
		sol = sol ^ nums[i]
	}
	return sol
}
