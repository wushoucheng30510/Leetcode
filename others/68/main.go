package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	l := 1
	r := x
	for l <= r {
		mid := (l + r) / 2
		if mid <= x/mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
