package main

import "fmt"

func main() {
	fmt.Println(addBinary("111", "101"))
	fmt.Println(addBinary2("111", "101"))
}

func addBinary(a string, b string) string {
	shortString, longString := stringCompare(reverse(a), reverse(b))
	carry, currentDigit, output := "0", "", ""

	for i := 0; i < len(shortString); i++ {
		carry, currentDigit = digitsAdd(string(shortString[i]), string(longString[i]), carry)
		output = currentDigit + output
	}

	diff := len(longString) - len(shortString)
	if diff > 0 {
		for i := len(longString) - diff; i < len(longString); i++ {
			carry, currentDigit = digitsAdd("0", string(longString[i]), carry)
			output = currentDigit + output
		}
	}
	if carry == "1" {
		output = "1" + output
	}
	return output
}

func reverse(input string) string {
	output := ""
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	return output
}

func stringCompare(a, b string) (string, string) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}

func digitsAdd(a, b, oriCarry string) (carry, current string) {
	compareString := a + b + oriCarry
	switch compareString {
	case "000":
		return "0", "0"
	case "001", "010", "100":
		return "0", "1"
	case "110", "011", "101":
		return "1", "0"
	case "111":
		return "1", "1"
	}
	return "0", "0"
}

func addBinary2(a string, b string) string {
	longString, shortString := stringCompare(a, b)
	output := ""
	carry := "0"
	for times := 0; times < len(shortString); times++ {
		if string(longString[len(longString)-1-times]) == "1" && string(shortString[len(shortString)-1-times]) == "1" {
			if carry == "0" {
				carry = "1"
				output = "0" + output
			} else {
				carry = "1"
				output = "1" + output
			}
		} else if string(longString[len(longString)-1-times]) == "1" || string(shortString[len(shortString)-1-times]) == "1" {
			if carry == "0" {
				carry = "0"
				output = "1" + output
			} else {
				carry = "1"
				output = "0" + output
			}
		} else {
			output = carry + output
			carry = "0"
		}
	}
	for i := len(longString) - 1 - len(shortString); i >= 0; i-- {
		if carry == "0" {
			output = string(longString[0:i+1]) + output
			break
		} else {
			if string(longString[i]) == "1" {
				carry = "1"
				output = "0" + output
			} else {
				carry = "0"
				output = "1" + output
			}
		}
	}
	if carry == "1" {
		output = "1" + output
	}
	return output
}
