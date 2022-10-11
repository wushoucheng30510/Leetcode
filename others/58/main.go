package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	start := false
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if start == true && s[i] == 32 {
			break
		}
		if s[i] == 32 {
			continue
		} else {
			start = true
		}
		if start && s[i] != 32 {
			count++
		}

	}
	return count
}

func lengthOfLastWord2(s string) int {
	start := false
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if start == true && string(s[i]) == " " {
			break
		}
		if string(s[i]) == " " {
			continue
		} else {
			start = true
		}
		if start && string(s[i]) != " " {
			count++
		}

	}
	return count
}
