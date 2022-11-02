package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))

}

func longestPalindrome(s string) string {
	var max string

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if palindrome(s[i : j+1]) {
				if len(max) < len(s[i:j+1]) {
					max = s[i : j+1]
				}
			}
		}
	}

	return max
}

func palindrome(s string) bool {
	var mid int

	if len(s)%2 == 0 {
		mid = len(s)/2 - 1
	} else {
		mid = len(s) / 2
	}

	for i := 0; i <= mid; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	longestString := ""

	// abcba
	for i := range s {
		left := i
		right := i

		for {
			if left-1 < 0 || right+1 >= len(s) {
				break
			}
			if s[left-1] == s[right+1] {
				left--
				right++
			} else {
				break
			}
		}
		currentString := s[left : right+1]
		if len(longestString) < len(currentString) {
			longestString = currentString
		}
	}

	// c'bb'd
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			continue
		}
		left := i
		right := i + 1
		for {
			if left-1 < 0 || right+1 >= len(s) {
				break
			}
			if s[left-1] == s[right+1] {
				left--
				right++
			} else {
				break
			}
		}
		currentString := s[left : right+1]
		if len(longestString) < len(currentString) {
			longestString = currentString
		}

	}
	return longestString
}

func longestPalindrome3(s string) string {
	longestPalyndrome := ""

	// Centered on the letter
	for i := range s {
		currentLongestPalyndrome := ""
		left := i
		right := i

		for {

			if left-1 < 0 || right+1 >= len(s) {
				break
			}
			if s[left-1] == s[right+1] {
				left--
				right++
			} else {
				break
			}
		}
		currentLongestPalyndrome = s[left : right+1]

		if len(currentLongestPalyndrome) > len(longestPalyndrome) {
			longestPalyndrome = currentLongestPalyndrome
		}
	}

	// Centered between letter
	for i := 0; i < len(s)-1; i++ {
		currentLongestPalyndrome := ""

		if s[i] != s[i+1] {
			continue
		}

		left := i
		right := i + 1

		for {

			if left-1 < 0 || right+1 >= len(s) {
				break
			}
			if s[left-1] == s[right+1] {
				left--
				right++
			} else {
				break
			}
		}
		currentLongestPalyndrome = s[left : right+1]

		if len(currentLongestPalyndrome) > len(longestPalyndrome) {
			longestPalyndrome = currentLongestPalyndrome
		}
	}

	return longestPalyndrome
}
