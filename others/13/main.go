package main

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for index, word := range s {
		sum += romanMap[string(word)]
		if index-1 >= 0 {
			if (string(word) == "V" || string(word) == "X") && s[index-1] == 'I' {
				sum -= 2
			}
			if (string(word) == "L" || string(word) == "C") && s[index-1] == 'X' {
				sum -= 20
			}
			if (string(word) == "D" || string(word) == "M") && s[index-1] == 'C' {
				sum -= 200
			}
		}

	}
	return sum
}
