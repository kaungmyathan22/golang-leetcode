package main

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}

// func main() {
// 	fmt.Println(lengthOfLastWord("Hello World"))
// 	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
// 	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
// }
