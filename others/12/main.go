package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

//I can be placed before V (5) and X (10) to make 4 and 9.
//X can be placed before L (50) and C (100) to make 40 and 90.
//C can be placed before D (500) and M (1000) to make 400 and 900.

type roman struct {
	symbol string
	value  int
}

func intToRoman(num int) string {
	output := ""
	sliceRoman := []roman{{"M", 1000}, {"CM", 900}, {"D", 500}, {"CD", 400}, {"C", 100}, {"XC", 90}, {"L", 50}, {"XL", 40}, {"X", 10}, {"IX", 9}, {"V", 5}, {"IV", 4}, {"I", 1}}

	for _, roman := range sliceRoman {
		for num >= roman.value {
			output = output + roman.symbol
			num -= roman.value
		}
	}
	return output
}
