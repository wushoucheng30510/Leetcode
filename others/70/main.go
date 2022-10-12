package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	s1, s2, result := 1, 2, 0

	for i := 3; i <= n; i++ {
		result = s1 + s2
		s1 = s2
		s2 = result
	}

	return result
}
