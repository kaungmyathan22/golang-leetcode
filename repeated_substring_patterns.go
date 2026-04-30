package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	substrLen := len(s) / 2
	for i := 1; i <= substrLen; i++ {
		if len(s)%i == 0 {
			comp := strings.Repeat(s[:i], len(s)/i)
			if comp == s {
				return true
			}
		}
	}
	return false
}
