package main

func main() {

}

func romanToInt(s string) int {

	romanValue := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	out := 0
	for index, num := range s {
		out += romanValue[num]
		var minus2, minus20, minus200 bool

		if index > 0 {
			minus2 = (num == 'V' || num == 'X') && s[index-1] == 'I'
			minus20 = (num == 'L' || num == 'C') && s[index-1] == 'X'
			minus200 = (num == 'D' || num == 'M') && s[index-1] == 'C'
		}

		if minus2 {
			out -= 2
		}

		if minus20 {
			out -= 20
		}

		if minus200 {
			out -= 200
		}
	}
	return out

}
