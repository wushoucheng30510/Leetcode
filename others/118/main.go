package main

import (
	"fmt"
)

func main() {
	output := generate(10)
	for _, k := range output {
		fmt.Println(k)
	}
}

func generate(numRows int) [][]int {
	output := [][]int{}
	currentSlice := []int{}
	// i means layers
	for i := 0; i < numRows; i++ {
		currentSlice = nil
		// j means the elements
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				currentSlice = append(currentSlice, 1)
			} else {
				currentSlice = append(currentSlice, output[i-1][j-1]+output[i-1][j])
			}
		}
		output = append(output, currentSlice)
	}
	return output
}
