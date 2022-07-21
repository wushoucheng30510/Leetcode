package main

import "fmt"

func main() {
	out := strStr("Hello", "ll")
	fmt.Println(out)
}

func strStr(haystack string, needle string) int {
	for index, wordHayStack := range haystack {
		if len(haystack)-index < len(needle) {
			return -1
		}
		if wordHayStack == rune(needle[0]) {
			fmt.Println("Haystack: ", haystack[index:index+len(needle)])
			same := compareTwoString(string(haystack[index:index+len(needle)]), needle)
			if same {
				return index
			}
		}
	}
	return -1
}

func compareTwoString(a, b string) bool {
	if a == b {
		return true
	}
	return false
}
