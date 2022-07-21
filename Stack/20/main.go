package main

func main() {

}

func isValid(s string) bool {
	paraSlice := []rune{}
	for _, parath := range s {
		// push if needed
		if pushParathess(parath) {
			paraSlice = append(paraSlice, parath)
		} else {
			// need pop
			// if slice is empty, return false
			if len(paraSlice) == 0 {
				return false
			}
			// check if the corresponding parathness is correct
			if checkParathess(parath, paraSlice[len(paraSlice)-1]) {
				paraSlice = paraSlice[0 : len(paraSlice)-1]
			} else {
				return false
			}

		}
	}
	// if there is any elements in the slice before at the end of the function, return false
	if len(paraSlice) != 0 {
		return false
	}
	return true
}

func pushParathess(s rune) bool {
	if s == '{' || s == '(' || s == '[' {
		return true
	}
	return false
}

func checkParathess(checkString, stackString rune) bool {
	if checkString == '}' && stackString != '{' {
		return false
	}
	if checkString == ']' && stackString != '[' {
		return false
	}
	if checkString == ')' && stackString != '(' {
		return false
	}
	return true
}
