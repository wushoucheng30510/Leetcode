package main

import "fmt"

func main() {
	if 01&1 == 1 {
		fmt.Println("correct")
	}
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
