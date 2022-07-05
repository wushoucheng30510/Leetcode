package main

import "fmt"

func main() {

	fmt.Print("Leetcode practice: ")

}

func TwoSum(nums []int, target int) []int {
	mapTwoSum := map[int]int{}
	var output []int
	for index, value := range nums {
		if _, ok := mapTwoSum[target-value]; ok {
			if index != mapTwoSum[target-value] {
				output = append(output, mapTwoSum[target-value], index)
			}
		}
		mapTwoSum[value] = index
	}
	return output
}
