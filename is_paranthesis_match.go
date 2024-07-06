package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		if v, ok := mapping[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

// func main() {
// 	fmt.Println(isValid("()"))
// 	fmt.Println(isValid("()[]{}"))
// 	fmt.Println(isValid("(]"))
// }
