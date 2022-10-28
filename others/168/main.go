package main

func main() {
	// fmt.Println((byte('A')))
	// fmt.Println(string(byte(65)))

	convertToTitle(53)
}

func convertToTitle(columnNumber int) string {
	var output string
	var left int
	for columnNumber > 0 {
		left = columnNumber % 26
		columnNumber /= 26
		if left == 0 {
			output = "Z" + output
			columnNumber -= 1
		} else {
			output = string(byte('A')-1+byte(left)) + output
		}

	}
	return output
}
