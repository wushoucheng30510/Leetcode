package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(1111111))
}

func isHappy(n int) bool {
	sum := 0
	mapHistroy := map[int]int{}

	var cal int
	for {
		for n != 0 {
			cal = n % 10
			sum += cal * cal
			n /= 10
		}

		if sum == 1 {
			return true
		}

		if _, ok := mapHistroy[sum]; ok {
			return false
		} else {
			mapHistroy[sum] = 1
		}

		n = sum
		sum = 0
	}
	return false
}
