package main

import (
	"strconv"
)

func main() {

}

func isPalindrome(x int) bool {
	originalX := x
	if x < 0 {
		return false
	}
	sum := 0

	for x != 0 {
		sum = sum*10 + x%10
		x = x / 10
	}
	return originalX == sum

}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	numString := strconv.FormatInt(int64(x), 10)

	for i, _ := range numString {
		if i >= len(numString) {
			return true
		}
		if numString[i] != numString[len(numString)-i] {
			return false
		}
	}
	return true
}
