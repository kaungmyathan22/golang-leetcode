package main

import (
	"unicode/utf8"
)

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func FirstPalindrome(words []string) string {
	for _, str := range words {
		reverse := Reverse((str))
		if reverse == str {
			return str
		}

	}
	return ""
}

// func main() {
// 	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
// 	fmt.Println(firstPalindrome([]string{"notapalindrome", "racecar"}))
// 	fmt.Println(firstPalindrome([]string{"def", "ghi"}))
// }
