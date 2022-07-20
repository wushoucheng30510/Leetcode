package main

import "fmt"

func main() {

	test1 := "flower"
	test2 := "flight"
	test := []string{test1, test2}
	out := longestCommonPrefix(test)
	fmt.Println(out)
}

func longestCommonPrefix(strs []string) string {

	compareString := strs[0]
	max := len(strs[0])
	i := 0

	for _, word := range strs {
		for index, letter := range word {
			if index > len(compareString)-1 {
				break
			}
			if byte(letter) != compareString[index] {
				break
			}
			i++
		}
		if i <= max {
			max = i
		}
		i = 0
	}
	return strs[0][0:max]

}
