package main

import (
	"math"
	"strings"
)

func StringToInteger(s string) int {
	sign := 1
	s = strings.TrimSpace(s)
	if len(s) > 0 && s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if len(s) > 0 && s[0] == '+' {
		sign = 1
		s = s[1:]
	}
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num = num*10 + int(s[i]-'0')
		if sign*num > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*num < math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * num
}
