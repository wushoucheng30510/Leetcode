package main

import "fmt"

func main() {
	input := []string{"abcabcbb", "bbbbb", "pwwkew", " ", "dvdf", "jbpnbwwd"}
	for _, k := range input {
		fmt.Println("Method 1:", lengthOfLongestSubstring(k), "  Method 2: ", lengthOfLongestSubstring2(k))
		fmt.Println()
	}

}

func lengthOfLongestSubstring(s string) int {
	count := 0
	maxCount := 0
	mapHistory := map[byte]int{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := mapHistory[s[j]]; ok {
				for key, _ := range mapHistory {
					delete(mapHistory, key)
				}
				maxCount = max(maxCount, count)
				count = 0
				break
			} else {
				mapHistory[s[j]] = 1
				count++

			}
		}
		maxCount = max(maxCount, count)
		count = 0
	}
	return maxCount
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLongestSubstring2(str string) int {
	m := make(map[rune]int)
	left, max := 0, 0
	for index, ch := range str {
		if _, ok := m[ch]; ok && m[ch] >= left {
			if index-left > max {
				max = index - left
			}
			left = m[ch] + 1
		}
		m[ch] = index
	}

	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}
