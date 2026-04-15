package main

import "strings"

func IntegerToRoman(num int) string {
	roman := ""
	roman += strings.Repeat("M", num/1000)
	num %= 1000
	roman += strings.Repeat("CM", num/900)
	num %= 900
	roman += strings.Repeat("D", num/500)
	num %= 500
	roman += strings.Repeat("CD", num/400)
	num %= 400
	roman += strings.Repeat("C", num/100)
	num %= 100
	roman += strings.Repeat("XC", num/90)
	num %= 90
	roman += strings.Repeat("L", num/50)
	num %= 50
	roman += strings.Repeat("XL", num/40)
	num %= 40
	roman += strings.Repeat("X", num/10)
	num %= 10
	roman += strings.Repeat("IX", num/9)
	num %= 9
	roman += strings.Repeat("V", num/5)
	num %= 5
	roman += strings.Repeat("IV", num/4)
	num %= 4
	roman += strings.Repeat("I", num)
	return roman
}
