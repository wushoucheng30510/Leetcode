package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1

	for l < r {
		for !isAlphabetInLowerCase(byte(s[l])) && l < r {
			l++
		}
		for !isAlphabetInLowerCase(byte(s[r])) && l < r {
			r--
		}
		if l < r && s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphabetInLowerCase(s byte) bool {
	if s >= byte('a') && s <= byte('z') {
		return true
	}
	return false
}
